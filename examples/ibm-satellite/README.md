# terraform-satellite

Use this project to set up satellite on IBM Cloud, using Terraform.

## Overview
Deployment of 'Satelite on IBM Cloud' is divided into separate steps.
	
* Step 1: You can create Satellite locations for each place that you like, such as your company's ports in the north and south of the country. After you set up your locations, you can bring IBM Cloud services and consistent application management to the machines that already exist in your environments in the location.
  
* Step 2: After you create the location, you must add compute capacity to your location so that you can run the Satellite control plane or set up OpenShift clusters.<br>
you can add hosts that you run in your on-prem data center, in IBM Cloud, or in other cloud providers. Make sure that your host has public network connectivity and that you have access to the host machine to run the Satellite script.
  
* Step 3: Log in to each host machine that you want to add to your location and run the script. The steps for how to log in to your machine and run the script vary by cloud provider. When you run the script on the machine, the machine is made visible to your Satellite location, but is not yet assigned to the Satellite control plane or a Satellite cluster. The script also disables the ability to SSH in to the machine for security purposes. If you later remove the host from the Satellite location, you must reload the host machine to SSH into the machine again

* Step 4: Setup Satellite Control Plane. The Satellite control plane serves as the .... TODO. To create the control plane, you must add at least 3 compute hosts to your location that meet the [minimum requirements](https://test.cloud.ibm.com/docs/satellite?topic=satellite-limitations#limits-host) . Assign these hosts to your location.

*Step 5: Create DNS for your new Location.
 

## Prerequisite

* Set up the IBM Cloud command line interface (CLI), the Satellite plug-in, and other related CLIs.
* Install cli and plugin package

  ``` console
    ibmcloud login -a https://test.cloud.ibm.com -r $region --apikey=$API_KEY
    ibmcloud ks init --host https://containers.test.cloud.ibm.com
    ibmcloud plugin repo-add stage https://plugins.test.cloud.ibm.com
    ibmcloud plugin install container-service -r stage

  ```

## Steps to execute

### 1. Setup the IBM terraform-satellite Project


* Generate the private and public key pair which is required to provision the   virtual machines in softlayer.(Put the private key inside ~/.ssh/id_rsa).Follow the instruction [here](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/) to generate ssh key pair

* Paste the private key file in the root directory of this project.


### 2. Provision satellite location for IBM Cloud Infrastructure.

* Update variables.tf file 

* Provision the infrastructure using the following command.Expected time to execute this command is `13 min`
    ```
    make location
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.   
   ```

check for the available satellite locations using
```
ibmcloud sat location ls
```

### 3. Provision hosts for satellite location on IBM Cloud Infrastructure.

* Update variables.tf file 
* A Minimum of 3 hosts are to be assigned for the location to be in normal state. Here we create 4 hosts. 3 for location to be assigned and one for staellite cluster
* Provision the infrastructure using the following command.Expected time to execute this command is `10 min`
    ```
      make hosts
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 4 added, 0 changed, 0 destroyed.
   ```

### 4. Run SSH on the hosts

* Update variables.tf file 

* Run SSH using the following command.Expected time to execute this command is `3 min`
    ```
      make runssh
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 5 added, 0 changed, 0 destroyed.
   ```
check for the hosts in the satellite using 
```
ibmcloud host ls --location <location name_or_id>
```

### 5. Assign hosts to satellite location

* Update variables.tf file 

* assign hosts using the following command.Expected time to execute this command is `14 min`
    ```
    make assignhost
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 4 added, 0 changed, 0 destroyed.
   ```

### 6. Register DNS

* Update variables.tf file 

* To register DNS loaction should be in ready state. Location will be ready only if the 3 hosts are assigned to it.Expected time to execute this command is `30 min`
    ```
    make registerdns
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
   ```

### 6. Create satellite cluster and assign it to host

* Update variables.tf file 
Expected time to execute this command is 30 min
    ```
    make cluster
    ```

On successful completion, you will see the following message.
   ```
   ...

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
   ```

Once the setup is complete, you can:

Using your ROKS Cluster via CLI


Easiest way to get started is to grab the admin kubeconfig via the CLI. Running both of these commands ensures your user gets created for console access as well. Note that for now you will need to use --insecure-skip-tls-verify=true or get the lets encrypt non prod CAs added to your system (more details soon).
  ``` console
$ ibmcloud ks cluster config -c <your ROKS cluster id>
# OR fetch the admin kubeconfig
$ ibmcloud ks cluster config --admin -c <your ROKS cluster id>
```
## References 
* [Getting started with IBM Cloud Satellite](https://cloud.ibm.com/docs/satellite?topic=satellite-getting-started)
* [IBM Cloud Satellite CLI Docs](https://cloud.ibm.com/docs/satellite?topic=satellite-satellite-cli-reference)