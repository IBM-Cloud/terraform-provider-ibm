# Service Instance example

This example shows how to perform CRUD operation on cloud foundry service instance and how to associates a service key.

This creates a service instance of specified service offering and plan type. After successfull creation it associates a service key, deletes service key and service instance.

Example: 

```
go run main.go -org example.com -space test
```
 
If user doesn't want to delete service instance and service key

Example:

```
go run main.go -org example.com -space test -no-delete true
```



