# IBM VSI (virtual machine) example

This example shows how to create virtual machines on IBM Cloud using datacenter_choice. you can retry to create a VM instance using a datacenter_choice. If VM fails to place order on first datacenter or vlans it retries to place order on subsequent datacenters and vlans untill place order is successfull.

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