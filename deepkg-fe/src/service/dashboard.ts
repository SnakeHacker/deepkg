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
export interface ListEntityResp extends BasicTotalResp {}

// 工作空间实体统计接口
export interface WorkspaceEntityCount {
  workspace_id: number;
  workspace_name: string;
  entity_count: number;
}

export interface GetWorkspaceEntityCountsResp {
  workspace_entity_counts: WorkspaceEntityCount[];
}

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
export async function GetTotalEntityCountAllWorkspaces(): Promise<number> {
  try {
    const workspaceResp = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
    const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];

    let total = 0;
    for (const ws of workspaces) {
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        try {
          // 先获取该工作空间下的所有任务
          const taskResp = await request.post('/extract_task/list', {
            work_space_id: ws.id,
            page_size: 1000, // 获取所有任务
            page_number: 1,
          });
          
          const tasks = (taskResp as any).extract_tasks || [];
          
          // 统计每个任务的实体数量
          for (const task of tasks) {
            if (task && typeof task.id === 'number' && task.id > 0) {
              try {
                const entityResp: ListEntityResp = await request.post('/entity/list', {
                  task_id: task.id,
                  page_size: 1,
                  page_number: 1,
                });
                total += entityResp.total || 0;
              } catch (taskError) {
                console.warn(`获取任务 ${task.id} 的实体数量失败:`, taskError);
                // 继续处理下一个任务
              }
            }
          }
        } catch (workspaceError) {
          console.warn(`获取工作空间 ${ws.id} 的任务列表失败:`, workspaceError);
          // 继续处理下一个工作空间
        }
      }
    }
    return total;
  } catch (error) {
    console.error('获取实体总数失败:', error);
    return 0;
  }
}

/** 获取每个工作空间的实体数量统计 */
export async function GetWorkspaceEntityCounts(): Promise<WorkspaceEntityCount[]> {
  try {
    const workspaceResp = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
    const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];

    const workspaceEntityCounts: WorkspaceEntityCount[] = [];

    for (const ws of workspaces) {
      if (ws && typeof ws.id === 'number' && ws.id > 0) {
        try {
          // 先获取该工作空间下的所有任务
          const taskResp = await request.post('/extract_task/list', {
            work_space_id: ws.id,
            page_size: 1000, // 获取所有任务
            page_number: 1,
          });
          
          const tasks = (taskResp as any).extract_tasks || [];
          let workspaceEntityCount = 0;
          
          // 统计每个任务的实体数量
          for (const task of tasks) {
            if (task && typeof task.id === 'number' && task.id > 0) {
              try {
                const entityResp: ListEntityResp = await request.post('/entity/list', {
                  task_id: task.id,
                  page_size: 1,
                  page_number: 1,
                });
                workspaceEntityCount += entityResp.total || 0;
              } catch (taskError) {
                console.warn(`获取任务 ${task.id} 的实体数量失败:`, taskError);
                // 继续处理下一个任务
              }
            }
          }

          workspaceEntityCounts.push({
            workspace_id: ws.id,
            workspace_name: ws.knowledge_graph_workspace_name || `工作空间-${ws.id}`,
            entity_count: workspaceEntityCount,
          });
        } catch (workspaceError) {
          console.warn(`获取工作空间 ${ws.id} 的任务列表失败:`, workspaceError);
          // 添加该工作空间，实体数量为0
          workspaceEntityCounts.push({
            workspace_id: ws.id,
            workspace_name: ws.knowledge_graph_workspace_name || `工作空间-${ws.id}`,
            entity_count: 0,
          });
        }
      }
    }

    return workspaceEntityCounts;
  } catch (error) {
    console.error('获取工作空间实体数量统计失败:', error);
    return [];
  }
}


/** 实体趋势（近七日每日总数） */
export interface DailyCountItem {
  date: string;
  count: number;
}

/** 获取最近7天的日期范围 */
function getLast7Days(): string[] {
  const today = new Date();
  const dates: string[] = [];
  
  for (let i = 6; i >= 0; i--) {
    const date = new Date(today);
    date.setDate(date.getDate() - i);
    const dateStr = date.toISOString().split('T')[0]; // 格式: YYYY-MM-DD
    dates.push(dateStr);
  }
  
  return dates;
}

/** 获取所有工作空间的每日实体数量统计 */
async function getDailyEntityCounts(): Promise<{ [key: string]: number }> {
  const workspaceResp = await ListKnowledgeGraphWorkspace({ page_size: 1000, page_number: 1 });
  const workspaces: KnowledgeGraphWorkspace[] = workspaceResp.knowledge_graph_workspaces || [];
  const dates = getLast7Days();
  const dailyCounts: { [key: string]: number } = {};
  
  // 初始化每天的计数为0
  dates.forEach(date => {
    dailyCounts[date] = 0;
  });

  // 遍历所有工作空间和任务，统计每天的实体数量
  for (const ws of workspaces) {
    if (ws && typeof ws.id === 'number' && ws.id > 0) {
      try {
        // 获取该工作空间下的所有任务
        const taskResp = await request.post('/extract_task/list', {
          work_space_id: ws.id,
          page_size: 1000,
          page_number: 1,
        });
        
        const tasks = (taskResp as any).extract_tasks || [];
        
        // 统计每个任务的实体数量
        for (const task of tasks) {
          if (task && typeof task.id === 'number' && task.id > 0) {
            try {
              const entityResp: ListEntityResp = await request.post('/entity/list', {
                task_id: task.id,
                page_size: 1000,
                page_number: 1,
              });
              
              // 按任务创建时间统计实体数量
              if (task.created_at) {
                const taskDate = new Date(task.created_at).toISOString().split('T')[0];
                if (dailyCounts.hasOwnProperty(taskDate)) {
                  dailyCounts[taskDate] += entityResp.total || 0;
                }
              }
            } catch (taskError) {
              console.warn(`获取任务 ${task.id} 的实体数量失败:`, taskError);
            }
          }
        }
      } catch (workspaceError) {
        console.warn(`获取工作空间 ${ws.id} 的任务列表失败:`, workspaceError);
      }
    }
  }
  
  return dailyCounts;
}

/** 按日期统计近七天实体新增数量 */
export async function GetEntityDailyCount(): Promise<DailyCountItem[]> {
  try {
    const dates = getLast7Days();
    const dailyCounts = await getDailyEntityCounts();

    // 转换为返回格式
    const result: DailyCountItem[] = dates.map(date => ({
      date,
      count: dailyCounts[date]
    }));

    return result;
  } catch (error) {
    console.error('获取实体每日数量统计失败:', error);
    // 返回默认的7天数据
    const dates = getLast7Days();
    return dates.map(date => ({
      date,
      count: 0
    }));
  }
}

/** 按日期统计近七天实体累计总数量 */
export async function GetEntityTotalDailyCount(): Promise<DailyCountItem[]> {
  try {
    const dates = getLast7Days();
    const dailyCounts = await getDailyEntityCounts();

    // 计算累计总数
    let cumulativeSum = 0;
    const result: DailyCountItem[] = dates.map(date => {
      cumulativeSum += dailyCounts[date];
      return {
        date,
        count: cumulativeSum
      };
    });
    
    return result;
  } catch (error) {
    console.error('获取实体累计数量统计失败:', error);
    // 返回默认的7天数据
    const dates = getLast7Days();
    return dates.map(date => ({
      date,
      count: 0
    }));
  }
}
