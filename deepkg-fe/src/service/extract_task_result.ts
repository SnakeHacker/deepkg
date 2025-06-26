import request from '../utils/req';

export interface GetExtractTaskResultParams {
    task_id: number;
}

export async function GetExtractTaskResult(params: GetExtractTaskResultParams): Promise<any> {
    return request.post(`/extract_task_result/get`, params);
}