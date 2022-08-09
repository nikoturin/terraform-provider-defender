terraform {         
	required_providers {                 
		defender = {                    
				source = "nikoturin/defender"
				version = ">=0.2.5"
		}
	}  
}

provider "defender" {}

variable "summary_name" {
  type    = string
  default = "Vagrante espresso"
}

data "defender_summary" "all" {}

# Returns all relays
output "all_summary" {
  value = data.defender_summary.all.summary
}

# Only returns packer spiced latte
output "summar" {
  value = {
    for summar in data.defender_summary.all.summary :
    summar.relayerId => summary
    if summar.name == var.summar_name
  }
}
