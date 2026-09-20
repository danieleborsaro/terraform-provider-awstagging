locals {
  LOGS-S_role = "logstrm"

  LOGS-S_versioning_sources = [
    "logstrm one",
    "logstrm two"
  ]
}

data "awstagging_aws_cloudwatch_log_stream" "awsCloudwatchLogStream" {
  role               = local.LOGS-S_role
  versioning_sources = local.LOGS-S_versioning_sources
}
