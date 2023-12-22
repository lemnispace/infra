output "api_id" {
  value = aws_apigatewayv2_api.lemnispace_services_api.id
}

output "api_endpoint" {
  value = aws_apigatewayv2_api.lemnispace_services_api.api_endpoint
}

output "htp_method" {
  value = aws_apigatewayv2_api.lemnispace_services_api.htp_method
}

output "stage_name" {
  value = aws_apigatewayv2_api.lemnispace_services_api.stage_name
}
