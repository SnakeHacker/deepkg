import request from '../utils/req';
export const uploadFile = (file: any, callback: Function) => {

    const formData = new FormData()
    formData.append('file', file)

    request
      .post('/file/upload', formData)
      .then((res: any) => {
        console.log(res);
        callback(res)
      })
      .catch((err: any) => {
        console.log(err)
      })

};
