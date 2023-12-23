
# API Gateway outputs
output "api_endpoint" {
  value = module.lemnispace_api.api_endpoint
}

output "api_id" {
  value = module.lemnispace_api.api_id
}

output "api_stage_name" {
  value = module.lemnispace_api_stage_deployment.stage_name
}

# API Gateway Stage Deployment outputs
output "api_stage_id" {
  value = module.lemnispace_api_stage_deployment.stage_id
}

output "api_deployment_id" {
  value = module.lemnispace_api_stage_deployment.deployment_id
}

# IAM Role outputs
output "execute_lambda_role_arn" {
  value = module.lemnispace_roles.execute_lambda_arn
}

output "execute_lambda_role_name" {
  value = module.lemnispace_roles.execute_lambda_name
}

# S3 Bucket outputs
output "services_s3_bucket_id" {
  value = module.lemnispace_services_s3.bucket_id
}

output "services_s3_bucket_arn" {
  value = module.lemnispace_services_s3.bucket_arn
}
