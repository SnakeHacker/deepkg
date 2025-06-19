
export interface SchemaOntologyProp {
    id?: number;
    prop_name: string;
    prop_desc?: string;
    work_space_id?: number;
    ontology_id: number;

    creator_id?: number;
    creator_name?: string;

    created_at?: string;
    updated_at?: string;
}
