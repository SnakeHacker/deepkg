import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { CreateKnowledgeGraphWorkspace, DeleteKnowledgeGraphWorkspaces, ListKnowledgeGraphWorkspace, UpdateKnowledgeGraphWorkspace } from "../../service/workspace";
import { Button, Form, Input, Modal, Pagination, Popconfirm, Table } from "antd";
import { PlusOutlined } from "@ant-design/icons";

const WorkspacePage: React.FC = () => {
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });
    const [total, setTotal] = useState(0);
    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [workspaceID, setWorkspaceID] = useState(0);
    const [form] = Form.useForm();

    useEffect(() => {
        listWorkspaces();
    }, [pagination.current, pagination.pageSize]);

    const listWorkspaces = async () => {
        const res = await ListKnowledgeGraphWorkspace({
            page_size: pagination.pageSize,
            page_number: pagination.current
        });
        setTotal(res.total);
        setWorkspaces(res.knowledge_graph_workspaces || []);
    };

    const handleCreateWorkspaceOk = async () => {
        const values = await form.validateFields();
        const { knowledge_graph_workspace_name } = values;
        try {
            setIsModalOpen(false);
            if (workspaceID > 0) {
                await UpdateKnowledgeGraphWorkspace({
                    knowledge_graph_workspace: {
                        id: workspaceID,
                        knowledge_graph_workspace_name: knowledge_graph_workspace_name,
                    },
                });
                form.resetFields();
                setWorkspaceID(0);
                listWorkspaces();
            } else {
                await CreateKnowledgeGraphWorkspace({
                    knowledge_graph_workspace: {
                        knowledge_graph_workspace_name: knowledge_graph_workspace_name,
                    }
                });
                form.resetFields();
                setWorkspaceID(0);
                listWorkspaces();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    };

    const handleCancelCreateWorkspace = () => {
        setIsModalOpen(false);
        form.resetFields();
        setWorkspaceID(0);
    };

    const deleteWorkspace = async (id: number) => {
        await DeleteKnowledgeGraphWorkspaces({ ids: [id] });
        listWorkspaces();
    };

    const handleEdit = (record: KnowledgeGraphWorkspace) => {
        setWorkspaceID(record.id!);
        form.setFieldsValue({
            knowledge_graph_workspace_name: record.knowledge_graph_workspace_name,
        });
        setIsModalOpen(true);
    };

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id',
            width: 80,
            ellipsis: true,
        },
        {
            title: '空间名称',
            dataIndex: 'knowledge_graph_workspace_name',
            key: 'knowledge_graph_workspace_name',
            ellipsis: true,
            render: (text: string) => (
                <span title={text}>{text}</span>
            ),
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            width: 180,
            ellipsis: true,
        },
        {
            title: '操作',
            key: 'action',
            width: 160,
            render: (_: any, record: KnowledgeGraphWorkspace) => (
                <div key={record.id}>
                    <Button
                        style={{ marginRight: '8px' }}
                        onClick={() => handleEdit(record)}
                        size='small'
                    >
                        编辑
                    </Button>
                    <Popconfirm
                        title={""}
                        description="确认删除空间?"
                        onConfirm={() => deleteWorkspace(record.id!)}
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

    const handlePageChange = (page: number, pageSize?: number) => {
        setPagination(prev => ({
            ...prev,
            current: page,
            pageSize: pageSize || prev.pageSize
        }));
    };

    return (
        <div className={styles.container} style={{ backgroundImage: `url(${BgSVG})` }}>
            <div className={styles.header}>
                <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={() => {
                        setWorkspaceID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                    }}
                >
                    新建空间
                </Button>
            </div>
            <div className={styles.body}>
                <div className={styles.tableContainer}>
                    <Table
                        dataSource={workspaces}
                        columns={columns}
                        pagination={false}
                        rowKey="id"
                        //scroll={{ y: 'calc(100vh - 300px)', x: 'max-content' }}
                    />
                </div>
                <div className={styles.footer}>
                    <Pagination
                        current={pagination.current}
                        pageSize={pagination.pageSize}
                        total={total}
                        onChange={handlePageChange}
                        showSizeChanger
                        showQuickJumper
                        showTotal={(total, range) => `第 ${range[0]}-${range[1]} 条/共 ${total} 条`}
                        pageSizeOptions={['10', '20', '50', '100']}
                    />
                </div>
            </div>
            <Modal
                title={workspaceID > 0 ? '编辑空间' : `新建空间`}
                open={isModalOpen}
                onOk={handleCreateWorkspaceOk}
                onCancel={handleCancelCreateWorkspace}
                okText="确定"
                cancelText="取消"
                width={450}
            >
                <Form
                    form={form}
                    name="userForm"
                    labelAlign='left'
                    labelCol={{ span: 5 }}
                >
                    <Form.Item
                        label="空间名称"
                        name="knowledge_graph_workspace_name"
                        rules={[{ required: true, message: '请输入空间名称' }]}
                    >
                        <Input
                            style={{ 'width': '100%' }}
                            placeholder="请输入空间名称"
                        />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
};

export default WorkspacePage;
