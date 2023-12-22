module "lemnispace_services_api" {
  source         = "../api"
  api_stage_name = var.api_stage_name
}

resource "aws_apigatewayv2_integration" "lemnispace_services_integration" {
  api_id                    = module.lemnispace_services_api.id
  description               = "Lambda integration for Lemnispace Services API"
  integration_type          = "AWS_PROXY"
  integration_method        = "POST"
  connection_type           = "INTERNET"
  content_handling_strategy = "CONVERT_TO_TEXT"
  passthrough_behavior      = "WHEN_NO_MATCH"
  integration_uri           = var.lambda_arn
}

resource "aws_apigatewayv2_route" "lemnispace_services_route" {
  api_id    = module.lemnispace_services_api.id
  route_key = "ANY /${var.lambda_endpoint}/{proxy+}"
}
