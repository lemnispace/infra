variable "api_stage_name" {
  type        = string
  description = "The name of the API stage"
}
variable "api_id" {
  type        = string
  description = "The ID of the API to which to connect the stage"
}
variable "deployment_id" {
  type        = string
  description = "The deployment ID of the API stage"
  default     = null
}
variable "auto_deploy" {
  type        = bool
  description = "Whether to automatically deploy changes to the stage"
  default     = false
}
