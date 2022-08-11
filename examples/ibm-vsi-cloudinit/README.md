# IBM VSI (virtual machine) example with CloudInit configuration

This example shows how to create a virtual machine on IBM Cloud using CloudInit to perform post
boot configuration of the VSI. An Apache (httpd) webserver is installed and started.  

VSI is configured with private network nic only to protect from threats on the public Internet. 
To access the Apache server, start your VPN connection to IBM Cloud and run curl to return the 
Apache splash page. Note the Apache webserver will take a couple of minutes to start after the
Terraform apply completes. 

'curl <vsi private ip address>'


To run, configure your IBM Cloud provider

Running the example

For planning phase

```shell
terraform init
terraform plan
```

For apply phase

```shell
terraform apply
```

For destroy

```shell
terraform destroy
```