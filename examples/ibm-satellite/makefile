location:
	
	terraform init && terraform get && terraform apply --target=module.create_location --auto-approve

hosts:
	terraform init && terraform get && terraform apply --target=ibm_compute_ssh_key.test_ssh_key  --auto-approve -var 'ssh_public_key=${public_key}'
	
	terraform init && terraform get && terraform apply --target=module.vms --auto-approve 

runssh:
	
	terraform init && terraform get && terraform apply --target=module.run_ssh --auto-approve -var 'private_ssh_key=${private_ssh_key}'

assignhost:	
	terraform init && terraform get && terraform apply --target=module.assign --auto-approve 

registerdns:
	
	terraform init && terraform get && terraform apply --target=module.register_dns --auto-approve 

cluster:
	
	terraform init && terraform get && terraform apply --target=module.create_cluster --auto-approve 
	terraform init && terraform get && terraform apply --target=module.assign_cluster --auto-approve 
	

destroy:
	terraform init && terraform get && terraform destroy --auto-approve