# IBM VSI (virtual machine) example

This example shows how to create bulk virtual machines on IBM Cloud and use data source to retrieve the each vm information. 
This is a minimal configuration to demonstrate the Terraform lifecycle of apply and destroy. 

VSI is configured with private network nic only to protect from threats on the public Internet. 

The bulk_vms can not be updated and also when the bulk_vms is used one cannot update the other parameters like cores, memory, disks, flavor_key_name etc.


To run, configure your IBM Cloud provider

Running the example

For planning phase

```shell
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