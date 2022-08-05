
variable "summary_name" {
  type    = string
  default = "Vagrante espresso"
}

data "hashicups_defender" "all" {}

# Returns all relays
output "all_summary" {
  value = data.hashicups_defender.all.summary
}

# Only returns packer spiced latte
output "summary" {
  value = {
    for summary in data.hashicups_defender.all.summary :
    summary.relayerId => summary
    if summary.name == var.summary_name
  }
}
