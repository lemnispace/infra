### EXECUTE LAMBDA ROLE ###
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "execute_lambda_role" {
  name               = "execute_lambda_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.execute_lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

### WEBHOOK LAMBDA ROLE ###
data "aws_iam_policy_document" "ssm_parameter_access" {
  statement {
    effect = "Allow"

    actions = [
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:GetParametersByPath"
    ]

    resources = [
      "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/*"
    ]
  }
}

resource "aws_iam_policy" "lambda_ssm_access_policy" {
  name   = "lambda_ssm_access_policy"
  policy = data.aws_iam_policy_document.ssm_parameter_access.json
}

resource "aws_iam_role" "lemnispace_services_webhook_lambda_role" {
  name               = "lemnispace_services_webhook_lambda_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "ssm_policy_attachment" {
  role       = aws_iam_role.lemnispace_services_webhook_lambda_role.name
  policy_arn = aws_iam_policy.lambda_ssm_access_policy.arn
}
