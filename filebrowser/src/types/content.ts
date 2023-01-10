export interface CreateContentRequest {
  content_locate_url: string
  managed_contract: string
  network: string
  payment_token_address: string
  payment_token_amount: number
  key_id: number
  content_name: string
  encryption_type: number
  file_extension: string
  description: string
}

export interface CreateContentResponse {
  data: {
    content_id: number
  }
}

export interface ShowContentRequest {
  content_id: number
}

export interface ShowContentResponse {
  data: {
    content_id: number
    content_name: string
    preview_image: Buffer
    description: string
    encryption_type: number
    file_extension: string
    file_size: string
    asset_id: number
    network: string
    payment_token_address: string
    payment_token_amount: number
  }
}
