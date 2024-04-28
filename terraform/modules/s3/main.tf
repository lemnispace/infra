# create s3 bucket name
resource "random_pet" "lemnispace_services_bucket_name" {
  length = 4
  prefix = "lemnispace-services-bucket"
}
resource "random_pet" "lemnispace_user_product_files_bucket_name" {
  length = 4
  prefix = "lemnispace-user-product-files-bucket"
}

# create s3 bucket
resource "aws_s3_bucket" "lemnispace_services_bucket" {
  bucket = random_pet.lemnispace_services_bucket_name.id
  tags = {
    Name = "lemnispace-services-bucket"
  }
}
resource "aws_s3_bucket" "lemnsipace_user_product_files_bucket" {
  bucket = random_pet.lemnispace_user_product_files_bucket_name.id
  tags = {
    Name = "lemnispace-user-product-files-bucket"
  }

}

# add bucket ownership controls
resource "aws_s3_bucket_ownership_controls" "lemnispace_services_ownership_controls" {
  bucket = aws_s3_bucket.lemnispace_services_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}
resource "aws_s3_bucket_ownership_controls" "lemnispace_user_product_files_ownership_controls" {
  bucket = aws_s3_bucket.lemnsipace_user_product_files_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

# add bucket acl (access control list)
resource "aws_s3_bucket_acl" "lemnispace_services_acl" {
  depends_on = [aws_s3_bucket_ownership_controls.lemnispace_services_ownership_controls]
  bucket     = aws_s3_bucket.lemnispace_services_bucket.id
  acl        = "private"
}
resource "aws_s3_bucket_acl" "lemnispace_user_product_files_acl" {
  depends_on = [aws_s3_bucket_ownership_controls.lemnispace_user_product_files_ownership_controls]
  bucket     = aws_s3_bucket.lemnsipace_user_product_files_bucket.id
  acl        = "private"
}
