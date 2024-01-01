output "webhook_arn" {
  value = aws_lambda_function.WebhookFunction.arn
}

output "webhook_invoke_arn" {
  value = aws_lambda_function.WebhookFunction.invoke_arn
}
