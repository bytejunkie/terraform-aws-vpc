resource "aws_vpc" "main" {
  cidr_block       = var.cidr_block
  instance_tenancy = var.instance_tenancy

  enable_dns_support = var.enable_dns_support

  tags = merge({
    Module-Name = "AWS VPC"
    Author      = "bytejunkie - matt@bytejunkie.dev"
    },
    var.tags
  )
}
