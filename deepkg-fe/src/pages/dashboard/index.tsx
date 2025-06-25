import React from 'react';
import { Card, Col, Row } from 'antd';
import styles from './index.module.less';
import WelcomeCard from './components/WelcomeCard';
import DashboardCards from './components/DashboardCards';
import WorkspaceEntityPieChart from './components/WorkspaceEntityPieChart';
import EntityTotalTrendChart from './components/EntityTotalTrendChart';
import EntityAdditonTrendChart from './components/EntityAdditonTrendChart';
import ExtractTaskList from './components/ExtractTaskList';

const Dashboard: React.FC = () => {
    return (
        <div className={styles.dashboardWrapper}>
            <WelcomeCard />
            <DashboardCards />

            {/* 第一行：实体分布图 + 抽取任务列表 */}
            <Row gutter={[16, 16]} style={{ marginTop: 16 }}>
                <Col xs={24} md={10}>
                    <Card title="各工作空间实体数量分布" bordered={false} style={{ height: '420px' }}>
                        <div style={{ height: '360px' }}>
                            <WorkspaceEntityPieChart />
                        </div>
                    </Card>
                </Col>
                <Col xs={24} md={14}>
                    <Card title="抽取任务列表" bordered={false} style={{ height: '420px' }}>
                        <div style={{ maxHeight: '360px', overflowY: 'auto' }}>
                            <ExtractTaskList />
                        </div>
                    </Card>
                </Col>
            </Row>

            {/* 第二行：实体新增趋势图 + 实体数量趋势图 */}
            <Row gutter={[16, 16]} style={{ marginTop: 16 }}>
                <Col xs={24} md={12}>
                    <Card title="实体新增趋势图" bordered={false} style={{ height: '420px' }}>
                        <div style={{ height: '360px' }}>
                            <EntityAdditonTrendChart />
                        </div>
                    </Card>
                </Col>
                <Col xs={24} md={12}>
                    <Card title="实体数量趋势图" bordered={false} style={{ height: '420px' }}>
                        <div style={{ height: '360px' }}>
                            <EntityTotalTrendChart />
                        </div>
                    </Card>
                </Col>
            </Row>
        </div>
    );
};

export default Dashboard;
