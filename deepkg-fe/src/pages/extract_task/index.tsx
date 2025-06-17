import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button, Form, Input, Modal, Select, Table, Tree, type TreeDataNode, Popconfirm, Pagination } from "antd";
import type { ExtractTask } from "../../model/extract_task";
import { CreateExtractTask, DeleteExtractTasks, ListExtractTask, UpdateExtractTask } from "../../service/extract_task";
import { PlusOutlined } from "@ant-design/icons";
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { ListKnowledgeGraphWorkspace } from "../../service/workspace";

const ExtractTaskPage: React.FC = () => {
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });

    const [total, setTotal] = useState(0);

    const [tasks, setTasks] = useState<ExtractTask[]>([]);
    const [taskID, setTaskID] = useState(0);
    const [workspaceID, setWorkspaceID] = useState(0);
    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isLoadDocModalOpen, setIsLoadDocModalOpen] = useState(false);
    const [isLoadTripleModalOpen, setIsLoadTripleModalOpen] = useState(false);

    const [form] = Form.useForm();

    useEffect(() => {
        listWorkspaces();
    }, []);

    useEffect(() => {
        listTasks();
    }, []);

    const listTasks = async () => {
        const res = await ListExtractTask(
            {
                work_space_id: workspaceID,
                page_size: pagination.pageSize,
                page_number: pagination.current
            }
        );
        console.log(res);

        setTotal(res.total);
        setTasks(res.extract_tasks);
    };

    const listWorkspaces = async () => {
        const res = await ListKnowledgeGraphWorkspace({
            page_size: 10,
            page_number: -1
        });
        setWorkspaces(res.knowledge_graph_workspaces);
        if (res.knowledge_graph_workspaces.length > 0){
            setWorkspaceID(res.knowledge_graph_workspaces[0].id)
        }
    };

    const handleCreateTaskOk = async () => {
        const values = await form.validateFields();
        const { doc_ids, triple_ids, work_space_id, task_name, remark } = values;
        try {
            setIsModalOpen(false);
            if (taskID > 0) {
                const res = await UpdateExtractTask({
                    extract_task: {
                        id: taskID,
                        task_name: task_name,
                        remark: remark,
                        work_space_id: work_space_id,
                        docs: [],
                        triples: [],
                    },
                });
                console.log(res)
                form.resetFields();
                setTaskID(0);
                listTasks();
            } else {
                const res = await CreateExtractTask({
                    extract_task: {
                        task_name: task_name,
                        remark: remark,
                        work_space_id: work_space_id,
                        docs: [],
                        triples: [],
                    }
                });
                console.log(res)
                form.resetFields();
                setTaskID(0);
                listTasks();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateTask = () => {
        setIsModalOpen(false);
        form.resetFields();
        setTaskID(0);
    }

    const handleCancelLoadDocModal = () => {
        setIsLoadDocModalOpen(false);
    }

    const handleCancelLoadTripleModal = () => {
        setIsLoadTripleModalOpen(false);
    }

    const deleteTasks = async (id: number) => {
        const res = await DeleteExtractTasks({ ids: [id] });
        console.log(res)
        listTasks();
    }

    const handleEdit = (record: ExtractTask) => {
        setTaskID(record.id!);
        form.setFieldsValue({
            task_name: record.task_name,
            remark: record.remark,
            // TODO(mickey)
            // docs: record.docs,
            // triples: record.triples,
            work_space_id: record.work_space_id=== 0 ? '': record.work_space_id,
        });
        setIsModalOpen(true);
    }

    const handlePageChange = (page: number, pageSize?: number) => {
        setPagination(prev => ({
            ...prev,
            current: page,
            pageSize: pageSize || prev.pageSize
        }));
    };

    const columns = [
        {
          title: 'ID',
          dataIndex: 'id',
          key: 'id',
          width: '10%',
        },
        {
          title: '任务名称',
          dataIndex: 'task_name',
          key: 'task_name',
          width: '60%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
        },
        {
            title: '操作',
            key: 'action',
            width: 200,
            render: (_: any, record: ExtractTask) => (
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
                        description="确认删除任务?"
                        onConfirm={() => deleteTasks(record.id!)}
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
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={() => {
                        setTaskID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                    }}
                >
                    新建任务
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        taskID > 0 ? '编辑非结构化抽取任务' : `新建非结构化抽取任务`
                    }
                    open={isModalOpen}
                    onOk={handleCreateTaskOk}
                    onCancel={handleCancelCreateTask}
                    okText="确定"
                    cancelText="取消"
                    width={600}
                >
                    <Form
                        form={form}
                        name="userForm"
                        labelAlign='left'
                        labelCol={{ span: 5 }}
                    >

                        <Form.Item
                            label="任务名称"
                            name="task_name"
                            rules={[{ required: true, message: '请输入任务名称' }]}
                        >
                            <Input
                                style={{'width': '100%'}}
                                placeholder="请输入任务名称"
                            />
                        </Form.Item>

                        <Form.Item
                            label="导入文件"
                            name="doc_ids"
                            rules={[{ required: true, message: '请选择文件' }]}
                        >
                            <Button onClick={()=>{setIsLoadDocModalOpen(true)}}>
                              导入文件
                            </Button>
                        </Form.Item>

                        <Form.Item
                            label="导入三元组"
                            name="triple_ids"
                            rules={[{ required: true, message: '请选择三元组' }]}
                        >
                            <Button onClick={()=>{setIsLoadTripleModalOpen(true)}}>
                                导入三元组
                            </Button>
                        </Form.Item>

                        <Form.Item
                            label="备注"
                            name="remark"
                        >
                            <Input.TextArea
                                style={{'width': '100%'}}
                            />
                        </Form.Item>
                    </Form>
                </Modal>

                <Modal
                    title={"导入文件"}
                    open={isLoadDocModalOpen}
                    // onOk={handleCreateTaskOk}
                    onCancel={handleCancelLoadDocModal}
                    okText="确定"
                    cancelText="取消"
                    width={800}
                ></Modal>

<Modal
                    title={"导入三元组"}
                    open={isLoadTripleModalOpen}
                    // onOk={handleCreateTaskOk}
                    onCancel={handleCancelLoadTripleModal}
                    okText="确定"
                    cancelText="取消"
                    width={800}
                ></Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={tasks}
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
    )
};

export default ExtractTaskPage;
