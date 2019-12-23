# IBM VSI (virtual machine) example

This example shows how to create virtual machines on IBM Cloud. 
This is a minimal configuration to demonstrate the Terraform lifecycle of apply and destroy. 

VSI is configured with private network nic only to protect from threats on the public Internet. 


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