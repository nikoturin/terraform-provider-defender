
provider "defender" {
	
	api_version = "v1"
	hostname = "localhost"

}

module "psl" {
  source = "./summary"

  summary_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.summary
}
