
import type { KnowledgeGraphWorkspace } from '../model/kg_workspace';
import request from '../utils/req';

export interface ListKnowledgeGraphWorkspaceParams {
  page_size: number;
  page_number: number;
}

export async function ListKnowledgeGraphWorkspace(params: ListKnowledgeGraphWorkspaceParams): Promise<any> {
    return request.post(`/knowledge_graph_workspace/list`, {
        ...params
    } as any);
}


export interface DeleteKnowledgeGraphWorkspaceParams {
    ids: number[];
}

export async function DeleteKnowledgeGraphWorkspaces(params: DeleteKnowledgeGraphWorkspaceParams): Promise<any> {
    return request.post(`/knowledge_graph_workspace/delete`, params);
}

export interface GetKnowledgeGraphWorkspaceParams {
    id: number;
}
export async function GetKnowledgeGraphWorkspace(params: GetKnowledgeGraphWorkspaceParams): Promise<any> {
    return request.post(`/knowledge_graph_workspace/get`, params);
}

export interface CreateKnowledgeGraphWorkspaceParams {
    knowledge_graph_workspace: KnowledgeGraphWorkspace;
}

export async function CreateKnowledgeGraphWorkspace(params: CreateKnowledgeGraphWorkspaceParams): Promise<any> {
    return request.post(`/knowledge_graph_workspace/create`, params);
}


export interface UpdateKnowledgeGraphWorkspaceParams {
    knowledge_graph_workspace: KnowledgeGraphWorkspace;
}

export async function UpdateKnowledgeGraphWorkspace(params: UpdateKnowledgeGraphWorkspaceParams): Promise<any> {
    return request.post(`/knowledge_graph_workspace/update`, params);
}
