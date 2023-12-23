terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

module "lemnispace_api" {
  source         = "./modules/api"
  api_stage_name = var.stage_name
}

module "lemnispace_api_stage_deployment" {
  source         = "./modules/stage_deployment"
  api_id         = module.lemnispace_api.api_id
  api_stage_name = var.stage_name
}

module "lemnispace_roles" {
  source = "./modules/roles"
}

module "lemnispace_services_s3" {
  source = "./modules/s3"
}
