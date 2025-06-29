import React, { useEffect, useRef, useState } from 'react';
import * as echarts from 'echarts/core';
import { BarChart } from 'echarts/charts';
import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent
} from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import ReactECharts from 'echarts-for-react';
import { GetEntityTotalDailyCount } from '../../../../service/dashboard';
import { Spin } from 'antd';
import styles from './index.module.less';

// 注册 ECharts 组件
echarts.use([
  BarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent,
  CanvasRenderer
]);

const EntityTotalTrendChart: React.FC = () => {
  const chartRef = useRef(null);
  const [loading, setLoading] = useState(true);
  const [chartData, setChartData] = useState<{ dates: string[], counts: number[] }>({
    dates: [],
    counts: []
  });

  // 获取近七天实体累计总数数据
  const fetchDailyData = async () => {
    try {
      setLoading(true);
      const response = await GetEntityTotalDailyCount();

      const dates = response.map(item => {
        const date = new Date(item.date);
        return `${date.getMonth() + 1}月${date.getDate()}日`;
      });

      const counts = response.map(item => item.count);

      setChartData({ dates, counts });
    } catch (error) {
      console.error('获取实体累计数量数据失败:', error);
      // 如果获取失败，使用默认数据
      const today = new Date();
      const defaultDates: string[] = [];
      const defaultCounts: number[] = [];

      for (let i = 6; i >= 0; i--) {
        const date = new Date(today);
        date.setDate(date.getDate() - i);
        defaultDates.push(`${date.getMonth() + 1}月${date.getDate()}日`);
        defaultCounts.push(0);
      }

      setChartData({
        dates: defaultDates,
        counts: defaultCounts
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDailyData();
    const interval = setInterval(fetchDailyData, 60000);
    return () => clearInterval(interval);
  }, []);

  const option = {
    legend: {
      show: true,
      data: ["实体累计总数"],
      icon: "rect",
      itemWidth: 10,
      itemHeight: 10,
      right: "5%",
      top: "1%",
      textStyle: {
        color: "rgba(0,0,0,0.65)",
        fontSize: 15,
        lineHeight: 30,
      },
    },
    grid: {
      top: "18%",
      bottom: "20%",
      right: "5%",
      left: "5%"
    },
    xAxis: {
      type: "category",
      data: chartData.dates,
    },
    yAxis: {
      type: "value",
      max: (value: any) => Math.ceil(value.max / 10) * 10,
      minInterval: 1,
    },
    series: [
      {
        name: "实体累计总数",
        type: "bar",
        data: chartData.counts,
        itemStyle: {
          color: "#5285fd",
        },
        barWidth: 30,
      },
    ],
  };

  if (loading) {
    return (
      <div className={styles.loadingWrapper}>
        <Spin size="large" tip="加载累计数据中..." />
      </div>
    );
  }

  return (
    <div className={styles.chartWrapper}>
      <ReactECharts
        ref={chartRef}
        option={option}
        style={{ height: '100%', width: '100%' }}
      />
    </div>
  );
};

export default EntityTotalTrendChart;
