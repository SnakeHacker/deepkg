import request from '../utils/req';
import { ListKnowledgeGraphWorkspace } from './workspace';
import type { KnowledgeGraphWorkspace } from '../model/kg_workspace';


interface BasicTotalResp {
  total: number;
}

export interface GetDocumentTotalCountResp extends BasicTotalResp {}
export interface GetOrganizationTotalCountResp extends BasicTotalResp {}
export interface GetKnowledgeGraphWorkspaceTotalCountResp extends BasicTotalResp {}
export interface ListExtractTaskResp extends BasicTotalResp {}
export interface ListSchemaOntologyResp extends BasicTotalResp {}

/** 文档总数 */
export async function GetDocumentTotalCount(dir_id?: number): Promise<number> {
  try {
    const resp: GetDocumentTotalCountResp = await request.post('/document/list', {
      page_size: 1,
      page_number: 1,
      dir_id: dir_id ?? undefined,
    });
    return resp.total || 0;
  } catch (e) {
    console.error('获取文档总数失败:', e);
    return 0;
  }
}

/** 组织机构总数 */
export async function GetOrganizationTotalCount(): Promise<number> {
  try {
    const resp: GetOrganizationTotalCountResp = await request.post('/org/list', {
      page_size: 1,
      page_number: 1,
    });
    return resp.total || 0;
  } catch (e) {
    console.error('获取组织机构总数失败:', e);
    return 0;
  }
}

/** 工作空间总数 */
export async function GetKnowledgeGraphWorkspaceTotalCount(): Promise<number> {
  try {
    const resp: GetKnowledgeGraphWorkspaceTotalCountResp = await request.post('/knowledge_graph_workspace/list', {
      page_size: 1,
      page_number: 1,
    });
    return resp.total || 0;
  } catch (e) {
    console.error('获取工作空间总数失败:', e);
    return 0;
  }
}

/** 所有工作空间下抽取任务总数 */
export async function GetTotalExtractTaskCountAllWorkspaces(): Promise<number> {
  try {
    const workspaceResp = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
    const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];

    let total = 0;
    for (const ws of workspaces) {
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        const res: ListExtractTaskResp = await request.post('/extract_task/list', {
          work_space_id: ws.id,
          page_size: 1,
          page_number: 1,
        });
        total += res.total || 0;
      }
    }
    return total;
  } catch (error) {
    console.error('获取抽取任务总数失败:', error);
    return 0;
  }
}

/** 所有工作空间下实体总数 */
export async function GetTotalSchemaOntologyCountAllWorkspaces(): Promise<number> {
  try {
    const workspaceResp = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
    const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];

    let total = 0;
    for (const ws of workspaces) {
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        const res: ListSchemaOntologyResp = await request.post('/schema_ontology/list', {
          work_space_id: ws.id,
          page_size: 1,
          page_number: 1,
        });
        total += res.total || 0;
      }
    }
    return total;
  } catch (error) {
    console.error('获取实体总数失败:', error);
    return 0;
  }
}

/** 实体趋势（近七日每日总数） */
export interface DailyCountItem {
  date: string;
  count: number;
}

export interface GetSchemaOntologyDailyCountResp {
  items: DailyCountItem[];
}

export async function GetSchemaOntologyDailyCount(): Promise<GetSchemaOntologyDailyCountResp> {
  return request.post('/schema_ontology/daily_count', {
    work_space_id: 0,
  });
}
