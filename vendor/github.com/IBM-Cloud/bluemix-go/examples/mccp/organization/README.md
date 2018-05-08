# Organization example

This example shows how to perform CRUD operation on cloud foundry organization.

This creates an organization with name specified. To create an organization a user must have authorization. 
Creation of organization is commented. A user must uncomment it to create an organization.

After successful creation it performs to find the organization and then update it with new name specified.

Finally this example deletes the newly created organization. To delete an organization a user must have authorization.
Deletion of organization is commented. A user must uncomment it to delete an organization.

Example:

```go run main.go -org "example.com" -neworg "newexample.com"```
