variable "aws_region" {
  description = "AWS region to deploy to"
  default     = "us-east-1"
}
variable "stage_name" {
  type        = string
  description = "The name of the stage to deploy to"
}

variable "lambda_endpoint" {
  type        = string
  description = "The endpoint for the lambda function to be invoked by the API Gateway. Do not include a trailing slash. (e.g. '/posts' not '/posts/')"
  default     = "lemnispace-services"
}

variable "lambda_arn" {
  type        = string
  description = "ARN of the Lambda function to be invoked by the API Gateway"
}

variable "route_hashes" {
  type        = list(string)
  description = "The hashes of the routes resources associated with the api" # this is used to trigger a redeployment
  default     = []
}
