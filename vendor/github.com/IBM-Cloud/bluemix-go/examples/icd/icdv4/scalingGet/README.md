# Cloud Database Get Scaling Groups information 

Scaling groups represent the various resources allocated to a deployment. ICD  resource instance must be created first via UI or using the `resource/service-instance` example to create an icd instance. Any of the ICD instance types are supported, Postgres, ElasticSearch, RabbitMQ, Etcd, etc. The 64 digit CRN of the icd instance should be supplied as the icdId. Environment variable IC_API_KEY must be set with API key and IC_REGION set to the ICD deployment region. 

Details of the API function implemented can be found in the IBM CLoud API docs: 
https://console.bluemix.net/apidocs/cloud-databases-api#get-currently-available-scaling-groups-from-a-depl

parameters:
Required:  --icdId - 64 digit CRN for ICD instance

Example: 
```
go run main.go --icdId crn:v1:bluemix:public:databases-for-postgresql:us-south:a%2F4ea1882a2d3401ed1e459979941966ea:9504e257-8916-4c00-890a-23189eeebdfd:: 
```