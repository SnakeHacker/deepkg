import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button, Form, Input, Modal, Pagination, Popconfirm, Table } from "antd";
import { PlusOutlined, RollbackOutlined } from "@ant-design/icons";
import type { SchemaOntologyProp } from "../../model/schema_ontology_prop";
import { CreateSchemaOntologyProp, DeleteSchemaOntologyProps, ListSchemaOntologyProp, UpdateSchemaOntologyProp } from "../../service/schema_ontology_prop";

const OntologyPropPage: React.FC = () => {

    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });
    const [total, setTotal] = useState(0);
    const [props, setProps] = useState<SchemaOntologyProp[]>([]);
    const [ontologyID, setOntologyID] = useState(0);
    const [propID, setPropID] = useState(0);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [form] = Form.useForm();

    useEffect(() => {
        const hash = window.location.hash;
        const hashParts = hash.split('?');
        if (hashParts.length > 1) {
            const searchParams = new URLSearchParams(hashParts[1]);
            const ontologyId = searchParams.get('ontology_id');
            setOntologyID(Number(ontologyId))
        } else {
            console.warn('URL hash 中不包含查询参数');
        }

    }, []);

    useEffect(() => {
        ontologyID > 0 && listOntologyProps()
    }, [ontologyID]);


    const listOntologyProps = async () => {
        const res = await ListSchemaOntologyProp({
            ontology_id: ontologyID,
            page_size: pagination.pageSize,
            page_number: pagination.current
        });
        setTotal(res.total)
        setProps(res.schema_ontology_props);
    };


    const handleCreateOntologyPropOk = async () => {
        const values = await form.validateFields();
        const { prop_name, prop_desc } = values;
        try {
            setIsModalOpen(false);
            if (propID > 0) {
                const res = await UpdateSchemaOntologyProp({
                    schema_ontology_prop: {
                        id: propID,
                        ontology_id: ontologyID,
                        prop_name: prop_name,
                        prop_desc: prop_desc
                    },
                });
                console.log(res)
                form.resetFields();
                setPropID(0);
                listOntologyProps();
            } else {

                const res = await CreateSchemaOntologyProp({
                    schema_ontology_prop: {
                        ontology_id: ontologyID,
                        prop_name: prop_name,
                        prop_desc: prop_desc
                    }
                });
                console.log(res)
                form.resetFields();
                setPropID(0);
                listOntologyProps();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateOntologyProp = () => {
        setIsModalOpen(false);
        form.resetFields();
        setOntologyID(0);
    }

    const deleteOntologyProps = async (id: number) => {
        const res = await DeleteSchemaOntologyProps({ ids: [id] });
        console.log(res)
        listOntologyProps();
    }

    const handleEdit = (record: SchemaOntologyProp) => {
        setPropID(record.id!);
        form.setFieldsValue({
            prop_name: record.prop_name,
            prop_desc: record.prop_desc
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
          title: '属性名称',
          dataIndex: 'prop_name',
          key: 'prop_name',
          width: '20%',
        },
        {
            title: '属性描述',
            dataIndex: 'prop_desc',
            key: 'prop_desc',
            width: '40%',
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
            render: (_: any, record: SchemaOntologyProp) => (
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
                        description="确认删除属性?"
                        onConfirm={() => deleteOntologyProps(record.id!)}
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

                <Button
                    type="default"
                    icon={<RollbackOutlined />}
                    onClick={() => {
                        window.history.back();
                    }}
                    style={{ marginRight: '10px'}}
                >
                    返回
                </Button>
                <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={() => {
                        setPropID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                    }}
                >
                    新建属性
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        propID > 0 ? '编辑空间' : `新建空间`
                    }
                    open={isModalOpen}
                    onOk={handleCreateOntologyPropOk}
                    onCancel={handleCancelCreateOntologyProp}
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
                            label="属性名称"
                            name="prop_name"
                            rules={[{ required: true, message: '请输入属性名称' }]}
                        >
                            <Input
                                style={{'width': '100%'}}
                                placeholder="请输入本体名称"
                            />
                        </Form.Item>

                        <Form.Item
                            label="属性描述"
                            name="prop_desc"
                        >
                            <Input.TextArea
                                style={{'width': '100%'}}
                                placeholder="请输入属性描述"
                            />
                        </Form.Item>
                    </Form>
                </Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={props}
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

export default OntologyPropPage;
