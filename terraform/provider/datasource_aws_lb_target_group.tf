locals {
  LB-TG_role_extend = "tgt"

  LB-TG_versioning_sources = [
    "tg one",
    "tg two"
  ]
}

data "awstagging_aws_lb_target_group" "awsLbTargetGroup" {
  role_extend        = local.LB-TG_role_extend
  versioning_sources = local.LB-TG_versioning_sources
}
