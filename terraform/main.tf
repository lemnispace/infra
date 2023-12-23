terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket         = "lemnispace-terraform-state"
    key            = "infra/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-lock"
  }
}

provider "aws" {
  region = var.aws_region
}

module "lemnispace_api" {
  source = "./modules/api"
}

module "lemnispace_api_dev_stage" {
  source         = "./modules/stage"
  api_id         = module.lemnispace_api.api_id
  api_stage_name = "Dev"
}
module "lemnispace_api_prod_stage" {
  source         = "./modules/stage"
  api_id         = module.lemnispace_api.api_id
  api_stage_name = "Prod"
}

module "lemnispace_roles" {
  source = "./modules/roles"
}

module "lemnispace_services_s3" {
  source = "./modules/s3"
}
