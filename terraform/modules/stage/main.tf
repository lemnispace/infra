resource "aws_apigatewayv2_stage" "lemnispace_services_stage" {
  api_id      = var.api_id
  name        = var.api_stage_name
  description = "${var.api_stage_name} stage for Lemnispace Services API"
  auto_deploy = true
  stage_variables = {
    "stage" = var.api_stage_name
  }
}
