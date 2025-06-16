import { create } from 'zustand'

export interface MessageInfo {
    successMsg: string
    errorMsg: string
    warningMsg: string

    success: Function
    warning: Function
    error: Function
}


export const useStore = create((set) => ({
    successMsg: "",
    errorMsg: "",
    warningMsg: "",

    success: (msg: string) => set(() => ({ successMsg: msg})),
    error: () => set((state: MessageInfo) => ({ errorMsg: state.errorMsg})),
    warning: () => set((state: MessageInfo) => ({ successMsg: state.successMsg})),
}))