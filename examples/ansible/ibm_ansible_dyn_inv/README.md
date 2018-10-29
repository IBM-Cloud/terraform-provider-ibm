# Terraform-Ansible dynamic inventory for IBM Cloud

Sample Python script to dynamically parse terraform.tfstate file into Ansible inventory

Copyright (c) 2018, IBM UK
steve_strutt@uk.ibm.com
ti_version = '0.6'

## Ansible dynamic inventory for Terraform with IBM Cloud ##
This dynamic inventory script is written for use with Ansible and Terraform on IBM Cloud. Details of how to setup the script can be found in the IBM Cloud [Terraform and Ansible](https://console.bluemix.net/docs/terraform/manage_resources.html) documentation.


## Static and dynamic inventory
This script can be used alongside static inventory files in the same directory 


## Terraform dependencies

This inventory script expects to find Terraform tags of the form 
group: host_group associated with each tf instance to define the 
host group membership for Ansible. Multiple group tags are allowed per host

## Installation

Create directory **inventory** under Ansible project directory
`mkdir inventory` 
Create directory **temp** under Ansible project directory
Navigate to **temp** directory 
`cd temp`
Download files to **temp** directory using `svn export` (subversion)
`svn export https://github.com/stevestrutt/terraform-provider-ibm/trunk/examples/ansible/ibm_ansible_dyn_inv
Move terraform_inv.py and terraform_inv.ini to **inventory** folder. 
'mv ./ibm_ansible_dyn_inv/terraform* ../inventory`
nagivate to **inventory** directory
`cd inventory`
Update permissions of terraform_inv.py to include **execute**
`chmod +x terraform_inv.py`

## Configuration

The terraform.tfstate defining the current state of the deployed infrastructure, 
is read from the Terraform direcfory. This is specified using the  
terraform_inv.ini file in the same directory as the terraform_inv.py script, pointing to the 
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
-  With supplied test files - `./terraform_inv.py -t ../temp/ibm_ansible_dyn_inv/tr_test_files/terraform.tfstate` 
-  With ini file `./terraform.py --list` 
Successful execution returns groups with lists of hosts and _meta/hostvars with a detailed
host listing. 
Validate successful operation with ansible from the playbook directory:
-   `ansible-inventory -i inventory --list`


  
