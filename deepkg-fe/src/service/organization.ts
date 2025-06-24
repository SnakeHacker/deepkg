import type { Organization } from '../model/organization';
import request from '../utils/req';

export interface ListOrgParams {
  page_size: number;
  page_number: number;
}

export async function ListOrg(params: ListOrgParams): Promise<any> {
    return request.post(`/org/list`, {
        ...params
    } as any);
}

export interface DeleteOrgParams {
    ids: number[];
}

export async function DeleteOrgs(params: DeleteOrgParams): Promise<any> {
    return request.post(`/org/delete`, params);
}

export interface CreateOrgParams {
    organization: Organization;
}

export async function CreateOrg(params: CreateOrgParams): Promise<any> {
    return request.post(`/org/create`, params);
}

export interface UpdateOrgParams {
    organization: Organization;
}

export async function UpdateOrg(params: UpdateOrgParams): Promise<any> {
    return request.post(`/org/update`, params);
}

export interface GetOrgListResp {
  total: number;
  organizations: Organization[];
  page_size: number;
  page_number: number;
}

/**
 * 获取组织机构总数
 * 通过调用分页查询接口，page_size设置为1以减少数据传输量
 * @returns Promise<number> 组织总数
 */
export async function GetOrganizationTotalCount(): Promise<number> {
  const resp: GetOrgListResp = await request.post('/org/list', {
    page_size: 1,
    page_number: 1,
  });

  return resp.total || 0;
}
