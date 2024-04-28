output "services_bucket_id" {
  value = aws_s3_bucket.lemnispace_services_bucket.id
}

output "services_bucket_arn" {
  value = aws_s3_bucket.lemnispace_services_bucket.arn
}

output "user_product_files_bucket_id" {
  value = aws_s3_bucket.lemnsipace_user_product_files_bucket.id
}

output "user_product_files_bucket_arn" {
  value = aws_s3_bucket.lemnsipace_user_product_files_bucket.arn
}
