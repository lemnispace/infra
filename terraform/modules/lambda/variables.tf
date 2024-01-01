variable "services_s3_bucket_id" {
  description = "The S3 bucket to store the Lambda functions in"
  type        = string
}

variable "execute_lambda_role_arn" {
  description = "The ARN of the role that can execute Lambda functions"
  type        = string
}
