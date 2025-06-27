import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Button } from "antd";
import { PlusOutlined, RollbackOutlined } from "@ant-design/icons";
import type { Entity, EntityProp, Relationship } from "../../model/extract_task_result";
import { Graph } from '@antv/g6';
import { GetExtractTaskResult } from "../../service/extract_task_result";
const ExtractTaskResultPage: React.FC = () => {

    const [taskID, setTaskID] = useState(0);
    const [graphData, setGraphData]= useState<any>({});

    let graph: Graph | null = null;

    // const data = {
    //     nodes: [
    //     { id: 'node0', size: 50 },
    //     { id: 'node1', size: 30 },
    //     { id: 'node2', size: 30 },
    //     { id: 'node3', size: 30 },
    //     { id: 'node4', size: 30, isLeaf: true },
    //     { id: 'node5', size: 30, isLeaf: true },
    //     { id: 'node6', size: 15, isLeaf: true },
    //     { id: 'node7', size: 15, isLeaf: true },
    //     { id: 'node8', size: 15, isLeaf: true },
    //     { id: 'node9', size: 15, isLeaf: true },
    //     { id: 'node10', size: 15, isLeaf: true },
    //     { id: 'node11', size: 15, isLeaf: true },
    //     { id: 'node12', size: 15, isLeaf: true },
    //     { id: 'node13', size: 15, isLeaf: true },
    //     { id: 'node14', size: 15, isLeaf: true },
    //     { id: 'node15', size: 15, isLeaf: true },
    //     { id: 'node16', size: 15, isLeaf: true },
    //     ],
    //     edges: [
    //     { source: 'node0', target: 'node1' },
    //     { source: 'node0', target: 'node2' },
    //     { source: 'node0', target: 'node3' },
    //     { source: 'node0', target: 'node4' },
    //     { source: 'node0', target: 'node5' },
    //     { source: 'node1', target: 'node6' },
    //     { source: 'node1', target: 'node7' },
    //     { source: 'node2', target: 'node8' },
    //     { source: 'node2', target: 'node9' },
    //     { source: 'node2', target: 'node10' },
    //     { source: 'node2', target: 'node11' },
    //     { source: 'node2', target: 'node12' },
    //     { source: 'node2', target: 'node13' },
    //     { source: 'node3', target: 'node14' },
    //     { source: 'node3', target: 'node15' },
    //     { source: 'node3', target: 'node16' },
    //     ],
    // };


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
                },
            },
            layout: {
                type: 'd3-force',
                link: {
                    distance: (d: any) => {
                        return 300;
                    },
                    strength: (d: any) => {
                        return 0.1;
                    },
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
    }, [graphData])

    useEffect(() => {
        taskID > 0 && getExtractTaskResult()
    }, [taskID]);


    const getExtractTaskResult = async () => {
        const res = await GetExtractTaskResult({
            task_id: taskID,
        });


        if (res){

            const {entities, relationships }=res.extract_task_result

            const data = {
                nodes: entities.flatMap((entity: Entity) => {
                    const entityNode = {
                        id: `entityNode${entity.id}`,
                        size: 30,
                        labelText: entity.entity_name
                        // ...entity
                    };
                    const propNodes = (entity.props || []).map((prop: EntityProp) => ({
                        id: `propNode${prop.id}`,
                        size: 15,
                        isLeaf: true,
                        labelText: prop.prop_value
                        // ...prop
                    }));
                    return [entityNode, ...propNodes];
                }),
                edges: relationships.map((relationship: Relationship) => ({
                    source: `entityNode${relationship.source_entity_id}`,
                    target: `entityNode${relationship.target_entity_id}`,
                    labelText: relationship.relationship_name
                })),
            };

            setGraphData(data)
        }
    };


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
                    onClick={() => {
                    }}
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
