
import type { DocumentDir } from '../model/document_dir';
import request from '../utils/req';



export async function ListDocumentDir(): Promise<any> {
    return request.post(`/document_dir/list`, {
    } as any);
}


export interface DeleteDocumentDirParams {
    ids: number[];
}

export async function DeleteDocumentDirs(params: DeleteDocumentDirParams): Promise<any> {
    return request.post(`/document_dir/delete`, params);
}

export interface CreateDocumentDirParams {
    document_dir: DocumentDir;
}

export async function CreateDocumentDir(params: CreateDocumentDirParams): Promise<any> {
    return request.post(`/document_dir/create`, params);
}


export interface UpdateDocumentDirParams {
    document_dir: DocumentDir;
}

export async function UpdateDocumentDir(params: UpdateDocumentDirParams): Promise<any> {
    return request.post(`/document_dir/update`, params);
}
