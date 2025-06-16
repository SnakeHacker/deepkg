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
