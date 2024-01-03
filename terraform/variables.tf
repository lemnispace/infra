variable "aws_region" {
  description = "AWS region to deploy to"
  default     = "us-east-1"
}

variable "deployment_repo_owner" {
  description = "The owner of the deployment repository"
  type        = string
  default     = "lemnispace"
}

variable "deployment_repo_name" {
  description = "The name of the deployment repository"
  type        = string
  default     = "infra"
}

variable "deployment_file_name" {
  description = "The name of the deployment file"
  type        = string
  default     = "deploy.yaml"
}
