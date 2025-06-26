import React, { useEffect, useState } from 'react';
import { Table, Tag, Card, message } from 'antd';
import { ListExtractTask } from '../../../../service/extract_task';
import { ListKnowledgeGraphWorkspace } from '../../../../service/workspace';
import styles from './index.module.less';

const EXTRACT_TASK_STATUS_WAITING = 1;
const EXTRACT_TASK_STATUS_RUNNING = 2;
const EXTRACT_TASK_STATUS_FAILED = 3;
const EXTRACT_TASK_STATUS_SUCCESSED = 4;

const ExtractTaskList: React.FC = () => {
  const [data, setData] = useState<any[]>([]);
  const [loading, setLoading] = useState(false);
  const [pageSize, setPageSize] = useState(4);
  const [pageNumber, setPageNumber] = useState(1);
  const [total, setTotal] = useState(0);

  // 获取所有工作空间下的所有任务
  const fetchAllTasks = async () => {
    setLoading(true);
    try {
      const wsRes = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
      const workspaces = wsRes.knowledge_graph_workspaces || [];
      let allTasks: any[] = [];
      for (const ws of workspaces) {
        const res = await ListExtractTask({
          work_space_id: ws.id,
          page_size: 1000, // 拉取所有任务
          page_number: 1,
        });
        if (Array.isArray(res.extract_tasks)) {
          allTasks = allTasks.concat(res.extract_tasks.map((task: any) => ({
            ...task,
            workspace_name: ws.knowledge_graph_workspace_name,
          })));
        }
      }
      setTotal(allTasks.length);
      setData(allTasks);
    } catch (error) {
      message.error('获取所有抽取任务失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAllTasks();
    const interval = setInterval(fetchAllTasks, 30000);
    return () => clearInterval(interval);
  }, []);

  // 分页处理
  const pagedData = data.slice((pageNumber - 1) * pageSize, pageNumber * pageSize);

  const columns = [
    { title: '任务ID', dataIndex: 'id', key: 'id' },
    { title: '任务名称', dataIndex: 'task_name', key: 'task_name' },
    { title: '工作空间', dataIndex: 'workspace_name', key: 'workspace_name' },
    {
      title: '发布状态',
      dataIndex: 'published',
      key: 'published',
      render: (published: boolean | number) => {
        const status = typeof published === 'boolean' ? (published ? 1 : 0) : Number(published);
        const label = status === 1 ? '已发布' : '未发布';
        const color = status === 1 ? 'geekblue' : 'default';
        return <Tag color={color}>{label}</Tag>;
      },
    },
    {
      title: '任务状态',
      dataIndex: 'task_status',
      key: 'task_status',
      render: (status: number) => {
        let label = '未知';
        let color = 'default';
        switch (status) {
          case EXTRACT_TASK_STATUS_WAITING:
            label = '等待';
            color = 'orange';
            break;
          case EXTRACT_TASK_STATUS_RUNNING:
            label = '运行';
            color = 'blue';
            break;
          case EXTRACT_TASK_STATUS_FAILED:
            label = '失败';
            color = 'red';
            break;
          case EXTRACT_TASK_STATUS_SUCCESSED:
            label = '成功';
            color = 'green';
            break;
        }
        return <Tag color={color}>{label}</Tag>;
      },
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text: string) => text || '-',
    },
  ];

  return (
    <div className={styles.tableWrapper}>
      <Table
        columns={columns}
        dataSource={pagedData}
        rowKey="id"
        loading={loading}
        pagination={{
          pageSize,
          current: pageNumber,
          total,
          onChange: (page, size) => {
            setPageNumber(page);
            setPageSize(size);
          },
        }}
        bordered
      />
    </div>
  );
};

export default ExtractTaskList;
