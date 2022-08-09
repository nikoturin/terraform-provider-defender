terraform {         
	required_providers {                 
		defender = {                    
				version = "0.2.5"                    
				source = "nikoturin/defender"         
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
output "summary" {
  value = {
    for summary in data.defender_summary.all.summary :
    summary.relayerId => summary
    if summary.name == var.summary_name
  }
}
