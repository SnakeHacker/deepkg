import { RobotOutlined, DashboardOutlined, BankOutlined, UserOutlined, SettingOutlined} from "@ant-design/icons";

export const menuList = [
    {
        key: '/dashboard',
        icon: <DashboardOutlined />,
        label: '首页',
    },
    {
        key: '/document_center',
        icon: <RobotOutlined />,
        label: '文档中心',
        children: [
            {
                key: '/document_center/dir',
                icon: <BankOutlined />,
                label: '目录管理',
            },
            {
                key: '/document_center/document',
                icon: <BankOutlined />,
                label: '文件管理',
            },
        ]
    },
    {
        key: '/extraction',
        icon: <RobotOutlined />,
        label: '知识抽取',
        children: [
            {
                key: '/extraction/workspace',
                icon: <BankOutlined />,
                label: '图空间管理',
            },
            {
                key: '/extraction/ontology',
                icon: <BankOutlined />,
                label: '本体管理',
            },
            {
                key: '/extraction/relationship',
                icon: <BankOutlined />,
                label: '关系管理',
            },
            {
                key: '/user',
                icon: <UserOutlined />,
                label: '非结构化抽取',
            },
        ]
    },
    {
        key: '/system',
        icon: <SettingOutlined />,
        label: '系统管理',
        children: [
            {
                key: '/org',
                icon: <BankOutlined />,
                label: '组织管理',
            },
            {
                key: '/user',
                icon: <UserOutlined />,
                label: '用户管理',
            },
        ]
    }

]