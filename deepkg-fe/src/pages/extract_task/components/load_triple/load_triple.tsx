import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import { Modal, Pagination, Table , type TableProps} from "antd";

import { useStore, type LoadTriple } from "../../../../store";
import type { SchemaTriple } from "../../../../model/schema_triple";
import { ListSchemaTriple } from "../../../../service/schema_triple";

interface LoadTripleProps {
    visible: boolean;
    onCancel: Function;
    onOk: Function;
    workspaceID: number;
}

const LoadTripleComponent: React.FC<LoadTripleProps> = ({ visible, onCancel, onOk, workspaceID }) => {
    const [selectedRowKeysState, setSelectedRowKeysState] = useState<React.Key[]>([]);
    const { setTripleList } = useStore() as LoadTriple;
    const [pagination, setPagination] = useState({
        current: -1,
        pageSize: 10,
    });

    const [triples, setTriples] = useState<SchemaTriple[]>([]);
    const [total, setTotal] = useState(0);

    useEffect(() => {
        visible && workspaceID>0 && listTriples();
    }, [visible, workspaceID]);

    const listTriples = async () => {
        const res = await ListSchemaTriple(
            {
                work_space_id: workspaceID,
                page_size: pagination.pageSize,
                page_number: pagination.current
            }
        );
        console.log(res);

        setTriples(res.schema_triples);
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
          title: '起始节点',
          dataIndex: 'source_ontology_name',
          key: 'source_ontology_name',
          width: '20%',
        },
        {
            title: '关系名称',
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
    ];



    const rowSelection: TableProps<SchemaTriple>['rowSelection'] = {
        selectedRowKeys: selectedRowKeysState,
        onChange: (selectedRowKeys: React.Key[], selectedRows: SchemaTriple[]) => {
            console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
            setTripleList(selectedRows);
            setSelectedRowKeysState(selectedRowKeys);
        },
        getCheckboxProps: (record: SchemaTriple) => ({
            id: record.id?.toString(),
        }),
    };

    const clearSelection = () => {
        setSelectedRowKeysState([]);
    };

    const handleCancel = () => {
        setTriples([]);
        onCancel();
        clearSelection();
    }

    const handleOk = () => {
        setTriples([]);
        clearSelection();
        onOk();
    };``

    return (
        <Modal
            title={"导入三元组"}
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
        </Modal>
    )
};

export default LoadTripleComponent;
