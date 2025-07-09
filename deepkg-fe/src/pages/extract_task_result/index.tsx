import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button } from "antd";
import { PlusOutlined, RollbackOutlined } from "@ant-design/icons";
import { Graph } from '@antv/g6';
import { GetExtractTaskResult } from "../../service/extract_task_result";
import {PublishExtractTask} from "../../service/extract_task.ts";
const ExtractTaskResultPage: React.FC = () => {

    const [taskID, setTaskID] = useState(0);
    const [graphData, setGraphData]= useState<any>({});

    let graph: Graph | null = null;

    useEffect(() => {
        const hash = window.location.hash;
        const hashParts = hash.split('?');
        if (hashParts.length > 1) {
            const searchParams = new URLSearchParams(hashParts[1]);
            const taskID = searchParams.get('task_id');
            setTaskID(Number(taskID))
        } else {
            console.warn('URL hash 中不包含查询参数');
        }

    }, []);

    useEffect(()=>{

        console.log(graphData)

        if (!graphData.nodes ){
            return
        }
        graph = new Graph({
            container: 'container',
            data: graphData,
            node: {
                style: {
                    size: (d: any) => d.size,
                    labelText: (d: any) => d.labelText,
                },
            },
            edge: {
                style: {
                    labelText: (d: any) => d.labelText,
                    endArrow: true,
                },
            },
            layout: {
                type: 'd3-force',
                link: {
                    distance: 300,
                    strength: 0.1,
                },
                manyBody: {
                    strength: (d: any) => {
                    if (d.isLeaf) {
                        return -50;
                    }
                    return -10;
                    },
                },
            },
            behaviors: [
                {
                  type: 'drag-element-force',
                  key: 'drag-element-force-1',
                  fixed: true, // 拖拽后固定节点位置
                },
                'zoom-canvas'
            ],
            // behaviors: ['drag-node'],
        });

        graph.render();
    }, [graphData]);

    useEffect(() => {
        taskID > 0 && getExtractTaskResult()
    }, [taskID]);


    const getExtractTaskResult = async () => {
        const res = await GetExtractTaskResult({
            task_id: taskID,
        });


        if (res){

            const {entities, relationships} = res.extract_task_result

            const nodes = [];
            const edges = [];

            for (const entity of entities) {
                const entityNode = {
                    id: `entityNode${entity.id}`,
                    size: 30,
                    labelText: entity.entity_name
                };
                nodes.push(entityNode);

                for (const prop of entity.props || []) {
                    const propNode = {
                        id: `propNode${prop.id}`,
                        size: 15,
                        labelText: prop.prop_value
                    };
                    nodes.push(propNode);

                    const propEdge = {
                        source: `entityNode${entity.id}`,
                        target: `propNode${prop.id}`,
                        labelText: prop.prop_name
                    }
                    edges.push(propEdge);
                }
            }

            for (const relationship of relationships) {
                const sourceId = `entityNode${relationship.source_entity_id}`;
                const targetId = `entityNode${relationship.target_entity_id}`;

                // 查找是否已有相同起始节点和目标节点的边
                const existingEdge = edges.find(edge =>
                    edge.source === sourceId && edge.target === targetId
                );

                if (existingEdge) {
                    // 如果找到已有的边，在标签文本中添加新的关系名
                    existingEdge.labelText = existingEdge.labelText + ', ' + relationship.relationship_name;
                } else {
                    // 如果没有找到，创建新的边
                    const edge = {
                        source: sourceId,
                        target: targetId,
                        labelText: relationship.relationship_name
                    };
                    edges.push(edge);
                }
            }

            const data = {
                nodes: nodes,
                edges: edges,
            };

            setGraphData(data)
        }
    };

    const publishExtractTask = async () => {
        const res = await PublishExtractTask({id: taskID})
        if (res) {
            console.log('发布成功');
        } else {
            console.error('发布失败');
        }
    }


    return (
        <div className={styles.container}
            style={{
                backgroundImage: `url(${BgSVG})`,
            }}
        >

            <div className={styles.header}>

                <Button
                    type="default"
                    icon={<RollbackOutlined />}
                    onClick={() => {
                        window.history.back();
                    }}
                    style={{ marginRight: '10px'}}
                >
                    返回
                </Button>
                <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={publishExtractTask}
                >
                    发布
                </Button>
            </div>
            <div className={styles.body}>
                <div id="container"className={styles.graphContainer}/>
            </div>

        </div>
    )
};

export default ExtractTaskResultPage;
