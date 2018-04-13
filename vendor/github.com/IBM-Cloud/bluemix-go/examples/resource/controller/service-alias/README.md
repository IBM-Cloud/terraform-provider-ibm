# Resource service-instance -alias example

This example shows how to perform CRUD operation on resource service-instance-alias

This creates a service instance of specified service offering, plan type, location in specified resource-group if provided else creates in default resource group. After successfull creation , it creates a alias of the instance in specified org and space and then  deletes service instance alias and service instance.

Example: 

```
go run main.go -service "cloud-object-storage" -plan "lite" -location "global" -name test -alias testalias -org abc@in.ibm.com -space prod
```




