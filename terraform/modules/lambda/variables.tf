variable "services_s3_bucket_id" {
  description = "The S3 bucket to store the Lambda functions in"
  type        = string
}

variable "lambda_role_arn" {
  description = "The ARN of the role to use for the Lambda functions"
  type        = string
}

variable "deployment_repo_owner" {
  description = "The owner of the deployment repository"
  type        = string
}

variable "deployment_repo_name" {
  description = "The name of the deployment repository"
  type        = string
}

variable "deployment_file_name" {
  description = "The name of the deployment file"
  type        = string
}
