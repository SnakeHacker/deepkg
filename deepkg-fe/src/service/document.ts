
import type { Document } from '../model/document';
import request from '../utils/req';

export interface ListDocumentParams {
  dir_id?: number;
  page_size: number;
  page_number: number;
}

export async function ListDocument(params: ListDocumentParams): Promise<any> {
    return request.post(`/document/list`, {
        ...params
    } as any);
}


export interface DeleteDocumentParams {
    ids: number[];
}

export async function DeleteDocuments(params: DeleteDocumentParams): Promise<any> {
    return request.post(`/document/delete`, params);
}

export interface GetDocumentParams {
    id: number;
}
export async function GetDocument(params: GetDocumentParams): Promise<any> {
    return request.post(`/document/get`, params);
}

export interface CreateDocumentParams {
    document: Document;
}

export async function CreateDocument(params: CreateDocumentParams): Promise<any> {
    return request.post(`/document/create`, params);
}


export interface UpdateDocumentParams {
    document: Document;
}

export async function UpdateDocument(params: UpdateDocumentParams): Promise<any> {
    return request.post(`/document/update`, params);
}

