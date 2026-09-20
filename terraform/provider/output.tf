output "configuration" {
  value = local.environment
}

output "resources" {
  value = { for k, v in local.resource_tagging : k => {
    name         = v.name
    name_prefix  = v.name_prefix
    name_hash    = v.name_hash
    tags         = v.tags
    tags_as_maps = v.tags_as_maps
    type_name    = v.type_name
    type_prefix  = v.type_prefix
    }
  }
}
