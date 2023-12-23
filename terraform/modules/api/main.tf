resource "aws_apigatewayv2_api" "lemnispace_services_api" {
  name          = "lemnispace-services-api"
  protocol_type = "HTTP"
  description   = "API for Lemnispace Services"
  version       = "1.0"
}
