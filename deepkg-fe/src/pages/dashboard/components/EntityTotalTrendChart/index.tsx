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
import { GetSchemaOntologyDailyCount } from '../../../../service/schema_ontology';
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
      const response = await GetSchemaOntologyDailyCount();
      const items = response.items || [];
      // 计算累计总数
      let sum = 0;
      const counts = items.map(item => {
        sum += item.count;
        return sum;
      });
      const dates = items.map(item => {
        const date = new Date(item.date);
        return `${date.getMonth() + 1}月${date.getDate()}日`;
      });
      setChartData({ dates, counts });
    } catch (error) {
      console.error('获取实体累计数量数据失败:', error);
      setChartData({
        dates: ['6月17日', '6月18日', '6月19日', '6月20日', '6月21日', '6月22日', '6月23日'],
        counts: [0, 0, 0, 0, 0, 0, 0]
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDailyData();
    const interval = setInterval(fetchDailyData, 30000);
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
      bottom: "10%",
      right: "5%",
      left: "5%"
    },
    xAxis: {
      type: "category",
      data: chartData.dates,
    },
    yAxis: {
      type: "value",
      max: Math.max(...chartData.counts, 10),
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
