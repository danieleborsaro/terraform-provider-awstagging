locals {
  ASG_custom_tags = {
    MyKey3 = "My asg value 3"
    MyKey4 = "My asg value 4"
  }

  ASG_custom_tags_verbatim = {
    MyKeyVerbatim3 = "My datasource value verbatim asg3"
    MyKeyVerbatim4 = "My datasource value verbatim asg4"
  }

  ASG_role_extend = "asg"

  ASG_versioning_sources = [
    "asg one",
    "asg two"
  ]
}


data "awstagging_aws_autoscaling_group" "awsAutoscalingGroup" {
  custom_tags                = local.ASG_custom_tags
  custom_tags_verbatim       = local.ASG_custom_tags_verbatim
  is_propagate_tags_at_lauch = local.is_propagate_tags_at_lauch
  role_extend                = local.ASG_role_extend
  versioning_sources         = local.ASG_versioning_sources
}
