variable "services_s3_bucket_id" {
  description = "The S3 bucket to store the Lambda functions in"
  type        = string
}

variable "lambda_role_arn" {
  description = "The ARN of the role to use for the Lambda functions"
  type        = string
}
