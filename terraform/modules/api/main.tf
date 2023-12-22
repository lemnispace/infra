resource "aws_apigatewayv2_api" "lemnispace_services_api" {
  name          = "lemnispace-services-api"
  protocol_type = "HTTP"
  description   = "API for Lemnispace Services"
  version       = "1.0"
}

resource "aws_apigatewayv2_stage" "lemnispace_services_stage" {
  api_id      = aws_apigatewayv2_api.lemnispace_services_api.id
  name        = var.api_stage_name
  description = "${var.api_stage_name} stage for Lemnispace Services API"
  auto_deploy = true
  stage_variables = {
    "stage" = var.api_stage_name
  }
}
