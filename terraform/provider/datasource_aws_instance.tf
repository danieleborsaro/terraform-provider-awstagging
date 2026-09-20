locals {
  EC2-I_custom_tags = {
    MyKey3 = "My ec2 value 3"
    MyKey4 = "My ec2 value 4"
  }

  EC2-I_custom_tags_verbatim = {
    MyKeyVerbatim3 = "My ec2 value verbatim 3"
    MyKeyVerbatim4 = "My ec2 value verbatim 4"
  }

  EC2-I_platform_name = "Linux"
  EC2-I_region        = "sa-east-1"
  EC2-I_role          = "ec2"

  EC2-I_versioning_sources = [
    "ec2 one",
    "ec2 two"
  ]
}

data "awstagging_aws_instance" "awsInstance" {
  availability_zone    = local.availability_zone
  role                 = local.EC2-I_role
  custom_tags          = local.EC2-I_custom_tags
  custom_tags_verbatim = local.EC2-I_custom_tags_verbatim
  platform_name        = local.EC2-I_platform_name
  region               = local.EC2-I_region
  versioning_sources   = local.EC2-I_versioning_sources
}
