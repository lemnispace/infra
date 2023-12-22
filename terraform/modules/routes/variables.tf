variable "lambda_arn" {
  type        = string
  description = "ARN of the Lambda function to be invoked by the API Gateway"
}

variable "lambda_endpoint" {
  type        = string
  description = "Endpoint of the Lambda function to be invoked by the API Gateway. Do not include a trailing slash. (e.g. '/posts' not '/posts/')"
}

variable "api_stage_name" {
  type        = string
  description = "Name of the API stage to(e.g. 'dev', 'prod')"
}
