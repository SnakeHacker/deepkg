import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button, Form, Input, Modal, Pagination, Popconfirm, Select, Table } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import type { SchemaOntology } from "../../model/schema_ontology";
import { CreateSchemaOntology, DeleteSchemaOntologys, ListSchemaOntology, UpdateSchemaOntology } from "../../service/schema_ontology";
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { ListKnowledgeGraphWorkspace } from "../../service/workspace";
import { useNavigate } from "react-router-dom";

const OntologyPage: React.FC = () => {
    const navigate = useNavigate();
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });
    const [total, setTotal] = useState(0);
    const [ontologies, setOntologies] = useState<SchemaOntology[]>([]);
    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [workspaceID, setWorkspaceID] = useState(0);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [ontologyID, setOntologyID] = useState(0);
    const [form] = Form.useForm();

    useEffect(() => {
        listWorkspaces()
    }, []);

    useEffect(() => {
        if (workspaceID > 0){
            listOntologies()
        }
    }, [workspaceID]);

    const listOntologies = async () => {
        const res = await ListSchemaOntology({
            work_space_id: workspaceID,
            page_size: pagination.pageSize,
            page_number: pagination.current
        });
        setTotal(res.total)
        setOntologies(res.schema_ontologys);
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

    const handleCreateOntologyOk = async () => {
        const values = await form.validateFields();
        const { work_space_id, ontology_name, ontology_desc } = values;
        try {
            setIsModalOpen(false);
            if (ontologyID > 0) {
                const res = await UpdateSchemaOntology({
                    schema_ontology: {
                        id: ontologyID,
                        work_space_id: work_space_id,
                        ontology_name: ontology_name,
                        ontology_desc: ontology_desc
                    },
                });
                console.log(res)
                form.resetFields();
                setOntologyID(0);
                listOntologies();
            } else {

                const res = await CreateSchemaOntology({
                    schema_ontology: {
                        work_space_id: work_space_id,
                        ontology_name: ontology_name,
                        ontology_desc: ontology_desc
                    }
                });
                console.log(res)
                form.resetFields();
                setOntologyID(0);
                listOntologies();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateOntology = () => {
        setIsModalOpen(false);
        form.resetFields();
        setOntologyID(0);
    }

    const deleteOntologies = async (id: number) => {
        const res = await DeleteSchemaOntologys({ ids: [id] });
        console.log(res)
        listOntologies();
    }

    const handleEdit = (record: SchemaOntology) => {
        setOntologyID(record.id!);
        form.setFieldsValue({
            ontology_name: record.ontology_name,
            ontology_desc: record.ontology_desc,
            work_space_id: record.work_space_id === 0 ? '': record.work_space_id,
        });
        setIsModalOpen(true);
    }

    const handleEditProp = (record: SchemaOntology) => {
        console.log(record.id)
        navigate(`/ontology_prop?ontology_id=${record.id}`);
    }

    const columns = [
        {
          title: 'ID',
          dataIndex: 'id',
          key: 'id',
          width: '10%',
        },
        {
          title: '本体名称',
          dataIndex: 'ontology_name',
          key: 'ontology_name',
          width: '20%',
        },
        {
            title: '本体描述',
            dataIndex: 'ontology_desc',
            key: 'ontology_desc',
            width: '30%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            width: 300,
        },
        {
            title: '操作',
            key: 'action',
            width: 200,
            render: (_: any, record: SchemaOntology) => (
                <div key={record.id}>
                    <Button
                        style={{ marginRight: '10px' }}
                        onClick={() => handleEdit(record)}
                        size='small'
                    >
                        编辑
                    </Button>

                    <Button
                        style={{ marginRight: '10px' }}
                        onClick={() => handleEditProp(record)}
                        size='small'
                    >
                        属性
                    </Button>

                    <Popconfirm
                        title={""}
                        description="确认删除本体?"
                        onConfirm={() => deleteOntologies(record.id!)}
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
                        setOntologyID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                        if (workspaceID > 0){
                            form.setFieldValue("work_space_id", workspaceID)
                        }
                    }}
                >
                    新建本体
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        ontologyID > 0 ? '编辑本体' : `新建本体`
                    }
                    open={isModalOpen}
                    onOk={handleCreateOntologyOk}
                    onCancel={handleCancelCreateOntology}
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
                            label="工作空间"
                            name="work_space_id"
                            rules={[{ required: true, message: '请选择工作空间' }]}
                        >
                            <Select
                                style={{'width': '100%'}}
                                placeholder="工作空间"
                                disabled={workspaces.length === 0}
                                options={[
                                    ...workspaces.map((workspaces) => (
                                        {key: workspaces.id, label: workspaces.knowledge_graph_workspace_name, value: workspaces.id}
                                    )),
                                ]}
                                onSelect={(value)=> setWorkspaceID(value)}
                            >
                            </Select>
                        </Form.Item>

                        <Form.Item
                            label="本体名称"
                            name="ontology_name"
                            rules={[{ required: true, message: '请输入本体名称' }]}
                        >
                            <Input
                                style={{'width': '100%'}}
                                placeholder="请输入本体名称"
                            />
                        </Form.Item>

                        <Form.Item
                            label="本体描述"
                            name="ontology_desc"
                        >
                            <Input.TextArea
                                style={{'width': '100%'}}
                                placeholder="请输入本体描述"
                            />
                        </Form.Item>
                    </Form>
                </Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={ontologies}
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

export default OntologyPage;
