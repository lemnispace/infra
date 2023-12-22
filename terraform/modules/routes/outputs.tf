output "api_id" {
  value = module.lemnispace_services_api.api_id
}

output "api_endpoint" {
  value = module.lemnispace_services_api.api_endpoint
}

output "api_stage_name" {
  value = module.lemnispace_services_api.stage_name
}

output "route_id" {
  value = aws_apigatewayv2_route.lemnispace_services_route.id
}

output "route_uri" {
  value = aws_apigatewayv2_route.lemnispace_services_route.target
}
