
import type { User } from '../model/user';
import request from '../utils/req';

export interface ListUserParams {
  page_size: number;
  page_number: number;
}

export async function ListUser(params: ListUserParams): Promise<any> {
    return request.post(`/user/list`, {
        ...params
    } as any);
}


export interface DeleteUserParams {
    ids: number[];
}

export async function DeleteUsers(params: DeleteUserParams): Promise<any> {
    return request.post(`/user/delete`, params);
}

export interface CreateUserParams {
    user: User;
}

export async function CreateUser(params: CreateUserParams): Promise<any> {
    return request.post(`/user/create`, params);
}


export interface UpdateUserParams {
    user: User;
}

export async function UpdateUser(params: UpdateUserParams): Promise<any> {
    return request.post(`/user/update`, params);
}
