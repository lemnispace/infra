### Webhook Deployment Lambda Function ###
data "archive_file" "WebhookFunction" {
  type        = "zip"
  source_dir  = "${path.root}/../build/deploy/bootstrap"
  output_path = "${path.root}/build/deploy/WebhookFunction.zip"
}

resource "aws_s3_object" "webhook_deploy" {
  bucket = var.services_s3_bucket_id
  key    = "WebhookFunction.zip"

  source = data.archive_file.WebhookFunction.output_path
  etag   = filemd5(data.archive_file.WebhookFunction.output_path)
}

resource "aws_lambda_function" "WebhookFunction" {
  filename         = data.archive_file.WebhookFunction.output_path
  function_name    = "WebhookFunction"
  role             = var.execute_lambda_role_arn
  handler          = "bootstrap"
  runtime          = "provided.al2023"
  architectures    = ["x86_64"]
  source_code_hash = data.archive_file.WebhookFunction.output_base64sha256
  timeout          = 30
  memory_size      = 512
}
