# DeepKG

[English](README.md) | 简体中文

DeepKG是一个基于大模型的多模态知识图谱平台，旨在简化知识图谱的构建、管理和推理过程。

## 系统展示

| 目录管理 | 文件管理 | 图空间管理 |
| --- | --- | --- |
| ![目录管理](assets/screenshots/directory.png "目录管理") | ![文件管理](assets/screenshots/document.png "文件管理") | !["图空间管理"](assets/screenshots/workspace.png "图空间管理") |

| 本体管理 | 关系管理 | 非结构化抽取 |
| --- | --- | --- |
| ![本体管理](assets/screenshots/ontology.png "本体管理") | ![关系管理](assets/screenshots/relationship.png "关系管理")| ![非结构化抽取](assets/screenshots/extraction.png "非结构化抽取") |

| 组织管理 | 用户管理 |
| --- | --- |
| ![组织管理](assets/screenshots/organization.png "组织管理") | ![用户管理](assets/screenshots/user.png "用户管理") |

## 时序图

![时序图](assets/sequence_diagram_zh.png "时序图")

## 依赖

- Go 1.20
- Node.js
- pnpm
- Docker

## 安装
```shell
# 1. 使用Docker部署数据库等服务
cd admin/deploy
docker compose -f docker-compose.yml up -d

# 2. 使用Docker部署Nebula
cd nebula
docker compose -f docker-compose-lite.yml up -d

# 3. 初始化数据库
cd ../../
make init_ddl

# 4. 启动后端服务
make run

# 5. 安装前端依赖
cd ../deepkg-fe
make install

# 6. 启动前端服务
make run
```

## 贡献者

- [@SnakeHacker](https://github.com/SnakeHacker)
- [@IsshikiSenn](https://github.com/IsshikiSenn)
- [@chenmiao8563](https://github.com/chenmiao8563)
- [@wuwuwukai](https://github.com/wuwuwukai)

## TODO
- [x] 系统架构设计、中间件选型
- [ ] 用户账号、组织、角色、菜单、鉴权体系构建
    - [x] 用户增删改查: Mickey
    - [x] 组织增删改查: 腾博
    - [ ] 角色增删改查
    - [ ] 菜单增删改查
    - [ ] 鉴权体系构建
- [ ] Nebula集成
    - [x] Nebula部署
    - [ ] Nebula接口验证: 陈淼
- [ ] 基于非结构化数据的图谱构建
- [ ] 知识图谱实体融合、消歧
- [ ] 图谱发布
- [ ] 基于图谱的推理
- [ ] 多模态知识图谱构建
- [ ] 基于多模态知识图谱的推理
- [ ] API文档: 吴凯
