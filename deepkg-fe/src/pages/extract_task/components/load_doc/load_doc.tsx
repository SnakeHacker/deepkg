import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import { ListDocument } from "../../../../service/document";
import { Modal, Pagination, Table , type TableProps} from "antd";
import { useStore, type LoadDoc } from "../../../../store";
import type { Document } from "../../../../model/document";

interface LoadDocProps {
    visible: boolean;
    onCancel: Function;
    onOk: Function;
}

const LoadDocComponent: React.FC<LoadDocProps> = ({ visible, onCancel, onOk }) => {
    const [selectedRowKeysState, setSelectedRowKeysState] = useState<React.Key[]>([]);
    const { setDocList } = useStore() as LoadDoc;
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
    });

    const [docs, setDocs] = useState<Document[]>([]);
    const [total, setTotal] = useState(0);


    useEffect(() => {
        visible && listDocuments();
    }, [visible]);

    const listDocuments = async () => {
        const res = await ListDocument(
            {
                page_size: pagination.pageSize,
                page_number: pagination.current
            }
        );
        console.log(res);

        setDocs(res.documents);
        setTotal(res.total);
    };

    const handlePageChange = (page: number, pageSize?: number) => {
        setPagination(prev => ({
            ...prev,
            current: page,
            pageSize: pageSize || prev.pageSize
        }));
    };

    const columns = [
        {
            title: 'id',
            dataIndex: 'id',
            key: 'id',
            width: '5%',
        },
        {
          title: '文件名称',
          dataIndex: 'doc_name',
          key: 'doc_name',
          width: '20%',
        },
        {
            title: '文件描述',
            dataIndex: 'doc_desp',
            key: 'doc_desp',
            width: '40%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            key: 'created_at',
            width: 350,
        },
    ];



const rowSelection: TableProps<Document>['rowSelection'] = {
    selectedRowKeys: selectedRowKeysState,
    onChange: (selectedRowKeys: React.Key[], selectedRows: Document[]) => {
        console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
        setDocList(selectedRows);
        setSelectedRowKeysState(selectedRowKeys);
    },
    getCheckboxProps: (record: Document) => ({
        name: record.doc_name,
    }),
};

    const clearSelection = () => {
        setSelectedRowKeysState([]);
    };

    const handleCancel = () => {
        setDocs([]);
        onCancel();
        clearSelection();
    }

    const handleOk = () => {
        setDocs([]);
        clearSelection();
        onOk();
    };``

    return (
        <Modal
            title={"导入文件"}
            open={visible}
            onOk={()=>{handleOk()}}
            onCancel={()=>{handleCancel()}}
            okText="确定"
            cancelText="取消"
            width={1000}
        >
        <div className={styles.container}
        >

            <div className={styles.body}>

                <Table
                    style={{ width: '100%' }}
                    rowSelection={{ type: 'checkbox', ...rowSelection }}
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
        </Modal>
    )
};

export default LoadDocComponent;
