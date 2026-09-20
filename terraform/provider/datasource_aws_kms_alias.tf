locals {
  KMS-A_custom_name = "alias/mykey"
  KMS-A_role        = "kms-alias"

  KMS-A_versioning_sources = [
    "kmsa one",
    "kmsa two"
  ]
}

data "awstagging_aws_kms_alias" "awsKmsAlias" {
  custom_name        = local.KMS-A_custom_name
  role               = local.KMS-A_role
  versioning_sources = local.KMS-A_versioning_sources
}
