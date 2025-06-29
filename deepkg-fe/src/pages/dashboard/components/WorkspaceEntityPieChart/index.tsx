import React, { useEffect, useState, useRef } from 'react';
import ReactECharts from 'echarts-for-react';
import * as echarts from 'echarts/core';
import { PieChart } from 'echarts/charts';
import {
  TooltipComponent,
  LegendComponent,
  TitleComponent,
} from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import { ListKnowledgeGraphWorkspace } from '../../../../service/workspace';
import type { ListSchemaOntologyResp } from '../../../../service/dashboard';
import request from '../../../../utils/req';
import styles from './index.module.less';

echarts.use([
  PieChart,
  TooltipComponent,
  LegendComponent,
  TitleComponent,
  CanvasRenderer,
]);

const WorkspaceEntityPieChart: React.FC = () => {
  const chartRef = useRef(null);
  const [data, setData] = useState<{ name: string; value: number }[]>([]);
  const [total, setTotal] = useState<number>(0);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await ListKnowledgeGraphWorkspace({
          page_size: 1000,
          page_number: 1,
        });

        const workspaces = res.knowledge_graph_workspaces || [];
        const chartData: { name: string; value: number }[] = [];

        for (const ws of workspaces) {
          const countRes = await request.post('/schema_ontology/list', {
            work_space_id: ws.id,
            page_number: 1,
            page_size: 1,
          }) as ListSchemaOntologyResp;

          const displayName = ws.knowledge_graph_workspace_name?.trim() || `ID-${ws.id}`;

          chartData.push({
            name: displayName,
            value: countRes.total || 0,
          });
        }

        const totalCount = chartData.reduce((sum, item) => sum + item.value, 0);
        setData(chartData);
        setTotal(totalCount);
      } catch (error) {
        console.error('获取实体数量失败:', error);
      }
    };

    fetchData();
  }, []);

  const colorList = [
    { c1: '#1E60FB', c2: '#5D8EFE' },
    { c1: '#6CD8D0', c2: '#1DC7B5' },
    { c1: '#F9D370', c2: '#F7BD26' },
    { c1: '#B28AE9', c2: '#9358E3' },
    { c1: '#EA7283', c2: '#F53D57' },
    { c1: '#73C0DE', c2: '#5470C6' },
    { c1: '#91CC75', c2: '#3BA272' },
  ];

  const option = {
    title: {
      text: '100%',
      left: '35%',
      top: '40%',  // 从 '50%' 调整为 '45%'
      textAlign: 'center',
      textVerticalAlign: 'middle',
      textStyle: {
        fontSize: 30,
        color: 'rgba(0,0,0,0.65)',
      },
    },
    tooltip: {
      trigger: 'item',
      formatter: (params: any) => {
        const { marker, name, value, percent } = params;
        return `${marker} ${name}<br/>实体数量：${value}<br/>占比：${percent}%`;
      }
    },
    legend: {
      type: 'scroll',
      orient: 'vertical',
      left: '75%',
      top: '0',
      icon: 'circle',
      itemWidth: 8,
      itemHeight: 8,
      formatter: (name: string) => `{a|${name}}`,
      textStyle: {
        lineHeight: 25,
        rich: {
          a: {
            fontSize: 14,
            width: 100,
            color: '#000',
          },
        },
      },
    },
    series: [
      {
        name: '实体数量',
        type: 'pie',
        center: ['35%', '40%'],
        radius: ['35%', '55%'],
        clockwise: false,
        label: {
          formatter: '{d}%',
          position: 'outside',
        },
        itemStyle: {
          color: (params: any) => {
            const color = colorList[params.dataIndex % colorList.length];
            return new echarts.graphic.LinearGradient(1, 0, 0, 0, [
              { offset: 0, color: color.c1 },
              { offset: 1, color: color.c2 },
            ]);
          },
        },
        data: data,
      },
    ],
  };


  return (
    <div className={styles.chartWrapper}>
      <ReactECharts
        ref={chartRef}
        option={option}
        style={{ width: '100%', height: '100%' }}
      />
    </div>
  );
};

export default WorkspaceEntityPieChart;
