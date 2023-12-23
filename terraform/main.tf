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

data "terraform_remote_state" "lemnispace_mosaic_service" {
  backend = "s3"
  config = {
    bucket         = "lemnispace-terraform-state"
    key            = "mosaic-service/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-state-lock"
  }
}

module "lemnispace_api" {
  source = "./modules/api"
}

module "lemnispace_dev_deployment" {
  source       = "./modules/deployment"
  api_id       = module.lemnispace_api.api_id
  route_hashes = [data.terraform_remote_state.lemnispace_mosaic_service.outputs.mosaic_route_hash]
}

module "lemnispace_api_dev_stage" {
  source         = "./modules/stage"
  api_id         = module.lemnispace_api.api_id
  api_stage_name = "Dev"
  deployment_id  = module.lemnispace_dev_deployment.deployment_id
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
