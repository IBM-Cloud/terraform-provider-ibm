# Global Search Query
Query to return attributes from global search and tagging service. API Docs can be found in staging. https://test.cloud.ibm.com/apidocs/search 

parameters:
Required:  --query

Example: 
```
go run main.go -query "crn:\"crn:v1:bluemix:public:kms:us-south:a/4ea1882a2d3401ed1e459979941966ea:254284b4-698d-4f10-b3f9-f1fdad7fb132::\"" 
```