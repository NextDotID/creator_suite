import http from '@/plugins/http-common'
import { defineStore } from 'pinia'
import { ListResponse, ListRequest } from '@/types/file'
import { useIpfsStore } from '@/stores'

// class FileService {
//   health(): Promise<any> {
//     return http.get('/health')
//   }
//   list(data: any): Promise<any> {
//     return http.get('/api/v1/file/list', data)
//   }
//   // create_content(data: any): Promise<any> {
//   //   return http.post('/api/test/create', data)
//   // }

//   // get_content(data: any): Promise<any> {
//   //   return http.post('/api/test/get_content', data)
//   // }
//   get(id: any): Promise<any> {
//     return http.get(`/todos/${id}`)
//   }
// }

// export default new FileService()
export const useBrowserStore = defineStore('browser', {
  state: () => ({
    path: './storage',
    folders: [
      {
        name: '1',
        type: 'dirs',
        path: 'storage/1',
        content_id: 1,
        update_time: '2021-09-20 11:00:01',
        children: [
          {
            name: 'mdi-folder-zip-outline',
            type: 'localfile',
            size: '15MB',
            extension: 'zip',
            path: 'storage/1/mdi-folder-zip-outline.zip',
            content_id: 1,
            asset_id: 0,
            key_id: 1,
            location_url: 'ipfs/123456ikjfghdgvsc',
            created_time: '2021-09-20 11:00:01',
            update_time: '2021-09-20 11:00:01',
          },
        ],
      },
      {
        name: '2',
        type: 'dirs',
        path: 'storage/2',
        content_id: 2,
        update_time: '2022-09-20 11:00:01',
        children: [
          {
            name: 'mdi-language-markdown-outline',
            type: 'localfile',
            size: '27KB',
            extension: 'md',
            path: 'storage/2/mdi-markdown.md',
            content_id: 2,
            asset_id: 0,
            key_id: 2,
            location_url: 'ipfs/123456ikjfghdgvsc',
            created_time: '2022-09-20 11:00:01',
            update_time: '2022-09-20 11:00:01',
          },
        ],
      },
      {
        name: '3',
        type: 'dirs',
        path: 'storage/3',
        content_id: 3,
        update_time: '2022-12-20 11:00:01',
        children: [
          {
            name: 'mdi-file-image',
            type: 'localfile',
            size: '27KB',
            extension: 'png',
            path: 'storage/3/mdi-file-image.png',
            content_id: 3,
            asset_id: 0,
            key_id: 3,
            location_url: 'ipfs/123456ikjfghdgvsc',
            created_time: '2022-12-20 11:00:01',
            update_time: '2022-12-20 11:00:01',
          },
        ],
      },
    ],
    files: [
      {
        name: 'mdi-folder-zip-outline',
        type: 'localfile',
        size: '15MB',
        extension: 'zip',
        path: 'storage/1/mdi-folder-zip-outline.zip',
        content_id: 1,
        asset_id: 0,
        key_id: 1,
        location_url: 'ipfs/123456ikjfghdgvsc',
        created_time: '2021-09-20 11:00:01',
        update_time: '2021-09-20 11:00:01',
      },
      {
        name: 'mdi-language-markdown-outline',
        type: 'localfile',
        size: '27KB',
        extension: 'md',
        path: 'storage/2/mdi-markdown.md',
        content_id: 2,
        asset_id: 0,
        key_id: 2,
        location_url: 'ipfs/123456ikjfghdgvsc',
        created_time: '2022-09-20 11:00:01',
        update_time: '2022-09-20 11:00:01',
      },
      {
        name: 'mdi-file-image',
        type: 'localfile',
        size: '27KB',
        extension: 'png',
        path: 'storage/3/mdi-file-image.png',
        content_id: 3,
        asset_id: 0,
        key_id: 3,
        location_url: 'ipfs/123456ikjfghdgvsc',
        created_time: '2022-12-20 11:00:01',
        update_time: '2022-12-20 11:00:01',
      },
    ],
  }),
  actions: {
    // 异步更新 message
    // async load(newMessage: string): Promise<string> {
    //   return new Promise((resolve) => {
    //     setTimeout(() => {
    //       // 这里的 this 是当前的 Store 实例
    //       this.message = newMessage
    //       resolve('Async done.')
    //     }, 8080)
    //   })
    // },
    changePath(p: string) {
      this.path = p
    },
    loadPath() {
      console.log(this.path)
      const ipfsStore = useIpfsStore()
      const cfg = ipfsStore.getIpfsCfg()
      var req: ListRequest = {
        path: this.path,
        cfg: cfg,
      }
      http
        .post('/api/v1/file/list', req)
        .then((response: ListResponse) => {
          console.log(response.data)
          this.folders = response.data.folders
          this.files = response.data.files
        })
        .catch((e: Error) => {
          console.log(e)
        })
    },
  },
})
