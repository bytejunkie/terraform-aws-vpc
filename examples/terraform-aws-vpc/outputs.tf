output "arn" {
  value = module.vpc.vpc.arn
}
output "id" {
  value = module.vpc.vpc.id
}
output "cidr_block" {
  value = module.vpc.vpc.cidr_block
}
output "instance_tenancy" {
  value = module.vpc.vpc.instance_tenancy
}
output "enable_dns_support" {
  value = module.vpc.vpc.enable_dns_support
}
output "enable_dns_hostnames" {
  value = module.vpc.vpc.enable_dns_hostnames
}
output "enable_classiclink" {
  value = module.vpc.vpc.enable_classiclink
}
output "enable_classiclink_dns_support" {
  value = module.vpc.vpc.enable_classiclink_dns_support
}
output "ipv6_association_id" {
  value = module.vpc.vpc.ipv6_association_id
}
output "ipv6_cidr_block" {
  value = module.vpc.vpc.ipv6_cidr_block
}
output "default_network_acl_id" {
  value = module.vpc.vpc.default_network_acl_id
}
output "default_security_group_id" {
  value = module.vpc.vpc.default_security_group_id
}
output "default_route_table_id" {
  value = module.vpc.vpc.default_route_table_id
}
output "owner_id" {
  value = module.vpc.vpc.owner_id
}
output "tags" {
  value = module.vpc.vpc.tags
}
