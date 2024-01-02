output "execute_lambda_arn" {
  value = aws_iam_role.execute_lambda_role.arn
}

output "execute_lambda_name" {
  value = aws_iam_role.execute_lambda_role.name
}

output "webhook_lambda_arn" {
  value = aws_iam_role.lemnispace_services_webhook_lambda_role.arn
}

output "webhook_lambda_name" {
  value = aws_iam_role.lemnispace_services_webhook_lambda_role.name
}
