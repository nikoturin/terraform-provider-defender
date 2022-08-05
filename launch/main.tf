terraform {   
	required_providers {     
		defender = {       
		   version = "0.2.4"
		   source = "nikoturin/defender"
	}   
    } 

}  


provider "defender" {}

module "relay" { 
	source = "./summary"
}
output "relay"{

	value = module.relay.summary
}
