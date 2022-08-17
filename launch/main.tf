terraform {         
	required_providers {                 
		defender = {                    
				version = "0.2"
				source = "github.com/def/defender"
		}
	}  
}

provider "defender" {}

variable "summary_name" {
  type    = string
  default = "Missing a lot actions but, it is the initial code :)"
}

data "defender_summary" "all" {}

# Returns all relays
output "all_items" {
  value = data.defender_summary.all.items
}

# Only returns summary relays created before
output "summary" {
  value = {
    for summary in data.defender_summary.all.items :
	    summary.relayerId => summary
    if summary.name == var.summary_name
  }
}

resource "defender_modules" "crud" {
	relay {
		name = "tf-tnx-openzeppelin"
		network = "rinkeby"
		minbalance = "1000000000"
		pendingtxcost = "0"
	}
}

output "crud_modules" {

	value = defender_modules.crud
}
