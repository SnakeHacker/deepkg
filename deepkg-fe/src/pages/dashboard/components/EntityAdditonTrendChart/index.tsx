import React, { useEffect, useRef, useState } from 'react';
import * as echarts from 'echarts/core';
import { LineChart } from 'echarts/charts';
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

echarts.use([
  LineChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  TitleComponent,
  CanvasRenderer
]);

const EntityAdditonTrendChart: React.FC = () => {
  const chartRef = useRef(null);
  const [loading, setLoading] = useState(true);
  const [chartData, setChartData] = useState<{ dates: string[], counts: number[] }>({
    dates: [],
    counts: []
  });

  // 获取近七天实体数量数据
  const fetchDailyData = async () => {
    try {
      setLoading(true);
      const response = await GetSchemaOntologyDailyCount();
      const items = response.items || [];
      
      const dates = items.map(item => {
        // 将日期格式从 '2025-06-23' 转换为 '6月23日'
        const date = new Date(item.date);
        return `${date.getMonth() + 1}月${date.getDate()}日`;
      });
      
      const counts = items.map(item => item.count);
      
      setChartData({ dates, counts });
    } catch (error) {
      console.error('获取实体数量趋势数据失败:', error);
      // 如果获取失败，使用默认数据
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
    // 每30秒刷新一次数据
    const interval = setInterval(fetchDailyData, 30000);
    return () => clearInterval(interval);
  }, []);

  const option = {
    legend: {
      show: true,
      itemGap: 50,
      data: ["实体新增趋势"],
      icon: "circle",
      itemWidth: 6,
      itemHeight: 6,
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
      max: Math.max(...chartData.counts, 10), // 动态设置最大值
    },
    series: [
      {
        name: "实体新增趋势",
        type: "line",
        data: chartData.counts,
        lineStyle: {
          color: "#5285fd",
        },
        symbolSize: 8,
        itemStyle: {
          color: "#427afd",
          borderColor: "#427afd",
          borderWidth: 1,
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: "rgba(204, 220, 254,1)" },
            { offset: 1, color: "rgba(204, 220, 254,0)" },
          ]),
        },
      },
    ],
  };

  if (loading) {
    return (
      <div className={styles.loadingWrapper}>
        <Spin size="large" tip="加载趋势数据中..." />
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

export default EntityAdditonTrendChart;
