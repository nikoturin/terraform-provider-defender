
variable "summary_name" {
  type    = string
  default = "Vagrante espresso"
}

data "nikoturin_defender" "all" {}

# Returns all relays
output "all_summary" {
  value = data.nikoturin_defender.all.summary
}

# Only returns packer spiced latte
output "summary" {
  value = {
    for summary in data.nikoturin_defender.all.summary :
    summary.relayerId => summary
    if summary.name == var.summary_name
  }
}
