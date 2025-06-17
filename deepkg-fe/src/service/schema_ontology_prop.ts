
import type { SchemaOntologyProp } from '../model/schema_ontology_prop';
import request from '../utils/req';

export interface ListSchemaOntologyPropParams {
  ontology_id: number;
  page_size: number;
  page_number: number;
}

export async function ListSchemaOntologyProp(params: ListSchemaOntologyPropParams): Promise<any> {
    return request.post(`/schema_ontology_prop/list`, {
        ...params
    } as any);
}


export interface DeleteSchemaOntologyPropParams {
    ids: number[];
}

export async function DeleteSchemaOntologyProps(params: DeleteSchemaOntologyPropParams): Promise<any> {
    return request.post(`/schema_ontology_prop/delete`, params);
}

export interface GetSchemaOntologyPropParams {
    id: number;
}
export async function GetSchemaOntologyProp(params: GetSchemaOntologyPropParams): Promise<any> {
    return request.post(`/schema_ontology_prop/get`, params);
}

export interface CreateSchemaOntologyPropParams {
    schema_ontology_prop: SchemaOntologyProp;
}


export async function CreateSchemaOntologyProp(params: CreateSchemaOntologyPropParams): Promise<any> {
    return request.post(`/schema_ontology_prop/create`, params);
}


export interface UpdateSchemaOntologyPropParams {
    schema_ontology_prop: SchemaOntologyProp;
}

export async function UpdateSchemaOntologyProp(params: UpdateSchemaOntologyPropParams): Promise<any> {
    return request.post(`/schema_ontology_prop/update`, params);
}
