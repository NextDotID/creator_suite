export interface CreateContentRequest {
  content_locate_url: string
  managed_contract: string
  payment_token_address: string
  payment_token_amount: number
  key_id: number
}

export interface CreateContentResponse {
  data: {
    content_id: number
  }
}
