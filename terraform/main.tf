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

module "lemnispace_api_stage_deployment" {
  source         = "./modules/stage_deployment"
  api_id         = module.lemnispace_api.api_id
  api_stage_name = var.stage_name
  route_hashes   = var.route_hashes
}

module "lemnispace_roles" {
  source = "./modules/roles"
}

module "lemnispace_services_s3" {
  source = "./modules/s3"
}
