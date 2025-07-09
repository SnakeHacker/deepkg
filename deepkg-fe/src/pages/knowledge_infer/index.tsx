import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';
import { Select } from "antd";
import type { KnowledgeGraphWorkspace } from "../../model/kg_workspace";
import { ListKnowledgeGraphWorkspace } from "../../service/workspace";
import ChatContainer from "./component/chat";

const KnowledgeInferPage: React.FC = () => {

    const [workspaces, setWorkspaces] = useState<KnowledgeGraphWorkspace[]>([]);
    const [workspaceID, setWorkspaceID] = useState(0);

    useEffect(() => {
        listWorkspaces()
    }, []);

    const listWorkspaces = async () => {
        const res = await ListKnowledgeGraphWorkspace({
            page_size: 10,
            page_number: -1
        });
        setWorkspaces(res.knowledge_graph_workspaces);
        if (res.knowledge_graph_workspaces.length > 0){
            setWorkspaceID(res.knowledge_graph_workspaces[res.knowledge_graph_workspaces.length-1].id)
        }
    };



    return (
        <div className={styles.container}
            style={{
                backgroundImage: `url(${BgSVG})`,
            }}
        >

            <div className={styles.header}>
                <Select
                    style={{'width': '200px', 'marginRight': '10px'}}
                    placeholder="工作空间"
                    disabled={workspaces.length === 0}
                    onSelect={(value) => setWorkspaceID(value)}
                    value={workspaceID}
                    options={[
                        ...workspaces.map((workspaces) => (
                            {key: workspaces.id, label: workspaces.knowledge_graph_workspace_name, value: workspaces.id}
                        )),
                    ]}
                />


            </div>
            <div className={styles.body}>
                <ChatContainer
                  key={`chat-${workspaceID}`} // 添加 key 属性
                  workspaceID={workspaceID}
                />

            </div>

            <div className={styles.footer}>

            </div>
        </div>
    )
};

export default KnowledgeInferPage;
