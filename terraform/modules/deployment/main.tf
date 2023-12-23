resource "aws_apigatewayv2_deployment" "service" {
  api_id      = var.api_id
  description = "LemniSpace Service deployment"

  triggers = {
    redeployment = sha1(join(",", var.route_hashes))
  }

  lifecycle {
    create_before_destroy = true
  }
}
