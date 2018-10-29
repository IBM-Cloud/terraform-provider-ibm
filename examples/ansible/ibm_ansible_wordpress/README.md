# Wordpress_ansible_ibmcloud

Demo ansible package to install Wordpress in a highly available 3-tier configuration on IBM Cloud
 - IBM Cloud Load Balancer
 - httpd app server
 - mariadb 

This is written as a capability demonstration of building high availability web sites using IBM Cloud IaaS and Terraform.


This pacakge supports two deployment options: 
- Single site deployment of multiple httpd webservers with a single Mariadb database host with an IBM Cloud Load Balancer (CLB) as a local LB. 
- Dual site high availability configuration, with webservers and DB's deployed in two data centers, each with CLBs, fronted by IBM Cloud Internet Services (CloudFlare) as a global load balancer. The Wordpress database is replicated master-master over the IBM Cloud private network. 

Deployment architecture is determined dynamically based on the Ansible inventory file specifying two Mariadb servers in different data centers. The inventory file can be statically specified with manual deployment of hosts on IBM Cloud, or used with Ansible dynamic inventory with Terraform automated deployment of servers and LBs.  


<p style="text-align: center;">
  <img src="images/WordpressCLB.png" alt="CLB single site" width="500"/>
</p>

<p style="text-align: center;">
  <img src="images/WordpressGLB.png" alt="GLB dual site" width="600"/>
</p>


This package uses three excellent Ansible roles maintained by Bert Van Vreckem and team. 
https://github.com/bertvv/ansible-role-mariadb
https://github.com/bertvv/ansible-role-wordpress
https://github.com/bertvv/ansible-role-httpd

DB replication setup is based on the work of Vitalii Michailovich
https://github.com/VitaliiMichailovich/Ansible-MySQL-Master-Master


All submodules/roles are included in this package as a small number of changes where made to support usage in this configuration. Its not ideal, but simplifies installation for users who are new to IBM Cloud, Ansible and Terraform. 

## Local setup
This package was developed on OSX and as such requires sudo rights to execute some of the updates performed to the host file on the OSX control workstation and install modules for monitoring the state of the application. The OSX user password is saved as the encrypted variable 'su_password' in an Ansible Vault file in group_vars/control. The vault password is expected to be stored in the users home directory  ~/vault_pass.txt. ansible.cfg in the root of the package defines the location of the vault password. 


## Inventory
Ansible inventory is defined in /inventory/hosts.

When used with a manually deployed environment, host details take the form:
`app101  ansible_host=10.72.58.78 ansible_user=root`

Alternatively an Ansible dynamic inventory script can be used. 

## IBM Cloud deployment
The required infrastructure for this play can be deployed manually or as intended using Terraform. These approaches are documented in two different tutorials on IBM Cloud Docs.

- Software Defined Network tutorial - The target environment is deployed using the IBM Cloud Terraform github example https://github.com/IBM-Cloud/terraform-provider-ibm/examples/ibm-website-single-region/. Full details of how to deploy Wordpress with this play can be found in the IBM Cloud [Terraform and Ansible](https://console.bluemix.net/docs/terraform/manage_resources.html) documentation. The Ansible Dynamic inventory solution https://github.com/IBM-Cloud/terraform-provider-ibm/examples/ansible/ibm_ansible_dyn_inv/ is a prerequisite. 



- Classic network infrastructure tutorial - The target deployment environment is deployed mnaually and is documented in the IBM Cloud Solution Tutorial [Web application serving from a secure private network](https://console.bluemix.net/docs/tutorials/web-app-private-network.html#web-application-serving-from-a-secure-private-network). The infrastructure deployment from this tutorial is manually completed, with host names and IP addresses manually populated into the Ansible inventory file. 



## Installation

This Ansible package and subdirectories can be downloaded separately to the other examples using the subversion command line client (svn). Install Subversion for your workstation from the [Apache.org](https://subversion.apache.org/packages.html) website. 

From your Ansible project directory, download this example using `svn export`. Browse to the project and subdirectory you want to clone, for example: https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ansible/ibm_ansible_wordpress. Replace `tree/master` with `trunk` in the URL, and run svn export on it, for example:

`svn export https://github.com/IBM-Cloud/terraform-provider-ibm/trunk/examples/ansible/ibm_ansible_wordpress`

nagivate to **ibm_ansible_wordpress/inventory** directory

`cd ibm_ansible_wordpress/inventory`

Update permissions of terraform_inv.py to include **execute**

`chmod +x terraform_inv.py`

## Configuration

Update the terraform_inv.ini file to point to the terraform.tfstate file of the Terraform configuration that Wordpress will be deployed to. The tfstate file should be referenced by is full path and file name. 

```
[TFSTATE]
#TFSTATE_FILE = /nnn/nnn/nnn/tr_test_files/terraform.tfstate
#TFSTATE_FILE = /Users/JohnDoe/terraform/ibm/app2x/terraform.tfstate
#TFSTATE_FILE = /usr/share/terraform/ibm/app2x/terraform.tfstate
``` 

Configure Ansible Vault for secure storage of your local user credentials. This package was developed on OSX and as such requires sudo rights to execute some of the updates performed to the host file on the OSX control workstation and to install modules for monitoring the state of the application. The OSX user password is saved as the encrypted variable su_password in an Ansible Vault file in the group_vars/control directory. The vault password is expected to be stored in the users home directory, ~/vault_pass.txt. The Ansible configuration file ansible.cfg in the root of the package defines the location of the vault password.

See the Ansible [Vault documentation](https://docs.ansible.com/ansible/2.6/user_guide/vault.html) for how to encrypt the control workstation password and store as a vault file. If this package is executed on Linux from a user with sudo rights, this step can be ignored.

## Execution
 
```
ansible-playbook -i inventory site.yml
```

