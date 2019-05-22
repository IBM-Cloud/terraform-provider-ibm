# Terraform-Ansible dynamic inventory for IBM Cloud

Sample Python script to dynamically parse terraform.tfstate file into Ansible inventory

Copyright (c) 2018, IBM UK
steve_strutt@uk.ibm.com
ti_version = '0.8'

## Ansible dynamic inventory for Terraform with IBM Cloud ##
This dynamic inventory script is written for use with Ansible and Terraform on IBM Cloud. Details of how to setup the script can be found in the IBM Cloud [Using Ansible to automate app deployment on Terraform-provided infrastructure](https://console.bluemix.net/docs/terraform/ansible/ansible.html#ansible).


## Static and dynamic inventory
This script can be used alongside static inventory files in the same directory 


## Terraform dependencies

This inventory script expects to find Terraform tags of the form 
group: host_group associated with each tf instance to define the 
host group membership for Ansible. Multiple group tags are allowed per host

This script was written for Terraform 0.11.07. It will break if Hashicorp change the format of the terraform.tf file. This is expected with 0.12.0. 

## Configuration

The terraform.tfstate defining the current state of the deployed infrastructure, 
is read from the Terraform direcfory. This is specified using the  
terraform_inv.ini file in the same directory as this script, pointing to the 
location of the terraform.tfstate file to be inventoried. 
The tfstate file should be referenced by is full path and file name. 

```
[TFSTATE]
#TFSTATE_FILE = /nnn/nnn/nnn/tr_test_files/terraform.tfstate
#TFSTATE_FILE = /Users/JohnDoe/terraform/ibm/app2x/terraform.tfstate
#TFSTATE_FILE = /usr/share/terraform/ibm/app2x/terraform.tfstate
``` 
 
## Testing  
 
Validate correct execution:
-  With supplied test files - `./terraform_inv.py -t ../tr_test_files/terraform.tfstate` 
-  With ini file `./terraform.py --list` 
Successful execution returns groups with lists of hosts and _meta/hostvars with a detailed
host listing. 
Validate successful operation with ansible from the playbook directory:
-   `ansible-inventory -i inventory --list`


  
