terraform {
  required_version = ">= 0.13"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }  
  backend "s3" {
    bucket         = "co.trainingdock.terraformstate"
    key            = "svelte-go-app/terraform.tfstate"
    region         = "us-west-1"
    dynamodb_table = "terraform-locks"
  }
}
provider "aws" {
  region  = "us-west-2"
  default_tags {
    tags = {
      application = "svelte-go-app"
    }
  }
}