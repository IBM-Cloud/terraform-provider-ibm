# Cloud Database Get Scaling Groups Default information 

Scaling groups represent the various resources allocated to a deployment. When a new deployment is created, there are a set of defaults for each database type. Any of the ICD types, etcd, rabbitmq, postgresql etc 

Details of the API function implemented can be found in the IBM CLoud API docs: 
https://console.bluemix.net/apidocs/cloud-databases-api#get-default-scaling-groups-for-a-new-deployment

parameters:
Required:  --icdId - 64 digit CRN for ICD instance

Example: 
```
go run main.go --groupType etcd
```