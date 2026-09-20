locals {
  LB-ALB_role_extend = "alb"

  LB-ALB_versioning_sources = [
    "alb one",
    "alb two"
  ]
}

data "awstagging_aws_lb" "awsAlb" {
  is_application_load_balancer = true
  role_extend                  = local.LB-ALB_role_extend
  versioning_sources           = local.LB-ALB_versioning_sources
}

locals {
  LB-NLB_role_extend = "nlb"

  LB-NLB_versioning_sources = [
    "nlb one",
    "nlb two"
  ]
}

data "awstagging_aws_lb" "awsNlb" {
  is_application_load_balancer = false
  role_extend                  = local.LB-NLB_role_extend
  versioning_sources           = local.LB-NLB_versioning_sources
}
