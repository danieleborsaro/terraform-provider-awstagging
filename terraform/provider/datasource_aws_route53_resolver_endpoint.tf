locals {
  R53-REI_role_extend = "r53re"

  R53-REI_versioning_sources = [
    "r53re one",
    "r53re two"
  ]
}

data "awstagging_aws_route53_resolver_endpoint" "awsRoute53ResolverEndpointInbound" {
  is_inbound         = true
  role_extend        = local.R53-REI_role_extend
  versioning_sources = local.R53-REI_versioning_sources
}


locals {
  R53-REO_role_extend = "r53re"

  R53-REO_versioning_sources = [
    "r53re one",
    "r53re two"
  ]
}

data "awstagging_aws_route53_resolver_endpoint" "awsRoute53ResolverEndpointOutbound" {
  is_inbound         = false
  role_extend        = local.R53-REO_role_extend
  versioning_sources = local.R53-REO_versioning_sources
}
