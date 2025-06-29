import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Spin, message } from 'antd';
import styles from './index.module.less';

import entityIcon from '../../../../assets/dashboard_icon_1.png';
import relationIcon from '../../../../assets/dashboard_icon_2.png';
import fileIcon from '../../../../assets/dashboard_icon_3.png';
import taskIcon from '../../../../assets/dashboard_icon_4.png';
import workspaceIcon from '../../../../assets/dashboard_icon_5.png';

import {
  GetTotalEntityCountAllWorkspaces,
  GetDocumentTotalCount,
  GetTotalExtractTaskCountAllWorkspaces,
  GetOrganizationTotalCount,
  GetKnowledgeGraphWorkspaceTotalCount
} from '../../../../service/dashboard';


interface CardData {
  name: string;
  value: number;
  up: boolean;
  speed: number;
  img: string;
  route: string;
}

const DashboardCards: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [cardData, setCardData] = useState<CardData[]>([
    { name: '实体总数', value: 0, up: true, speed: 10, img: entityIcon, route: '/ontology' },
    { name: '组织总数', value: 0, up: true, speed: 5, img: relationIcon, route: '/org' },
    { name: '文件总数', value: 0, up: true, speed: 8, img: fileIcon, route: '/document' },
    { name: '抽取任务总数', value: 0, up: false, speed: 3, img: taskIcon, route: '/extract_task' },
    { name: '工作空间总数', value: 0, up: true, speed: 7, img: workspaceIcon, route: '/workspace' },
  ]);

  const fetchStats = async (isInitialLoad: boolean) => {
    if (isInitialLoad) setLoading(true);

    try {
      const [entityTotal, relationTotal, docTotal, taskTotal, workspaceTotal] = await Promise.all([
        GetTotalEntityCountAllWorkspaces(),
        GetOrganizationTotalCount(),
        GetDocumentTotalCount(),
        GetTotalExtractTaskCountAllWorkspaces(),
        GetKnowledgeGraphWorkspaceTotalCount(),
      ]);

      const updatedData: CardData[] = [
        { name: '实体总数', value: entityTotal, up: true, speed: 10, img: entityIcon, route: '/ontology' },
        { name: '组织总数', value: relationTotal, up: true, speed: 5, img: relationIcon, route: '/org' },
        { name: '文件总数', value: docTotal, up: true, speed: 8, img: fileIcon, route: '/document' },
        { name: '抽取任务总数', value: taskTotal, up: false, speed: 3, img: taskIcon, route: '/extract_task' },
        { name: '工作空间总数', value: workspaceTotal, up: true, speed: 7, img: workspaceIcon, route: '/workspace' },
      ];

      setCardData(updatedData);
    } catch (error) {
      console.error('获取统计数据失败:', error);
      if (isInitialLoad) {
        message.error('获取统计数据失败，请稍后重试');
      }
    } finally {
      if (isInitialLoad) setLoading(false);
    }
  };

  useEffect(() => {
    fetchStats(true); // 首次加载
    const intervalId = setInterval(() => fetchStats(false), 30000); // 每30秒轮询一次
    return () => clearInterval(intervalId); // 组件卸载时清除定时器
  }, []);

  if (loading) {
    return (
      <div className={styles.module3}>
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '200px' }}>
          <Spin size="large" tip="加载统计数据中..." />
        </div>
      </div>
    );
  }

  return (
    <div className={styles.module3}>
      {cardData.map((item, index) => (
        <div key={index} className={styles.card}>
          <div className={styles.cardTop}>
            <div>
              <div className={styles.name}>{item.name}</div>
              <div
                className={styles.value}
                style={{ cursor: 'pointer' }}
                onClick={() => navigate(item.route)}
              >
                {item.value.toLocaleString()}
              </div>
            </div>
            <img src={item.img} className={styles.icon} alt="" />
          </div>
        </div>
      ))}
    </div>
  );
};

export default DashboardCards;
