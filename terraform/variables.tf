variable "aws_region" {
  description = "AWS region to deploy to"
  default     = "us-east-1"
}
variable "stage_name" {
  type        = string
  description = "The name of the stage to deploy to"
}

variable "route_hashes" {
  type        = list(string)
  description = "The hashes of the routes resources associated with the api" # this is used to trigger a redeployment
  default     = []
}
