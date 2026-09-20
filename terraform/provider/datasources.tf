
data "awstagging_configuration" "config" {}

locals {
  environment = data.awstagging_configuration.config

  resource_tagging = {
    awsAutoscalingGroup                = data.awstagging_aws_autoscaling_group.awsAutoscalingGroup
    awsInstance                        = data.awstagging_aws_instance.awsInstance
    awsKmsAlias                        = data.awstagging_aws_kms_alias.awsKmsAlias
    awsAlb                             = data.awstagging_aws_lb.awsAlb
    awsNlb                             = data.awstagging_aws_lb.awsNlb
    awsLbTargetGroup                   = data.awstagging_aws_lb_target_group.awsLbTargetGroup
    awsCloudwatchLogStream             = data.awstagging_aws_cloudwatch_log_stream.awsCloudwatchLogStream
    awsRoute53ResolverEndpointInbound  = data.awstagging_aws_route53_resolver_endpoint.awsRoute53ResolverEndpointInbound
    awsRoute53ResolverEndpointOutbound = data.awstagging_aws_route53_resolver_endpoint.awsRoute53ResolverEndpointOutbound
  }
}
