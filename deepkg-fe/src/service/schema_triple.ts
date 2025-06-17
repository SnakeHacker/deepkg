
import type { SchemaTriple } from '../model/schema_triple';
import request from '../utils/req';

export interface ListSchemaTripleParams {
  work_space_id?: number;
  page_size: number;
  page_number: number;
}

export async function ListSchemaTriple(params: ListSchemaTripleParams): Promise<any> {
    return request.post(`/schema_triple/list`, {
        ...params
    } as any);
}


export interface DeleteSchemaTripleParams {
    ids: number[];
}

export async function DeleteSchemaTriples(params: DeleteSchemaTripleParams): Promise<any> {
    return request.post(`/schema_triple/delete`, params);
}

export interface GetSchemaTripleParams {
    id: number;
}
export async function GetSchemaTriple(params: GetSchemaTripleParams): Promise<any> {
    return request.post(`/schema_triple/get`, params);
}

export interface CreateSchemaTripleParams {
    schema_triple: SchemaTriple;
}


export async function CreateSchemaTriple(params: CreateSchemaTripleParams): Promise<any> {
    return request.post(`/schema_triple/create`, params);
}


export interface UpdateSchemaTripleParams {
    schema_triple: SchemaTriple;
}

export async function UpdateSchemaTriple(params: UpdateSchemaTripleParams): Promise<any> {
    return request.post(`/schema_triple/update`, params);
}
