#App example

This example shows how to perform CRUD operation on cloud foundry app.

* Create _cloudantNoSQLDB_ service with _Lite_ plan.
* Create Application
* Create a Route
* Bind the router to the application
* Upload app bits 
* Start the application and wait for it to be running
* Bind the __cloudantNoSQLDB_ service instance to the application
* Update the application 

This creates a app defined in specified space. After successfull creation it perform update and deletes app.

Example: 

```
go run main.go -path <path to the application zip file> -name my_app -org my_org -space my_space --route my_route -timeout 120s

```


