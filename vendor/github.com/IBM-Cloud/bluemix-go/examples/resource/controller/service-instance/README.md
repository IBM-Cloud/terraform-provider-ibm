# Resource service-instance example

This example shows how to perform CRUD operation on resource service-instance

This creates a service instance of specified service offering, plan type, location in specified resource-group if provided else creates in default resource group. After successfull creation , updates and deletes service instance.

Example: 

```
go run main.go -service "cloud-object-storage" -plan "lite" -location "global" -name test
```




