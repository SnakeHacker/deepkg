import React, { useEffect, useState } from "react";
import { Layout, Avatar, Dropdown, message, Menu } from "antd";
import type { MenuProps } from 'antd';
import { MenuFoldOutlined, MenuUnfoldOutlined, UserOutlined } from "@ant-design/icons";
import { useNavigate, Outlet, useLocation } from "react-router-dom";
import styles from "./index.module.less";
import { Logout } from "../../service/session";
import { getCurrentUser } from "../../service/auth";
import type { MessageInfo } from "../../store";
import { useStore } from "../../store";
import logoImg from "../../assets/headerLogo.svg"
import BgSVG from '../../assets/background.png';
import { ROLE_ADMIN } from "../../utils/const";
import { menuList } from "./menu";

const { Header, Sider } = Layout;
type MenuItem = Required<MenuProps>['items'][number];

const HomePage: React.FC = () => {
    const { successMsg, errorMsg, warningMsg } = useStore() as MessageInfo;
     

    const [messageApi, contextHolder] = message.useMessage();

    //   const location = useLocation();

    useEffect(() => {
        successMsg !== "" && messageApi.success(successMsg)
    }, [successMsg]);

    useEffect(() => {
        errorMsg !== "" && messageApi.error(errorMsg)
    }, [errorMsg]);

    useEffect(() => {
        warningMsg !== "" && messageApi.warning(warningMsg)
    }, [warningMsg]);


    const navigate = useNavigate();
    const location = useLocation();

    const [collapsed, setCollapsed] = useState(false);
    const [selectedKeys, setSelectedKeys] = useState<string[]>(['']);
    const [menus , setMenus] = useState<MenuItem[]>(menuList);
    const [isAdmin, setIsAdmin] = useState(false);

    useEffect(() => {
        const currentUser = getCurrentUser();
        setIsAdmin(currentUser?.role === ROLE_ADMIN);
        if (currentUser?.role !== ROLE_ADMIN) {

            const filteredMenus = menuList.filter(item => item.key!== '/system');
            setMenus(filteredMenus);
        }
    }, []);

    const userMenu: MenuProps['items'] = [
        {
            key: 'logout',
            label: '退出登录',
            onClick: () => {
                startLogout();
                localStorage.clear();
                navigate(`/`)
            }
        },
    ];

    const startLogout = async () => {
        const res = await Logout()
        return res
    };

    const toMenu = (key: string) => {
        navigate(key);
        // setSelectedKeys([key]);
    }

    return (
        <Layout className={styles.container}>
            {contextHolder}
            <img src={BgSVG} className={styles.backgroundImage} />
            <Sider
                trigger={null}
                collapsible
                collapsed={collapsed}
                breakpoint="lg"
                collapsedWidth={80}
                style={{backgroundColor: 'white' }}
            >

                <div className={styles.logoContainer}>
                    <img src={logoImg} alt="Logo" className={styles.logo} />

                    <div className={styles.title}>知识图谱平台</div>
                </div>

                <Menu
                    theme="light"
                    mode="inline"
                    selectedKeys={[location.pathname]}
                    //   selectedKeys={selectedKeys}
                    //defaultSelectedKeys={[location.pathname]}
                    // 默认全部展开，获取所有菜单项的 key 作为 openKeys
                    defaultOpenKeys={(menus|| []).map(item => item!.key as string)}
                    onClick={({ key }) => toMenu(key)}
                    items={menus}
                    style={{ height: 'calc(100vh - 90px)', overflowY: 'auto' }}
                />

                {React.createElement(collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                    className: styles.trigger,
                    onClick: () => setCollapsed(!collapsed),
                })}
            </Sider>

            <Layout>
                <Header className={styles.header}>

                    <Dropdown menu={{ items: isAdmin ? [...userMenu] : [...userMenu] }}>
                        <div className={styles.userInfo}>
                            <Avatar className={styles.avatar} icon={<UserOutlined />} />
                            <span style={{ marginLeft: 8 }}>
                                {getCurrentUser()?.username || '未登录'}
                            </span>
                        </div>
                    </Dropdown>
                </Header>

                <Outlet />


            </Layout>
        </Layout>
    )
};

export default HomePage;
