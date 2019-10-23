# Global Tagging Delete Tag
Method to delete a tag from the IBM Cloud global search and tagging service. API Docs can be found in staging. https://test.cloud.ibm.com/apidocs/search The specified tag will be deleted from the users account as determined by the API key used and the users IAM permissions. Tags will only be deleted after they have been detached from all resources and the attach count is 0. Tags containing spaces must be enclosed in quotes. 

parameters:
Required:  -tag - tag to be deleted. Enclosed in quotes if it contains spaces. 

Example: 
```
go run main.go -tag "group one"
```