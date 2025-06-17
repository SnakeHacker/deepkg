package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SnakeHacker/deepkg/admin/internal/config"
	utils_nebula "github.com/SnakeHacker/deepkg/admin/internal/utils/nebula"
	"github.com/golang/glog"
	nebula "github.com/vesoft-inc/nebula-go/v3"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/admin.yaml", "the config file")

func main() {
	flag.Parse()
	defer glog.Flush()

	var c config.Config
	if err := conf.Load(*configFile, &c); err != nil {
		glog.Error("加载配置文件失败: %v", err)
	}

	if err := run(c.Nebula); err != nil {
		glog.Error(err)
	}
}
func run(conf utils_nebula.NebulaSessionConfig) (err error) {

	session, err := utils_nebula.NewNebulaSession(conf)
	if err != nil {
		glog.Error(err)
		return
	}
	defer session.Release()

	// 4. 检查服务状态
	glog.Info("检查 Nebula 服务状态...")

	// 4.1 检查 Graph 服务
	_, err = session.Execute("SHOW HOSTS;")
	if err != nil {
		glog.Error(err)
		return
	}

	// 4.2 检查 Storage 服务
	_, err = session.Execute("SHOW HOSTS STORAGE;")
	if err != nil {
		glog.Error(err)
		return
	}

	// 4.3 检查 Meta 服务
	_, err = session.Execute("SHOW HOSTS META;")
	if err != nil {
		glog.Error(err)
		return
	}

	// 5. 检查现有空间
	glog.Info("检查现有空间...")
	_, err = execNGQL(session, "SHOW SPACES;")
	if err != nil {
		glog.Error(err)
		return
	}

	// 6. 如果 demo_space 已存在，先删除
	_, err = execNGQL(session, "DROP SPACE IF EXISTS demo_space;")
	if err != nil {
		glog.Error(err)
		return
	}
	time.Sleep(2 * time.Second) // 等待空间删除完成

	// 7. 创建新空间
	glog.Info("创建新空间...")
	_, err = execNGQL(session, "CREATE SPACE demo_space(vid_type=INT64, partition_num=1, replica_factor=1);")
	if err != nil {
		glog.Error(err)
		return
	}
	time.Sleep(10 * time.Second)

	// 8. 使用新空间
	_, err = execNGQL(session, "USE demo_space;")
	if err != nil {
		glog.Error(err)
		return
	}

	// 9. 创建标签和边类型（通用三元组）
	glog.Info("创建三元组模型的 Schema...")
	_, err = execNGQL(session, "CREATE TAG entity(name string);")
	if err != nil {
		glog.Error(err)
		return
	}
	_, err = execNGQL(session, "CREATE EDGE relation(name string);")
	if err != nil {
		glog.Error(err)
		return
	}
	time.Sleep(5 * time.Second) // 等待 schema 同步

	glog.Info("创建并构建索引...")
	_, err = execNGQL(session, "CREATE TAG INDEX entity_index ON entity(name(20));")
	if err != nil {
		glog.Error(err)
		return
	}
	time.Sleep(10 * time.Second)
	_, err = execNGQL(session, "REBUILD TAG INDEX entity_index;")
	if err != nil {
		glog.Error(err)
		return
	}
	time.Sleep(10 * time.Second) // 等待索引构建完成

	// 10. 插入三元组数据
	glog.Info("插入三元组测试数据...")
	_, err = execNGQL(session, "INSERT VERTEX entity(name) VALUES 1001:('张三'), 1002:('李四'), 1003:('王五');")
	if err != nil {
		glog.Error(err)
		return
	}
	_, err = execNGQL(session, "INSERT EDGE relation(name) VALUES 1001->1002:('朋友'), 1002->1003:('老师');")
	if err != nil {
		glog.Error(err)
		return
	}

	// 11. 基于实体名查询其属性与相关实体
	glog.Info("查询实体“张三”的属性与关系: ")
	err = queryByEntityName(session, "张三")
	if err != nil {
		glog.Error(err)
		return
	}
	return nil
}

// 查询指定实体名对应的所有属性和关联实体
func queryByEntityName(session *nebula.Session, name string) (err error) {
	// 第一步：根据名称查 VID
	stmt := fmt.Sprintf("LOOKUP ON entity WHERE entity.name == '%s' YIELD id(vertex) AS vid;", name)
	resp, err := session.Execute(stmt)
	if err != nil || !resp.IsSucceed() || len(resp.GetRows()) == 0 {
		glog.Error(err)
		return
	}
	vid := resp.GetRows()[0].Values[0].GetIVal()

	// 第二步：获取实体属性
	fmt.Printf("\n实体 [%s] 的属性:\n", name)
	resp1, err := execNGQL(session, fmt.Sprintf("FETCH PROP ON entity %d YIELD entity.name;", vid))
	if err != nil {
		log.Printf("查询失败: %v", err)
		// 可以 return err 或做其他处理
	} else {
		printQueryResult(resp1)
	}

	// 第三步：获取关系
	fmt.Printf("\n实体 [%s] 的关联实体及关系:\n", name)
	resp2, err := execNGQL(session, fmt.Sprintf("GO FROM %d OVER relation YIELD dst(edge), relation.name;", vid))
	if err != nil {
		log.Printf("查询失败: %v", err)
		// 可以 return err 或做其他处理
	} else {
		printQueryResult(resp2)
	}
	return nil
}

func execNGQL(session *nebula.Session, stmt string) (*nebula.ResultSet, error) {
	//fmt.Printf("执行: %s\n", stmt)
	resp, err := session.Execute(stmt)
	if err != nil {
		return nil, fmt.Errorf("nGQL执行失败 [%s]: %v", stmt, err)
	}
	if !resp.IsSucceed() {
		return nil, fmt.Errorf("nGQL错误 [%s]: %s", stmt, resp.GetErrorMsg())
	}

	return resp, nil
}
func printQueryResult(resp *nebula.ResultSet) {
	// 打印结果
	colNames := resp.GetColNames()
	rows := resp.GetRows()
	if len(rows) == 0 {
		glog.Info("无数据")
		return
	}

	// 打印列名
	glog.Info("列名:", colNames)

	// 打印数据行
	// 使用 glog 输出数据行
	for i, row := range rows {
		var rowValues []string
		for _, val := range row.Values {
			switch {
			case val.IsSetSVal():
				rowValues = append(rowValues, string(val.GetSVal()))
			case val.IsSetIVal():
				rowValues = append(rowValues, strconv.FormatInt(val.GetIVal(), 10))
			case val.IsSetFVal():
				rowValues = append(rowValues, fmt.Sprintf("%.2f", val.GetFVal()))
			case val.IsSetBVal():
				rowValues = append(rowValues, strconv.FormatBool(val.GetBVal()))
			default:
				rowValues = append(rowValues, fmt.Sprintf("%v", val))
			}
		}
		glog.Infof("第%d行: %v", i+1, rowValues)
	}
}
