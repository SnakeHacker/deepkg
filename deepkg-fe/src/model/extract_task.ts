import type { Document } from "./document";
import type { SchemaTriple } from "./schema_triple";

export interface ExtractTask {
    id?: number;
    task_name: string;
    work_space_id: number;
    docs? : Document[];
    triples?: SchemaTriple[];
    remark?: string;

    creator_id?: number;
    creator_name?: string;

    created_at?: string;
    updated_at?: string;
}
