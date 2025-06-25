import type { ExtractTask } from '../model/extract_task';
import request from '../utils/req';

//分页查询参数类型
export interface ListExtractTaskParams {
  work_space_id?: number;
  page_size: number;
  page_number: number;
}

//分页查询方法
export async function ListExtractTask(params: ListExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/list`, {
        ...params
    } as any);
}


export interface DeleteExtractTaskParams {
    ids: number[];
}

export async function DeleteExtractTasks(params: DeleteExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/delete`, params);
}

export interface GetExtractTaskParams {
    id: number;
}
export async function GetExtractTask(params: GetExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/get`, params);
}

export interface CreateExtractTaskParams {
    extract_task: ExtractTask;
}


export async function CreateExtractTask(params: CreateExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/create`, params);
}


export interface UpdateExtractTaskParams {
    extract_task: ExtractTask;
}

export async function UpdateExtractTask(params: UpdateExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/update`, params);
}

export interface PublishExtractTaskParams {
    id: number;
}

export async function PublishExtractTask(params: PublishExtractTaskParams): Promise<any> {
    return request.post(`/extract_task/publish`, params);
}
