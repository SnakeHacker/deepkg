
export interface Document {
    id?: number;
    doc_name: string;
    doc_desc?: string;
    doc_path: string;
    dir_id: number;

    creator_id?: number;
    creator_name?: string;

    created_at?: string;
    updated_at?: string;
}
