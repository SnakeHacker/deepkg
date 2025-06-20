export interface ExtractTaskResult {
    task_id: number;
    nodes?: Entity[];
    edges?: Relationship[];
}

export interface Entity {
    id?: number;
    task_id?: number;
    entity_name: string;
    props?: EntityProp[];
}

export interface EntityProp {
    id?: number;
    task_id?: number;
    entity_id?: number;
    prop_name: string;
    prop_value: string;
}


export interface Relationship {
    id?: number;
    task_id?: number;
    source_entity_id: number;
    source_entity_name: string;
    target_entity_id: number;
    target_entity_name: string;
    relationship_name: string;
}