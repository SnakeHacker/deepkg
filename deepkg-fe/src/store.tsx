import { create } from 'zustand'

export interface MessageInfo {
    successMsg: string
    errorMsg: string
    warningMsg: string

    success: Function
    warning: Function
    error: Function
}

export interface LoadDoc {
    docList: any[]

    setDocList: Function
    removeDocListItem: Function
    clearDocList: Function
}

export interface LoadTriple {
    tripleList: any[]

    setTripleList: Function
    removeTripleListItem: Function
    clearTripleList: Function
}


export const useStore = create((set) => ({
    successMsg: "",
    errorMsg: "",
    warningMsg: "",

    success: (msg: string) => set(() => ({ successMsg: msg})),
    error: () => set((state: MessageInfo) => ({ errorMsg: state.errorMsg})),
    warning: () => set((state: MessageInfo) => ({ successMsg: state.successMsg})),

    docList: [],
    setDocList:(docs: any[])=> set(() => ({ docList: docs})),
    removeDocListItem: (id: string) => set((state: LoadDoc) => ({ docList: state.docList.filter((item) => item.id !== id)})),
    clearDocList: () => set(() => ({ docList: []})),

    tripleList: [],
    setTripleList:(triples: any[])=> set(() => ({ tripleList: triples})),
    removeTripleListItem: (id: string) => set((state: LoadTriple) => ({ tripleList: state.tripleList.filter((item) => item.id !== id)})),
    clearTripleList: () => set(() => ({ tripleList: []})),
}))