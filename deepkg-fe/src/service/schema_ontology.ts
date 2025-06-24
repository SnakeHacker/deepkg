import type { SchemaOntology } from '../model/schema_ontology';
import request from '../utils/req';
import { ListKnowledgeGraphWorkspace } from './workspace';
import type { KnowledgeGraphWorkspace } from '../model/kg_workspace';


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
/**
 * 获取某工作空间下的本体类总数
 * @param workSpaceId 工作空间ID
 * @returns 本体类数量
 */
export interface SchemaOntologyListResponse {
  total: number;
  schema_ontologys: any[];
  page_size: number;
  page_number: number;
}

export async function GetSchemaOntologyTotalCount(work_space_id?: number): Promise<number> {
  try {
    const res: SchemaOntologyListResponse = await request.post('/schema_ontology/list', {
      work_space_id: 1,
      page_number: 1,
      page_size: 1,
    });

    return res.total || 0;
  } catch (error) {
    console.error('获取本体类总数失败:', error);
    return 0;
  }
}
// 获取所有工作空间下的实体总数
export async function GetTotalSchemaOntologyCountAllWorkspaces(): Promise<number> {
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
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        const res: SchemaOntologyListResponse = await request.post('/schema_ontology/list', {
          work_space_id: ws.id,
          page_number: 1,
          page_size: 1,
        });
        total += res.total || 0;
      }
    }

    return total;
  } catch (error) {
    console.error('获取所有工作空间下实体总数失败:', error);
    return 0;
  }
}

// ✅ 实体数量趋势图接口（近七日每日总数）
export interface DailyCountItem {
  date: string; // 例如 '2025-06-23'
  count: number;
}

export interface GetSchemaOntologyDailyCountResp {
  items: DailyCountItem[];
}

export async function GetSchemaOntologyDailyCount(): Promise<GetSchemaOntologyDailyCountResp> {
  return request.post('/schema_ontology/daily_count', {
    work_space_id: 0
  });
}

