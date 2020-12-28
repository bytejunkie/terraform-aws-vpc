module "vpc" {
  source = "../../"

  cidr_block         = var.cidr_block
  tags               = var.tags
  enable_dns_support = var.enable_dns_support
  # enable_classiclink = var.enable_classiclink # not available
}
