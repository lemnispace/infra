output "stage_name" {
  value = aws_apigatewayv2_stage.lemnispace_services_stage.name
}

output "stage_id" {
  value = aws_apigatewayv2_stage.lemnispace_services_stage.id
}

output "deployment_id" {
  value = aws_apigatewayv2_deployment.lemnispace_services_deployment.id
}

