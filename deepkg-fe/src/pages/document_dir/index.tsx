import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import type { DocumentDir } from "../../model/document_dir";
import { Button, Form, Input, Modal, Select, Table, Popconfirm } from "antd";
import type { TableColumnsType, TableProps } from 'antd';
import { CreateDocumentDir, DeleteDocumentDirs, ListDocumentDir, UpdateDocumentDir } from "../../service/document_dir";
import { PlusOutlined } from "@ant-design/icons";

const { Option } = Select;

export interface DataType {
    key: string;
    dir_name: string;
    parent_id?: number;
    id: number;
    children?: DataType[];
    remark?: string;
    created_at?: string;
}
const DocumentDirPage: React.FC = () => {


    const [dirs, setDirs] = useState<DataType[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [dirID, setDirID] = useState(0);
    const [form] = Form.useForm();

    useEffect(() => {
        listDirs();
    }, []);

    const listDirs = async () => {
        const res = await ListDocumentDir();
        console.log(res)
        // 定义一个递归函数来处理多级目录
        const processDir = (dir: DataType): DataType => {
            return {
                dir_name: dir.dir_name,
                parent_id: dir.parent_id,
                id: dir.id,
                key: dir.id.toString(),
                created_at: dir.created_at,
                children: dir.children && dir.children.length > 0
                    ? dir.children.map((child: DataType) => processDir(child))
                    : undefined
            };
        };

        const document_dirs = (res.document_dirs || []).map((dir: DataType) => processDir(dir));

        setDirs(document_dirs);
    };

    const handleCreateDirOk = async () => {
        const values = await form.validateFields();
        const { dir_name, parent_id } = values;
        try {
            setIsModalOpen(false);
            if (dirID > 0) {
                const res = await UpdateDocumentDir({
                    document_dir: {
                        id: dirID,
                        parent_id: parent_id,
                        dir_name: dir_name,
                    },
                });
                console.log(res)
                form.resetFields();
                setDirID(0);
                listDirs();
            } else {

                const res = await CreateDocumentDir({
                    document_dir: {
                        parent_id: parent_id,
                        dir_name: dir_name,
                    }
                });
                console.log(res)
                form.resetFields();
                setDirID(0);
                listDirs();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateDocumentDir = () => {
        setIsModalOpen(false);
        form.resetFields();
        setDirID(0);
    }

    const deleteDir = async (id: number) => {
        const res = await DeleteDocumentDirs({ ids: [id] });
        console.log(res)
        listDirs();
    }

    const handleEdit = (record: DocumentDir) => {
        setDirID(record.id!);
        form.setFieldsValue({
            dir_name: record.dir_name,
            parent_id: record.parent_id,
            remark: record.remark
        });
        setIsModalOpen(true);
    }



    const columns: TableColumnsType<DataType> = [
        {
          title: 'ID',
          dataIndex: 'id',
          key: 'id',
          width: '10%',
        },
        {
          title: '目录名称',
          dataIndex: 'dir_name',
          key: 'dir_name',
          width: '12%',
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
            render: (_: any, record: DocumentDir) => (
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
                        description="确认删除目录?"
                        onConfirm={() => deleteDir(record.id!)}
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
                        setDirID(0);
                        form.resetFields();
                        setIsModalOpen(true)
                    }}
                >
                    新建目录
                </Button>
            </div>
            <div className={styles.body}>
                <Modal
                    title={
                        dirID > 0 ? '编辑目录' : `新建目录`
                    }
                    open={isModalOpen}
                    onOk={handleCreateDirOk}
                    onCancel={handleCancelCreateDocumentDir}
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
                            label="父级目录"
                            name="parent_id"
                        >

                            <Select
                                style={{'width': '100%'}}
                                placeholder="请选择父级目录"
                                disabled={dirs.length === 0}
                            >
                                {dirs.map((dir) => (
                                <Option key={dir.id} value={dir.id}>
                                    {dir.dir_name}
                                </Option>
                                ))}
                            </Select>
                        </Form.Item>
                        <Form.Item
                            label="目录名称"
                            name="dir_name"
                            rules={[{ required: true, message: '请输入目录名称' }]}
                        >
                            <Input
                                style={{'width': '100%'}}
                                placeholder="请输入目录名称"
                            />
                        </Form.Item>
                        <Form.Item
                            label="备注"
                            name="remark"
                        >
                            <Input.TextArea
                                style={{'width': '100%'}}
                                placeholder="请输入备注"
                            />
                        </Form.Item>



                    </Form>
                </Modal>

                <Table
                    style={{ width: '100%' }}
                    dataSource={dirs}
                    columns={columns}
                    pagination={false}
                    rowKey="id"
                />
            </div>

            <div className={styles.footer}>
            </div>
        </div>
    )
};

export default DocumentDirPage;
