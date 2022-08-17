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
output "all_summary" {
  value = data.defender_summary.all.summary
}

# Only return relays created before
output "summary" {
  value = {
    for summary in data.defender_summary.all.summary :
    summary.relayerId => summary
    if summary.name == var.summary_name
  }
}
