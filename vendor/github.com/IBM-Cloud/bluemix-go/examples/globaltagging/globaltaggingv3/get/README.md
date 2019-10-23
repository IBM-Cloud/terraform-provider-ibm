# Global Tagging Get Tags
Method to retrieve tags attached to a resource using the IBM Cloud global search and tagging service. API Docs can be found in staging. https://test.cloud.ibm.com/apidocs/search All tags will be returned for the identified resource in the users account as determined by the API key used and the users IAM permissions. 

parameters:
Required:  -id - 64 digit CRN for IBM Cloud resource

Example: 
```
go run main.go -id crn:v1:bluemix:public:kms:us-south:a/4ea1882a2d3401ed1e459979941966ea:254284b4-698d-4f10-b3f9-f1fdad7fb132::\" -tags webserver,dbserver 
```