import { defineStore } from 'pinia'
import http from '@/plugins/http-common'
import { IpfsCfg, Alive } from '@/types/ipfs'

type NewType = number

export const useIpfsStore = defineStore('ipfs', {
  state: () => ({
    peer_id: '',
    pub_key: '',
    host: 'http://localhost',
    api_port: 5001,
    gateway_port: 8080,
    message: 'Not Connected',
  }),
  getters: {
    ipfsCfg: (state) =>
      `The Ipfsdg is "${state.peer_id}" "${state.pub_key}" "${state.host}" "${state.api_port}" "${state.gateway_port}" .`,
  },
  actions: {
    // 异步更新 message
    // async updateMessage(newMessage: string): Promise<string> {
    //   return new Promise((resolve) => {
    //     setTimeout(() => {
    //       // 这里的 this 是当前的 Store 实例
    //       this.message = newMessage
    //       resolve('Async done.')
    //     }, 8080)
    //   })
    // },
    // sync
    getIpfsCfg() {
      return {
        peer_id: this.peer_id,
        pubkey: this.pub_key,
        host: this.host,
        api_port: Number(this.api_port),
        gateway_port: Number(this.gateway_port),
      }
    },
    setIpfsCfg(
      peer_id: string,
      pub_key: string,
      host: string,
      api_port: number,
      gateway_port: NewType
    ) {
      this.peer_id = peer_id
      this.pub_key = pub_key
      this.host = host
      this.api_port = api_port
      this.gateway_port = gateway_port
    },
    checkAliveSync(
      peer_id: string,
      pub_key: string,
      host: string,
      api_port: number,
      gateway_port: number
    ) {
      this.peer_id = peer_id
      this.pub_key = pub_key
      this.host = host
      this.api_port = api_port
      this.gateway_port = gateway_port
      this.alive()
    },
    alive() {
      var data: IpfsCfg = {
        peer_id: this.peer_id,
        pubkey: this.pub_key,
        host: this.host,
        api_port: Number(this.api_port),
        gateway_port: Number(this.gateway_port),
      }
      console.log(data)
      http
        .post('/api/v1/ipfs/alive', data)
        .then((response: Alive) => {
          console.log(response.data)
          console.log(response.data.is_alive)
          if (response.data.is_alive === true) {
            this.message = `Connected.`
          } else {
            this.message = `Connect failed: ' + "${response.data.is_alive}"`
          }
        })
        .catch((e: Error) => {
          console.log(e)
          this.message = `Connect failed: ' + "${e}"`
        })
    },
  },
})
