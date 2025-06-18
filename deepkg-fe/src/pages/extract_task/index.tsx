import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button, Form, Input, Modal, Table, Popconfirm, Pagination, Select } from "antd";
import type { ExtractTask } from "../../model/extract_task";
import { CreateExtractTask, DeleteExtractTasks, GetExtractTask, ListExtractTask, UpdateExtractTask } from "../../service/extract_task";
import { PlusOutlined } from "@ant-design/icons";
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { ListKnowledgeGraphWorkspace } from "../../service/workspace";
import LoadDocComponent from "./components/load_doc/load_doc";
import LoadTripleComponent from "./components/load_triple/load_triple";
import { useStore, type LoadDoc, type LoadTriple } from "../../store";

const EXTRACT_TASK_STATUS_WAITING = 1;
const EXTRACT_TASK_STATUS_RUNNING = 2;
const EXTRACT_TASK_STATUS_FAILED = 3;
const EXTRACT_TASK_STATUS_SUCCESSED = 4;

const ExtractTaskPage: React.FC = () => {
    const { docList, removeDocListItem, clearDocList, setDocList  } = useStore() as LoadDoc;
    const { tripleList, removeTripleListItem, clearTripleList, setTripleList  } = useStore() as LoadTriple;

    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });

    const [total, setTotal] = useState(0);

    const [tasks, setTasks] = useState<ExtractTask[]>([]);
    const [taskID, setTaskID] = useState(0);
    const [workspaceID, setWorkspaceID] = useState(0);
    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [curWorkspaceID, setCurWorkspaceID] = useState(0);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isLoadDocModalOpen, setIsLoadDocModalOpen] = useState(false);
    const [isLoadTripleModalOpen, setIsLoadTripleModalOpen] = useState(false);

    const [form] = Form.useForm();

    useEffect(() => {
        listWorkspaces();
    }, []);

    useEffect(() => {
        workspaceID >0 && listTasks();
    }, [workspaceID]);

    useEffect(() => {
        console.log(docList)
    }, [docList]);

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
            setWorkspaceID(res.knowledge_graph_workspaces[res.knowledge_graph_workspaces.length-1].id)
        }
    };

    const getExtractTask = async(tid: number)=>{
        const res = await GetExtractTask({
            id: tid,
        });

        const curTask = res.extract_task
        setDocList(curTask.docs)
        setTripleList(curTask.triples)


    }

    const handleCreateTaskOk = async () => {
        const values = await  form.validateFields(['task_name', 'work_space_id', 'remark', 'doc_ids', 'triple_ids'])
        const { work_space_id, task_name, remark } = values;
        const newDocList = docList.map(item => ({ id: item.id }));
        const newTripleList = tripleList.map(item => ({ id: item.id }));

        try {
            setIsModalOpen(false);
            if (taskID > 0) {
                const res = await UpdateExtractTask({
                    extract_task: {
                        id: taskID,
                        task_name: task_name,
                        remark: remark,
                        work_space_id: work_space_id,
                        docs: newDocList,
                        triples: newTripleList,
                    },
                });
                console.log(res)
                form.resetFields();
                setTaskID(0);
                setCurWorkspaceID(0);
                listTasks();
            } else {
                const res = await CreateExtractTask({
                    extract_task: {
                        task_name: task_name,
                        remark: remark,
                        work_space_id: work_space_id,
                        docs: newDocList,
                        triples: newTripleList,
                    }
                });
                console.log(res)
                form.resetFields();
                setTaskID(0);
                setCurWorkspaceID(0);
                listTasks();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleLoadDocsOk = async () => {
        setIsLoadDocModalOpen(false);
    }

    const handleLoadTriplesOk = async () => {
        setIsLoadTripleModalOpen(false);
    }

    const handleCancelCreateTask = () => {
        setIsModalOpen(false);
        form.resetFields();
        setTaskID(0);
        setCurWorkspaceID(0);
        clearDocList();
        clearTripleList();
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

    const handleEdit = async (record: ExtractTask) => {
        setTaskID(record.id!);
        await getExtractTask(record.id!);
        form.setFieldsValue({
            task_name: record.task_name,
            remark: record.remark,
            work_space_id: record.work_space_id=== 0 ? '': record.work_space_id,
        });
        setCurWorkspaceID(record.work_space_id);
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
          width: '30%',
        },
        {
            title: '发布状态',
            dataIndex: 'published',
            key: 'published',
            width: '15%',
            render: (_: any, record: ExtractTask) => (
                <div key={record.id}>
                    {record.published ? '已发布' : '未发布'}
                </div>
            )
        },
        {
            title: '任务状态',
            dataIndex: 'task_status',
            key: 'task_status',
            width: '10%',
            render: (_: any, record: ExtractTask) => (
                <div key={record.id}>
                    {record.task_status == EXTRACT_TASK_STATUS_WAITING ? '等待' :
                        record.task_status == EXTRACT_TASK_STATUS_RUNNING ? '运行' :
                            record.task_status == EXTRACT_TASK_STATUS_FAILED ? '失败' :
                                record.task_status == EXTRACT_TASK_STATUS_SUCCESSED ? '成功' : '未知'}
                </div>
            )
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
                <Select
                    style={{'width': '200px', 'marginRight': '10px'}}
                    placeholder="工作空间"
                    disabled={workspaces.length === 0}
                    onSelect={(value) => setWorkspaceID(value)}
                    value={workspaceID}
                    options={[
                        ...workspaces.map((workspaces) => (
                            {key: workspaces.id, label: workspaces.knowledge_graph_workspace_name, value: workspaces.id}
                        )),
                    ]}
                />
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
                    width={1000}
                >
                    <Form
                        form={form}
                        name="extract_task_form"
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
                            label="工作空间"
                            name="work_space_id"
                            rules={[{ required: true, message: '请选择工作空间' }]}
                        >
                            <Select
                                style={{'width': '100%'}}
                                placeholder="工作空间"
                                disabled={workspaces.length === 0}
                                onChange={(value)=>setCurWorkspaceID(value)}
                                options={[
                                    ...workspaces.map((workspaces) => (
                                        {key: workspaces.id, label: workspaces.knowledge_graph_workspace_name, value: workspaces.id}
                                    )),
                                ]}
                            >
                            </Select>
                        </Form.Item>

                        <Form.Item
                            label="导入文件"
                            name="doc_ids"
                            rules={[{
                                required: true,
                                validator: () => {
                                    if (docList.length == 0) {
                                      return Promise.reject('请导入文件');
                                    }
                                    return Promise.resolve();
                                }
                            }]}
                        >
                            <Button
                                style={{'marginBottom': '15px'}}
                                onClick={()=>{setIsLoadDocModalOpen(true)}}
                                disabled={taskID !=0}
                            >
                              导入文件
                            </Button>

                            <Table
                                dataSource={docList}
                                style={{'height': '230px', 'overflowY': 'auto'}}
                                columns={[
                                    {
                                        title: 'ID',
                                        dataIndex: 'id',
                                        key: 'id',
                                        width: '50px'
                                    },
                                    {
                                        title: '名称',
                                        dataIndex: 'doc_name',
                                        key: 'doc_name',
                                    },
                                    {
                                        title: '操作',
                                        key: 'action',
                                        width: '100px',
                                        render: (_, record) => (
                                            <Button
                                                danger
                                                size='small'
                                                onClick={() => {
                                                    removeDocListItem(record.id)
                                                }}
                                                disabled={taskID !=0}
                                            >
                                                删除
                                            </Button>
                                        ),
                                    },
                                ]}
                                rowKey="id"
                            />
                        </Form.Item>

                        <Form.Item
                            label="导入三元组"
                            name="triple_ids"
                            rules={[{
                                required: true,
                                validator: () => {
                                    if (tripleList.length == 0) {
                                      return Promise.reject('请导入三元组');
                                    }
                                    return Promise.resolve();
                                }
                            }]}
                        >
                            <Button
                                style={{'marginBottom': '15px'}}
                                onClick={()=>{setIsLoadTripleModalOpen(true)}}
                                disabled={curWorkspaceID == 0 || taskID !=0}

                            >
                                导入三元组
                            </Button>

                            <Table
                                dataSource={tripleList}
                                style={{'height': '230px', 'overflowY': 'auto'}}
                                columns={[
                                    {
                                        title: 'ID',
                                        dataIndex: 'id',
                                        key: 'id',
                                        width: '50px'
                                    },
                                    {
                                        title: '起始节点',
                                        dataIndex: 'source_ontology_name',
                                        key: 'source_ontology_name',
                                    },
                                    {
                                        title: '关系名称',
                                        dataIndex: 'relationship',
                                        key: 'relationship',
                                    },
                                    {
                                        title: '目标节点',
                                        dataIndex: 'target_ontology_name',
                                        key: 'target_ontology_name',
                                    },
                                    {
                                        title: '操作',
                                        key: 'action',
                                        width: '100px',
                                        render: (_, record) => (
                                            <Button
                                                danger
                                                size='small'
                                                onClick={() => {
                                                    removeTripleListItem(record.id)
                                                }}
                                                disabled={taskID !=0 }
                                            >
                                                删除
                                            </Button>
                                        ),
                                    },
                                ]}
                                rowKey="id"
                            />
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


                <LoadDocComponent
                    visible={isLoadDocModalOpen}
                    onOk={handleLoadDocsOk}
                    onCancel={handleCancelLoadDocModal}
                />

                <LoadTripleComponent
                    visible={isLoadTripleModalOpen}
                    onOk={handleLoadTriplesOk}
                    onCancel={handleCancelLoadTripleModal}
                    workspaceID={curWorkspaceID}
                />

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
