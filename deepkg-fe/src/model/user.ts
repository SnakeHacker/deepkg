
export interface User {
    id?: number;
    user_code?: string;
    org_id: number;
    org_name?: string;
    account?: string;
    username?: string;
    password?: string;
    phone?: string;
    mail?: string;
    enable: number;
    role: number;
    avatar?: string;
    expired_at?: string;
    created_at?: string;
    updated_at?: string;
}
