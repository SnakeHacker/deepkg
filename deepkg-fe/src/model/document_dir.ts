
export interface DocumentDir {
    id?: number;
    dir_name: string;
    parent_id?: number;
    children?: DocumentDir[];
    sort_index?: number;
    remark?: string;
    created_at?: string;
    updated_at?: string;
}
