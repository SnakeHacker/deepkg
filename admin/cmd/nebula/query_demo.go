package main

import (
	"fmt"
	"log"
	"time"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

func main() {
	if err := run(); err != nil {
		log.Printf("运行出错: %v", err)
	}
}
func run() error {
	// 1. 创建连接池配置
	hostAddress := nebula.HostAddress{Host: "127.0.0.1", Port: 9669}
	poolConfig := nebula.PoolConfig{
		TimeOut:         10 * time.Second,
		IdleTime:        0,
		MaxConnPoolSize: 10,
		MinConnPoolSize: 1,
	}

	// 2. 创建连接池
	pool, err := nebula.NewConnectionPool(
		[]nebula.HostAddress{hostAddress},
		poolConfig,
		nebula.DefaultLogger{},
	)
	if err != nil {
		return fmt.Errorf("连接池创建失败: %v", err)
	}
	defer pool.Close()

	// 3. 获取 session
	session, err := pool.GetSession("root", "nebula")
	if err != nil {
		return fmt.Errorf("获取session失败: %v", err)
	}
	defer session.Release()

	// 4. 检查服务状态
	fmt.Println("检查 Nebula 服务状态...")

	// 4.1 检查 Graph 服务
	_, err = session.Execute("SHOW HOSTS;")
	if err != nil {
		return fmt.Errorf("Graph 服务检查失败: %v", err)
	}

	// 4.2 检查 Storage 服务
	_, err = session.Execute("SHOW HOSTS STORAGE;")
	if err != nil {
		return fmt.Errorf("Storage 服务检查失败: %v", err)
	}

	// 4.3 检查 Meta 服务
	_, err = session.Execute("SHOW HOSTS META;")
	if err != nil {
		return fmt.Errorf("Meta 服务检查失败: %v", err)
	}

	// 5. 检查现有空间
	fmt.Println("\n检查现有空间...")
	execNGQL(session, "SHOW SPACES;")

	// 6. 如果 demo_space 已存在，先删除
	execNGQL(session, "DROP SPACE IF EXISTS demo_space;")
	time.Sleep(2 * time.Second) // 等待空间删除完成

	// 7. 创建新空间
	fmt.Println("\n创建新空间...")
	execNGQL(session, "CREATE SPACE demo_space(vid_type=INT64, partition_num=1, replica_factor=1);")
	time.Sleep(5 * time.Second)
	// 等待空间创建完成
	//waitForSpaceCreation(session, "demo_space")

	// 8. 使用新空间
	execNGQL(session, "USE demo_space;")

	// 9. 创建标签和边类型（通用三元组）
	fmt.Println("\n创建三元组模型的 Schema...")
	execNGQL(session, "CREATE TAG entity(name string);")
	execNGQL(session, "CREATE EDGE relation(name string);")
	time.Sleep(5 * time.Second) // 等待 schema 同步

	fmt.Println("\n创建并构建索引...")
	execNGQL(session, "CREATE TAG INDEX entity_index ON entity(name(20));") // 添加字符串长度限制
	time.Sleep(10 * time.Second)
	execNGQL(session, "REBUILD TAG INDEX entity_index;")
	time.Sleep(10 * time.Second) // 等待索引构建完成

	// 10. 插入三元组数据
	fmt.Println("\n插入三元组测试数据...")
	execNGQL(session, "INSERT VERTEX entity(name) VALUES 1001:('张三'), 1002:('李四'), 1003:('王五');")
	execNGQL(session, "INSERT EDGE relation(name) VALUES 1001->1002:('朋友'), 1002->1003:('老师');")

	// 11. 基于实体名查询其属性与相关实体
	fmt.Println("\n查询实体“张三”的属性与关系：")
	if err := queryByEntityName(session, "张三"); err != nil {
		return err // 或 return fmt.Errorf("查询失败: %w", err)
	}
	return nil
}

// 查询指定实体名对应的所有属性和关联实体
func queryByEntityName(session *nebula.Session, name string) error {
	// 第一步：根据名称查 VID
	stmt := fmt.Sprintf("LOOKUP ON entity WHERE entity.name == '%s' YIELD id(vertex) AS vid;", name)
	resp, err := session.Execute(stmt)
	if err != nil || !resp.IsSucceed() || len(resp.GetRows()) == 0 {
		return fmt.Errorf("查找实体 %s 失败: %v", name, err)
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
	//fmt.Println("执行成功")
	return resp, nil
}
func printQueryResult(resp *nebula.ResultSet) {
	// 打印结果
	colNames := resp.GetColNames()
	rows := resp.GetRows()
	if len(rows) == 0 {
		fmt.Println("无数据")
		return
	}

	// 打印列名
	fmt.Println("列名:", colNames)

	// 打印数据行
	for i, row := range rows {
		fmt.Printf("第%d行: ", i+1)
		for _, val := range row.Values {
			switch {
			case val.IsSetSVal(): // 字符串类型
				fmt.Printf("%s ", string(val.GetSVal()))
			case val.IsSetIVal(): // 整数类型
				fmt.Printf("%d ", val.GetIVal())
			case val.IsSetFVal(): // 浮点数类型
				fmt.Printf("%.2f ", val.GetFVal())
			case val.IsSetBVal(): // 布尔类型
				fmt.Printf("%t ", val.GetBVal())
			default: // 其他类型
				fmt.Printf("%v ", val)
			}
		}
		fmt.Println()
	}
}
