variable "api_stage_name" {
  type        = string
  description = "The name of the API stage"
}

variable "api_id" {
  type        = string
  description = "The ID of the API to which to connect the stage"
}

variable "route_hashes" {
  type        = list(string)
  description = "The hashes of the routes resources associated with the api stage" # this is used to trigger a redeployment
}
