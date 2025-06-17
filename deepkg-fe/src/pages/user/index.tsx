import React, { useEffect, useState } from 'react';
import { PlusOutlined, RollbackOutlined } from '@ant-design/icons';
import { Button, Form, Input, Modal, Pagination, Popconfirm, Select, Switch, Table } from 'antd';
import type { User } from '../../model/user';
import styles from './index.module.less';
import type { Organization } from '../../model/organization';
import { CreateUser, DeleteUsers, ListUser, UpdateUser } from '../../service/user';
import { ListOrg } from '../../service/organization';
import { GetPublicKey } from '../../service/session';
import {JSEncrypt} from 'jsencrypt';
import { ROLE_ADMIN, ROLE_USER, USER_DISABLED, USER_ENABLED } from '../../utils/const';
import BgSVG from '../../assets/bg.png';

const UserListPage: React.FC = () => {
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });

    const [total, setTotal] = useState(0);
    const [users, setUsers] = useState<User[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [userID, setUserID] = useState(0);
    const [form] = Form.useForm();
    const [orgs, setOrgs] = useState<Organization[]>([]);

    useEffect(() => {
        getOrgList();
    }, []);

    useEffect(() => {
        listUsers();
    }, [pagination]);

    const getOrgList = async () => {
        const res = await ListOrg({
            page_size: 1000,
            page_number: -1
        });
        console.log(res)
        setOrgs(res.organizations)
    };

    const listUsers = async () => {
        const res = await ListUser({
            page_size: pagination.pageSize,
            page_number: pagination.current
        });
        console.log(res)

        setTotal(res.total)
        setUsers(res.users)
    };

    const handlePageChange = (page: number, pageSize?: number) => {
        setPagination(prev => ({
            ...prev,
            current: page,
            pageSize: pageSize || prev.pageSize
        }));
    };

    const handleCreateUserOk = async () => {
        const values = await form.validateFields();
        const { password, enable } = values;
        try {
            setIsModalOpen(false);
            if (userID > 0) {
                const res = await updateUser({
                    id: userID,
                    enable: enable ? USER_ENABLED : USER_DISABLED,
                    ...values
                });
                console.log(res)
                form.resetFields();
                setUserID(0);
                listUsers();
            } else {
                const publicKeyRes = await GetPublicKey();
                const encryptor = new JSEncrypt();
                encryptor.setPublicKey(publicKeyRes.public_key);
                const encryptedPwd = encryptor.encrypt(password.trim());

                const res = await createUser({
                        ...values,
                        enable: enable ? USER_ENABLED : USER_DISABLED,
                        password: encryptedPwd
                });
                console.log(res)
                form.resetFields();
                setUserID(0);
                listUsers();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateUser = () => {
        setIsModalOpen(false);
        form.resetFields();
        setUserID(0);
    }

    const createUser = async (userData: any) => {
        const res = await CreateUser({
            user: {
                ...userData
            }
        })
        return res
    }

    const updateUser = async (userData: any) => {
        const res = await UpdateUser({
            user: {
                ...userData
            }
        })
        return res
    }

    const deleteUser = async (id: number) => {
        const res = await DeleteUsers({ ids: [id] });
        console.log(res)
        listUsers();
    }

    const handleEdit = (record: User) => {
        setUserID(record.id!);
        form.setFieldsValue({
            username: record.username,
            account: record.account,
            role: record.role,
            org_id: record.org_id,
        });
        setIsModalOpen(true);
    }

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id',
            width: 50,
        },
        {
            title: '用户名称',
            dataIndex: 'username',
            key: 'username',
            width: 200,
        },
        {
            title: '账号',
            dataIndex: 'account',
            key: 'account',
            width: 200,
        },
        {
            title: '组织',
            dataIndex: 'org_id',
            key: 'org_id',
            width: 350,
            render: (_: any, record: User) => {
                const org = orgs.find(org => org.id === record.org_id);
                return org ? org.org_name : '未知';
            },
        },
        {
            title: '角色',
            dataIndex: 'role',
            key: 'role',
            width: 200,
            render: (_: any, record: User) => {
                return record.role==ROLE_ADMIN? '管理员' : '普通用户';
            },
        },
        {
            title: '启用/停用',
            dataIndex: 'enable',
            key: 'enable',
            width: 200,
            render: (_: any, record: User) => {
                return (
                    <Switch
                        checked={record.enable == 1}
                        onChange={async (checked) => {
                            try {
                                await updateUser({
                                    ...record,
                                    id: record.id,
                                    enable: checked ? USER_ENABLED : USER_DISABLED,
                                });
                                listUsers();
                            } catch (error) {
                                console.error('更新用户启用状态失败:', error);
                            }
                        }}
                    />
                );
            },
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            width: 350,
        },
        {
            title: '更新时间',
            dataIndex: 'updated_at',
            key: 'updated_at',
            width: 350,
        },
        {
            title: '操作',
            key: 'action',
            width: 200,
            render: (_: any, record: User) => (
                <div key={record.id}>
                    <Button
                        style={{ marginRight: '10px' }}
                        onClick={() => handleEdit(record)}
                        size='small'
                    >
                        编辑
                    </Button>

                    <Popconfirm
                        title={""}
                        description="确认是否删除用户?"
                        onConfirm={() => deleteUser(record.id!)}
                        onCancel={() => { }}
                        okText="确认"
                        cancelText="取消"
                    >
                        <Button
                            size='small'
                            danger
                        >
                            删除
                        </Button>
                    </Popconfirm>

                </div>
            ),
        },
    ];

    return (
        <div className={styles.container}
            style={{
                backgroundImage: `url(${BgSVG})`,
            }}
        >

            <div className={styles.header}>
                <Button
                    style={{ marginRight: '10px' }}
                    type="default"
                    icon={<RollbackOutlined />}
                    onClick={() => {
                        window.history.back();
                    }}
                    >
                    返回
                </Button>
                <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={() => {
                        setUserID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                    }}
                >
                    新建用户
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        userID > 0 ? '编辑用户' : `新建用户`
                    }
                    open={isModalOpen}
                    onOk={handleCreateUserOk}
                    onCancel={handleCancelCreateUser}
                    okText="确定"
                    cancelText="取消"
                    width={420}
                >
                    <Form
                        form={form}
                        name="userForm"
                        labelAlign='left'
                        labelCol={{ span: 4 }}
                    >
                        <Form.Item
                            label="用户名"
                            name="username"
                            rules={[{ required: true, message: '请输入用户名' }]}
                        >
                            <Input
                                style={{'width': 300}}
                                placeholder="请输入用户名称"
                            />
                        </Form.Item>

                        <Form.Item
                            label="账号"
                            name="account"
                            rules={[{ required: true, message: '请输入账号' }]}
                        >
                            <Input
                                style={{'width': 300}}
                                placeholder="请输入用户账号"
                            />
                        </Form.Item>

                        {userID == 0 &&
                        <Form.Item
                            label="密码"
                            name="password"
                            rules={[{ required: true, message: '请输入密码' }]}
                        >
                            <Input.Password
                                style={{'width': 300}}
                                placeholder="请输入用户密码"
                            />
                        </Form.Item>
                        }

                        <Form.Item
                            label="角色"
                            name="role"
                            rules={[{ required: true, message: '请选择角色' }]}
                        >
                            <Select
                             style={{'width': 300}}
                            >
                                <Select.Option value={ROLE_USER}>普通用户</Select.Option>
                                <Select.Option value={ROLE_ADMIN}>超级管理员</Select.Option>
                            </Select>
                        </Form.Item>

                        <Form.Item
                            label="组织"
                            name="org_id"
                            rules={[{ required: true, message: '请选择组织' }]}
                        >
                            <Select
                             style={{'width': 300}}
                            >
                                {(orgs || []).map(org => (
                                    <Select.Option key={org.id} value={org.id}>
                                        {org.org_name}
                                    </Select.Option>
                                ))}
                            </Select>
                        </Form.Item>

                        <Form.Item
                            label="启用"
                            name="enable"
                            rules={[{ required: true, message: '请选择组织' }]}
                        >
                            <Switch checked={form.getFieldValue('enable')} />
                        </Form.Item>
                    </Form>
                </Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={users}
                    columns={columns}
                    pagination={false}
                    rowKey="id"
                />
            </div>

            <div className={styles.footer}>
                <Pagination
                    current={pagination.current}
                    pageSize={pagination.pageSize}
                    total={total}
                    onChange={handlePageChange}
                />
            </div>
        </div>
    );
};

export default UserListPage;