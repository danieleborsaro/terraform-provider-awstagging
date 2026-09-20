locals {
  availability_zone = "eu-west-1a"
  compliance        = "tbc"
  cost_centre       = "n/a"
  custom_tags = {
    MyKey1 = "My value 1"
    MyKey2 = "My value 2"
  }
  custom_tags_verbatim = {
    MyKeyVerbatim1 = "My value verbatim 1"
    MyKeyVerbatim2 = "My value verbatim 2"
  }
  description                = "This is my provider. I am proud of it."
  is_create_before_destroy   = true
  is_force_generated_name    = false
  is_propagate_tags_at_lauch = true
  owner                      = "daniele@borsaro.it"
  region                     = "eu-west-1"
  role                       = "provider"


  account = 811635568629

  accounts_coding = {
    ## Removing leading zero because of terraform issue https://github.com/hashicorp/terraform/issues/28619
    96935932367 = {
      class          = "management"
      name           = "management"
      name_encoded   = "MGM"
      name_canonical = "Management"
    }

    706980398008 = {
      class          = "development"
      name           = "devtest"
      name_encoded   = "DT"
      name_canonical = "DevTest"
    }

    668383449573 = {
      class          = "development"
      name           = "devopsdev"
      name_encoded   = "DVD"
      name_canonical = "DevOpsDev"
    }

    196786728864 = {
      class          = "development"
      name           = "secopsdev"
      name_encoded   = "SD"
      name_canonical = "SecOpsDev"
    }

    703294084643 = {
      class          = "production"
      name           = "prod"
      name_encoded   = "P"
      name_canonical = "UATProduction"
    }

    392399877490 = {
      class          = "preproduction"
      name           = "saraging"
      name_encoded   = "E"
      name_canonical = "UATShared"
    }

    811635568629 = {
      class          = "production"
      name           = "devopsprod"
      name_encoded   = "DVP"
      name_canonical = "DevOpsProd"
    }

    939276390372 = {
      class          = "production"
      name           = "devopsassets"
      name_encoded   = "DVA"
      name_canonical = "DevOpsAssets"
    }
  }

  app_ecosystem       = "dev1"
  app_environment     = "dev1"
  company_name_long   = "My Company Foo"
  company_name_short  = "Foo"
  infra_environment   = "dev1"
  project_name_long   = "My test project"
  project_name_short  = "mytstprjct"
  resource_set_long   = "My test resource set"
  resource_set_short  = "mytstrsrst"
  terraform_module    = basename(abspath(pathexpand(path.root)))
  terraform_workspace = "default"
}
