# Global Tagging Detach Tag
Method to detach one or more tags from a resource using the IBM Cloud global search and tagging service. API Docs can be found in staging. https://test.cloud.ibm.com/apidocs/search The specified tags will be detached from the identified resource in the users account as determined by the API key used and the users IAM permissions. The detached tags are not deleted, as they may still remain attached to other resource. Tags can be deleted using the Delete Method after they have been detached from all resources and the attach count is 0. 

parameters:
Required:  -tags - comma separated list of tags, without spaces
Required:  -id - 64 digit CRN for IBM Cloud resource

Example: 
```
go run main.go -id crn:v1:bluemix:public:kms:us-south:a/4ea1882a2d3401ed1e459979941966ea:254284b4-698d-4f10-b3f9-f1fdad7fb132::\" -tags webserver,dbserver 
```