# Resource service-instance-key example

This example shows how to perform CRUD operation on resource service-instance-key

This creates a service instance of specified service offering, plan type, location in specified resource-group if provided else creates in default resource group. After successfull creation , it assocaites a service-instance key with specified roles and deletes service-instance-key, service instance.

Example: 

```
go run main.go -service "cloud-object-storage" -plan "lite" -location "global" -name test -key "testkey" -role "Viewer"
```




