import type { ExtractTask } from '../model/extract_task';
import request from '../utils/req';
import { ListKnowledgeGraphWorkspace } from './workspace'; // 引入工作空间查询方法
import type { KnowledgeGraphWorkspace } from '../model/kg_workspace';

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

// 获取所有工作空间下的抽取任务总数
export async function GetTotalExtractTaskCountAllWorkspaces(): Promise<number> {
  try {
    const workspaceResp = await ListKnowledgeGraphWorkspace({
      page_size: 1000,
      page_number: 1,
    }) as {
      total: number;
      knowledge_graph_workspaces: KnowledgeGraphWorkspace[];
      page_size: number;
      page_number: number;
    };

    const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];
    let total = 0;

    for (const ws of workspaces) {
      // 确保ws.id是有效数字再调用
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        const res = await ListExtractTask({
          work_space_id: ws.id,
          page_size: 1,
          page_number: 1,
        });
        total += res.total || 0;
      }
    }

    return total;
  } catch (error) {
    console.error('获取所有工作空间下抽取任务总数失败:', error);
    return 0;
  }
}

