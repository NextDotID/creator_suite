export interface Result {
  is_alive: boolean
  message: string
}
export interface Alive {
  data: Result
}

export interface IpfsCfg {
  peer_id: string
  pubkey: string
  host: string
  api_port: number
  gateway_port: number
}

export interface IpfsUpload {
  local_file: string
  cfg: IpfsCfg
}

export interface IpfsUploadResult {
  data: {
    content_locate_url: string
  }
}
