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

This Ansible package and subdirectories can be downloaded separately to the other examples using the subversion command line client (svn). Install Subversion for your workstation from the [Apache.org website](https://subversion.apache.org/packages.html).

Create a temporary directory to download this package in your Ansible project directory.

`mkdir temp`

From your Ansible project directory, download this example using `svn export`. Browse to the github project and subdirectory you want to clone, for example: https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ansible/ibm_ansible_wordpress. Replace tree/master with trunk in the URL, and run svn export on it. 

`svn export https://github.com/IBM-Cloud/terraform-provider-ibm/trunk/examples/ansible/ibm_ansible_wordpress`

Move terraform_inv.py and terraform_inv.ini to the **inventory** folder of the Ansible project. 

`mv ./ibm_ansible_dyn_inv/terraform* ../inventory`

nagivate to the **inventory** directory

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


  
