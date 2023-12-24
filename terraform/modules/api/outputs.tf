output "api_id" {
  value = aws_apigatewayv2_api.lemnispace_services_api.id
}

output "api_endpoint" {
  value = aws_apigatewayv2_api.lemnispace_services_api.api_endpoint
}

output "api_execution_arn" {
  value = aws_apigatewayv2_api.lemnispace_services_api.execution_arn
}
