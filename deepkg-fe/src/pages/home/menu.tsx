import {
    ApartmentOutlined,
    AppstoreOutlined,
    BookOutlined,
    DashboardOutlined,
    DatabaseOutlined,
    FileSearchOutlined,
    FileTextOutlined,
    FolderOpenOutlined,
    SettingOutlined,
    ShareAltOutlined,
    TagsOutlined,
    UserOutlined
} from "@ant-design/icons";

import ChatIcon from '../../assets/chat.svg';

const ImageIcon = ( src : any) => (
  <img src={src} alt="icon" style={{ width: 14, height: 14, marginRight: 8 }} />
);


export const menuList = [
    {
        key: '/dashboard',
        icon: <DashboardOutlined />,
        label: '首页',
    },
    {
        key: '/document_center',
        icon: <BookOutlined />,
        label: '文档管理',
        children: [
            {
                key: '/document_dir',
                icon: <FolderOpenOutlined />,
                label: '目录管理',
            },
            {
                key: '/document',
                icon: <FileTextOutlined />,
                label: '文件管理',
            },
        ]
    },
    {
        key: '/knowledge',
        icon: <DatabaseOutlined />,
        label: '知识管理',
        children: [
            {
                key: '/workspace',
                icon: <AppstoreOutlined />,
                label: '图空间管理',
            },
            {
                key: '/ontology',
                icon: <TagsOutlined />,
                label: '本体管理',
            },
            {
                key: '/triple',
                icon: <ShareAltOutlined />,
                label: '关系管理',
            },
            {
                key: '/extract_task',
                icon: <FileSearchOutlined />,
                label: '非结构化抽取',
            },
        ]
    },
    {
        key: '/knowledge_infer',
        icon: ImageIcon(ChatIcon),
        label: '知识推理',
    },
    {
        key: '/system',
        icon: <SettingOutlined />,
        label: '系统管理',
        children: [
            {
                key: '/org',
                icon: <ApartmentOutlined />,
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