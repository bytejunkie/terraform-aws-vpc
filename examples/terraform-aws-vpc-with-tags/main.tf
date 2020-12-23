module "vpc" {
  source = "../../"

  cidr_block = var.cidr_block
  tags = {
    "Name" = "My_VPC"
  }
}
