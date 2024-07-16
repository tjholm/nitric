variable "bucket_name" {
  description = "The name of the bucket to create"
  type        = string
}

variable "storage_account_name" {
  description = "The name of the storage account to create the bucket in"
  type        = string
}

variable "bucket_subcribers" {
  description = "The list of subscribers to the bucket"
  type = map(object({
    host_url      = string
    event_token   = string
    sp_client_id  = string
    sp_tenant_id  = string
    prefix_filter = string
    event_types   = list(string)
  }))
}
