

resource "aws_apigatewayv2_stage" "lemnispace_services_stage" {
  api_id      = var.api_id
  name        = var.api_stage_name
  description = "${var.api_stage_name} stage for Lemnispace Services API"
  auto_deploy = true
  stage_variables = {
    "stage" = var.api_stage_name
  }
  deployment_id = aws_apigatewayv2_deployment.lemnispace_services_deployment.id
}

resource "aws_apigatewayv2_deployment" "lemnispace_services_deployment" {
  api_id      = var.api_id
  description = "Deployment for Lemnispace Services API"

  triggers = {
    redeployment = sha1(join(",", var.route_hashes))
  }

  lifecycle {
    create_before_destroy = true
  }
}
