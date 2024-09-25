terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.68.0"
    }
  }
}

provider "aws" {
  region = "us-east-2"
}

module "new_bucket" {
  source = "./modules/bucket"

  bucket_name                    = var.bucket_name
  bucket_tags                    = var.bucket_tags
  bucket_policy_allow_principals = var.bucket_policy_allow_principals
}