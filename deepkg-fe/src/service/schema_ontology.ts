
import type { SchemaOntology } from '../model/schema_ontology';
import request from '../utils/req';

export interface ListSchemaOntologyParams {
  work_space_id: number;
  page_size: number;
  page_number: number;
}

export async function ListSchemaOntology(params: ListSchemaOntologyParams): Promise<any> {
    return request.post(`/schema_ontology/list`, {
        ...params
    } as any);
}


export interface DeleteSchemaOntologyParams {
    ids: number[];
}

export async function DeleteSchemaOntologys(params: DeleteSchemaOntologyParams): Promise<any> {
    return request.post(`/schema_ontology/delete`, params);
}

export interface GetSchemaOntologyParams {
    id: number;
}
export async function GetSchemaOntology(params: GetSchemaOntologyParams): Promise<any> {
    return request.post(`/schema_ontology/get`, params);
}

export interface CreateSchemaOntologyParams {
    schema_ontology: SchemaOntology;
}


export async function CreateSchemaOntology(params: CreateSchemaOntologyParams): Promise<any> {
    return request.post(`/schema_ontology/create`, params);
}


export interface UpdateSchemaOntologyParams {
    schema_ontology: SchemaOntology;
}

export async function UpdateSchemaOntology(params: UpdateSchemaOntologyParams): Promise<any> {
    return request.post(`/schema_ontology/update`, params);
}
