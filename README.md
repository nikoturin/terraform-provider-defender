tag: 0.2.11

Steps to start the project Terraform - Defender

1.- Install binary file Terraform

2.- go version go1.18.4

3.- git clone https://github.com/nikoturin/terraform-provider-defender

4.- execute at CLI: go mod init terraform-provider-defender

5.- execute at CLI: go mod tidy

6.- execute at CLI: go build -o terraform-provider-defender

7.- create directory "mkdir -p /root/.terraform.d/plugins/github.com/def/defender/0.2/linux_amd64"

8.- cp -rvfp terraform-provider-defender /root/.terraform.d/plugins/github.com/def/defender/0.2/linux_amd64

8.- export DEFENDER_APIKEY=XXXXXXXXX

9.- export DEFENDER_TOKEN='Bearer XXXXXXXXXXX'

10.- cd launch && terraform init && terraform apply --auto-approve

11.- Check defender webPage to check relay created.

12.- check provider Terraform as "https://registry.terraform.io/providers/nikoturin/defender/0.2.11"

NOTE: Remember to create your apiKeys at defender webPage as Admin.

To create Token I share yout the repository oz-aws from blockdemy:https://github.com/blockdemy/oz-defender-bootcamp

I know is missing some actions from defender, but it's good to start to make some practice devops across Terraform - Defender

I'll continue working with source code to add all actions defender.

:) I hope you can run it
