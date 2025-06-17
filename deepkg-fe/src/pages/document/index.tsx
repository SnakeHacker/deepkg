import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { ListDocumentDir } from "../../service/document_dir";
import { Button, Form, Input, Modal, Select, Table, Tree, type TreeDataNode, Popconfirm, Upload, type UploadProps, Pagination } from "antd";
import { CreateDocument, DeleteDocuments, ListDocument, UpdateDocument } from "../../service/document";
import { InboxOutlined, PlusOutlined } from "@ant-design/icons";
import type { Document } from "../../model/document";
import request from "../../utils/req";
import  { type MessageInfo , useStore}from "../../store";

const { Option } = Select;
const { Dragger } = Upload;

export interface DataType {
    key: string;
    dir_name: string;
    title: string;
    id: number;
    children?: DataType[];
}
const DocumentPage: React.FC = () => {
    const { success, error } = useStore() as MessageInfo;
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });

    const [total, setTotal] = useState(0);
    const [dirs, setDirs] = useState<DataType[]>([]);
    const [dirID, setDirID] = useState(0);
    const [docs, setDocs] = useState<Document[]>([]);
    const [docID, setDocID] = useState(0);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [form] = Form.useForm();

    useEffect(() => {
        listDirs();
    }, []);

    useEffect(() => {
        listDocuments();
    }, [dirID]);

    const listDirs = async () => {
        const res = await ListDocumentDir();
        console.log(res)
        // 定义一个递归函数来处理多级目录
        const processDir = (dir: DataType): DataType => {
            return {
                id: dir.id,
                dir_name: dir.dir_name,
                title: dir.dir_name,
                key: dir.id.toString(),
                children: dir.children && dir.children.length > 0
                    ? dir.children.map((child: DataType) => processDir(child))
                    : []
            };
        };

        const document_dirs = (res.document_dirs || []).map((dir: DataType) => processDir(dir));

        setDirs(document_dirs);
    };

    const listDocuments = async () => {
        const res = await ListDocument(
            {
                dir_id: dirID,
                page_size: pagination.pageSize,
                page_number: pagination.current
            }
        );
        console.log(res);

        setDocs(res.documents);
        setTotal(res.total);
    };

    const handleSelectDir= (selectedKeys: React.Key[])=>{
        const selectedKeyAsInt = parseInt(selectedKeys[0] as string, 10);
        setDirID(selectedKeyAsInt);
    }

    const handleCreateDocOk = async () => {
        const values = await form.validateFields();
        const { doc_name, doc_path, doc_desc, dir_id } = values;
        try {
            setIsModalOpen(false);
            if (docID > 0) {
                const res = await UpdateDocument({
                    document: {
                        id: docID,
                        dir_id: dir_id,
                        doc_name: doc_name,
                        doc_path: doc_path,
                        doc_desc: doc_desc,
                    },
                });
                console.log(res)
                form.resetFields();
                setDocID(0);
                listDocuments();
            } else {

                const res = await CreateDocument({
                    document: {
                        dir_id: dir_id,
                        doc_name: doc_name,
                        doc_path: doc_path,
                        doc_desc: doc_desc,
                    }
                });
                console.log(res)
                form.resetFields();
                setDocID(0);
                listDocuments();
            }
        } catch (errorInfo) {
            console.log('Failed:', errorInfo);
        }
    }

    const handleCancelCreateDoc = () => {
        setIsModalOpen(false);
        form.resetFields();
        setDocID(0);
    }

    const deleteDoc = async (id: number) => {
        const res = await DeleteDocuments({ ids: [id] });
        console.log(res)
        listDirs();
    }

    const handleEdit = (record: Document) => {
        setDocID(record.id!);
        form.setFieldsValue({
            doc_name: record.doc_name,
            doc_path: record.doc_path,
            doc_desc: record.doc_desc,
            dir_id: record.dir_id === 0 ? '': record.dir_id,
        });
        setIsModalOpen(true);
    }

    const uploadFile = (options: any) => {
        const { file, onSuccess, onProgress, onError } = options

        const formData = new FormData()
        formData.append('file', file)

        request
          .post('/file/upload', formData, {
            onUploadProgress: ({ total, loaded }) => {
              onProgress({ percent: Math.round((loaded / (total!)) * 100).toFixed(2) }, file)
            },
          })
          .then((res: any) => {
            onSuccess({ ...res, name: file.name, status: 'done' }, file)
          })
          .catch(onError)

        return {
          abort() {
            console.log('上传进度中止')
          },
        }
    };
    const onUploadChange = (info: any) => {
        if (info.file.status === 'done') {
            success(`${info.file.name} 文件上传成功`);
            console.log(info.file.response)
            form.setFieldValue('doc_path',info.file.response.host+'/'+info.file.response.object_id)
            form.setFieldValue('doc_name', info.file.name)

        } else if (info.file.status === 'error') {
            error(`${info.file.name} 文件上传失败`);
        }
    };

    const uploadProps: UploadProps = {
        multiple: false,
        maxCount: 1, // 限制上传数量。当为 1 时，始终用最新上传的文件代替当前文件
        customRequest:uploadFile,
        onChange: onUploadChange,
        onRemove: (file) => {
            console.log('移除文件:', file);
            form.setFieldValue('doc_path', null);
        },
        accept: ".txt,.pdf,.docx,.md,.xsl,.xslx",
        beforeUpload: (file) => {
            console.log('上传前的文件:', file);
            return true; // 返回 true 允许上传，返回 false 阻止上传
        }
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
          title: '文件名称',
          dataIndex: 'doc_name',
          key: 'doc_name',
          width: '60%',
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
            render: (_: any, record: Document) => (
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
                        description="确认删除文件?"
                        onConfirm={() => deleteDoc(record.id!)}
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
            <div className={styles.left}>
                <Tree
                    onSelect={handleSelectDir}
                    treeData={dirs.map((dir) => ({
                        key: dir.key !== undefined ? String(dir.key) : '',
                        title: dir.dir_name,
                        children: dir.children,
                    } as TreeDataNode))}
                />
            </div>
            <div className={styles.right}>
                <div className={styles.header}>
                    <Button
                        type="primary"
                        icon={<PlusOutlined />}
                        onClick={() => {
                            setDocID(0);
                            form.resetFields();
                            setIsModalOpen(true)
                        }}
                    >
                        新建文档
                    </Button>
                </div>
                <div className={styles.body}>
                    <Modal
                        title={
                            dirID > 0 ? '编辑文档' : `新建文档`
                        }
                        open={isModalOpen}
                        onOk={handleCreateDocOk}
                        onCancel={handleCancelCreateDoc}
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
                                label="父级目录"
                                name="dir_id"
                            >

                                <Select
                                    style={{'width': '100%'}}
                                    placeholder="所属目录"
                                    disabled={dirs.length === 0}
                                    options={[
                                        ...dirs.map((dir) => (
                                            {key: dir.id, label: dir.dir_name, value: Number(dir.key)}
                                        )),
                                    ]}
                                >
                                </Select>
                            </Form.Item>
                            <Form.Item
                                label="文件名称"
                                name="doc_name"
                                rules={[{ required: true, message: '请输入文件名称' }]}
                            >
                                <Input
                                    style={{'width': '100%'}}
                                    placeholder="请输入文件名称"
                                />
                            </Form.Item>
                            <Form.Item
                                label="文件路径"
                                name="doc_path"
                                // rules={[{ required: true, message: '请上传文件' }]}
                            >
                                <Dragger
                                    style={{ backgroundColor: '#ffffff' }}
                                    {...uploadProps}
                                >
                                <p className="ant-upload-drag-icon">
                                    <InboxOutlined />
                                </p>
                                <p className="ant-upload-text">点击或拖拽文件到此区域上传</p>
                                <p className="ant-upload-hint">
                                    支持格式: .txt, .pdf, .docx, .xsl, .xslx
                                </p>
                                </Dragger>
                            </Form.Item>
                            <Form.Item
                                label="文件描述"
                                name="doc_desc"
                            >
                                <Input.TextArea
                                    style={{'width': '100%'}}
                                    placeholder="请输入文件描述"
                                />
                            </Form.Item>



                        </Form>
                    </Modal>

                    <div className={styles.tableContainer}>
                    <Table
                        style={{ width: '100%' }}
                        dataSource={docs}
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
            </div>


        </div>
    )
};

export default DocumentPage;
