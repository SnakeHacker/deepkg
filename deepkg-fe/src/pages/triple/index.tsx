import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button, Form, Input, Modal, Pagination, Popconfirm, Select, Table } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import type { SchemaTriple } from "../../model/schema_triple";
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { CreateSchemaTriple, DeleteSchemaTriples, ListSchemaTriple, UpdateSchemaTriple } from "../../service/schema_triple";
import { ListKnowledgeGraphWorkspace } from "../../service/workspace";
import type { SchemaOntology } from "../../model/schema_ontology";
import { ListSchemaOntology } from "../../service/schema_ontology";

const TriplePage: React.FC = () => {
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });
    const [total, setTotal] = useState(0);
    const [triples, setTriples] = useState<SchemaTriple[]>([]);
    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [workspaceID, setWorkspaceID] = useState(0);
    const [ontologies, setOntologies] = useState<SchemaOntology[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [tripleID, setTripleID] = useState(0);
    const [form] = Form.useForm();

    useEffect(() => {
        listWorkspaces()
    }, []);

    useEffect(() => {
        if (workspaceID > 0){
            listTriples()
            listOntologies()
        }
    }, [workspaceID]);

    const listTriples = async () => {
        const res = await ListSchemaTriple({
            work_space_id: workspaceID,
            page_size: pagination.pageSize,
            page_number: pagination.current
        });
        setTotal(res.total)
        setTriples(res.schema_triples);
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

    const listOntologies = async () => {
        const res = await ListSchemaOntology({
            work_space_id: workspaceID,
            page_size: 10,
            page_number: -1
        });
        setOntologies(res.schema_ontologys);
    };

    const handleCreateTripleOk = async () => {
        const values = await form.validateFields();
        const { source_ontology_id, target_ontology_id, relationship } = values;
        try {
            setIsModalOpen(false);
            if (tripleID > 0) {
                const res = await UpdateSchemaTriple({
                    schema_triple: {
                        id: tripleID,
                        work_space_id: workspaceID,
                        source_ontology_id: source_ontology_id,
                        target_ontology_id: target_ontology_id,
                        relationship: relationship,
                    },
                });
                console.log(res)
                form.resetFields();
                setTripleID(0);
                listTriples();
            } else {

                const res = await CreateSchemaTriple({
                    schema_triple: {
                        work_space_id: workspaceID,
                        source_ontology_id: source_ontology_id,
                        target_ontology_id: target_ontology_id,
                        relationship: relationship,
                    }
                });
                console.log(res)
                form.resetFields();
                setTripleID(0);
                listTriples();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateTriple = () => {
        setIsModalOpen(false);
        form.resetFields();
        setTripleID(0);
    }

    const deleteTriple = async (id: number) => {
        const res = await DeleteSchemaTriples({ ids: [id] });
        console.log(res)
        listTriples();
    }

    const handleEdit = (record: SchemaTriple) => {
        setTripleID(record.id!);
        form.setFieldsValue({
            relationship: record.relationship,
            source_ontology_id: record.source_ontology_id === 0 ? '': record.source_ontology_id,
            target_ontology_id: record.target_ontology_id === 0 ? '': record.target_ontology_id,
        });
        setIsModalOpen(true);
    }

    const columns = [
        {
          title: 'ID',
          dataIndex: 'id',
          key: 'id',
          width: '10%',
        },
        {
            title: '起始节点',
            dataIndex: 'source_ontology_name',
            key: 'source_ontology_name',
            width: '20%',
        },
        {
          title: '关系',
          dataIndex: 'relationship',
          key: 'relationship',
          width: '20%',
        },
        {
          title: '目标节点',
          dataIndex: 'target_ontology_name',
          key: 'target_ontology_name',
          width: '20%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            width: 350,
        },
        {
            title: '操作',
            key: 'action',
            width: 200,
            render: (_: any, record: SchemaTriple) => (
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
                        description="确认删除关系?"
                        onConfirm={() => deleteTriple(record.id!)}
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
                        setTripleID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                        if (workspaceID > 0){
                            form.setFieldValue("work_space_id", workspaceID)
                        }
                    }}
                >
                    新建关系
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        tripleID > 0 ? '编辑关系' : `新建关系`
                    }
                    open={isModalOpen}
                    onOk={handleCreateTripleOk}
                    onCancel={handleCancelCreateTriple}
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
                            label="起始节点"
                            name="source_ontology_id"
                            rules={[{ required: true, message: '请选择起始节点' }]}
                        >
                            <Select
                                style={{'width': '100%'}}
                                placeholder="起始节点"
                                disabled={ontologies.length === 0}
                                options={[
                                    ...ontologies.map((ontology) => (
                                        {key: ontology.id, label: ontology.ontology_name, value: ontology.id}
                                    )),
                                ]}
                            >
                            </Select>
                        </Form.Item>

                        <Form.Item
                            label="关系名称"
                            name="relationship"
                            rules={[{ required: true, message: '请输入关系名称' }]}
                        >
                            <Input
                                style={{'width': '100%'}}
                                placeholder="请输入关系名称"
                            />
                        </Form.Item>

                        <Form.Item
                            label="目标节点"
                            name="target_ontology_id"
                            rules={[{ required: true, message: '请选择目标节点' }]}
                        >
                            <Select
                                style={{'width': '100%'}}
                                placeholder="目标节点"
                                disabled={ontologies.length === 0}
                                options={[
                                    ...ontologies.map((ontology) => (
                                        {key: ontology.id, label: ontology.ontology_name, value: ontology.id}
                                    )),
                                ]}
                            >
                            </Select>
                        </Form.Item>
                    </Form>
                </Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={triples}
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

export default TriplePage;
