import { IpfsCfg } from '@/types/ipfs'

export interface Folder {
  name: string
  type: string
  path: string
  content_id: number
  created_time: string
  update_time: string
  children: File[]
}

export interface File {
  name: string
  type: string
  size: string
  extension: string
  icon?: string
  path: string

  content_id: number
  asset_id: number
  key_id: number
  location_url: string
  created_time: string
  update_time: string
}

export interface Result {
  folders: Folder[]
  files: File[]
}

export interface ListResponse {
  data: Result
}

export interface ListRequest {
  path: string
  cfg: IpfsCfg
}

export interface CreateRequest {
  encrypt_type: number
  key: string
  origin_file: string
  encrypt_file: string
}

export interface CreateResponse {
  data: {
    key_id: number
    encrypt_file: string
  }
}

export interface MoveRequest {
  src: string
  dst: string
}

export interface CopyRequest {
  src: string
  dst: string
}
