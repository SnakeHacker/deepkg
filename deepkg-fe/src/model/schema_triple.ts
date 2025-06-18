
export interface SchemaTriple {
    id?: number;
    source_ontology_id?: number;
    source_ontology_name?: string;
    target_ontology_id?: number;
    target_ontology_name?: string;
    relationship?: string;
    work_space_id?: number;

    creator_id?: number;
    creator_name?: string;

    created_at?: string;
    updated_at?: string;
}
