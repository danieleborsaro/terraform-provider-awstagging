terraform {
  required_providers {

    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.65.0"
    }

    awstagging = {
      # source = "localdev/foo/awstagging"

      # source = "github.com/danieleborsaro/awstagging"
      # version = "0.0.0-SNAPSHOT-64bf0c4"

      source  = "danieleborsaro/awstagging"
      version = "0.1.0"
    }
  }
}

provider "awstagging" {
  compliance               = local.compliance
  cost_centre              = local.cost_centre
  custom_tags              = local.custom_tags
  custom_tags_verbatim     = local.custom_tags_verbatim
  description              = local.description
  is_create_before_destroy = local.is_create_before_destroy
  is_force_generated_name  = local.is_force_generated_name
  owner                    = local.owner
  region                   = local.region
  role                     = local.role


  account         = local.account
  accounts_coding = local.accounts_coding

  app_ecosystem      = local.app_ecosystem
  app_environment    = local.app_environment
  company_name_long  = local.company_name_long
  company_name_short = local.company_name_short

  infra_environment  = local.infra_environment
  project_name_long  = local.project_name_long
  project_name_short = local.project_name_short
  resource_set_long  = local.resource_set_long
  resource_set_short = local.resource_set_short

  terraform_module    = local.terraform_module
  terraform_workspace = local.terraform_workspace
}
