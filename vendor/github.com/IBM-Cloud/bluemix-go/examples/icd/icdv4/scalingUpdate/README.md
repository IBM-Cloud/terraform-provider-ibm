# Cloud Database Set Scaling Group 

Set scaling value on a specified group. Can only be performed on is_adjustable=true groups. Values set are for the group as a whole and resources are distributed amongst the group. Values must be greater than or equal to the minimum size and must be a multiple of the step size. ICD  resource instance must be created first via UI or using the `resource/service-instance` example to create an icd instance. Any of the ICD instance types are supported, Postgres, ElasticSearch, RabbitMQ, Etcd, etc. The 64 digit CRN of the icd instance should be supplied as the icdId. Environment variable IC_API_KEY must be set with API key and IC_REGION set to the ICD deployment region. 

Details of the API function implemented can be found in the IBM CLoud API docs: 
https://console.bluemix.net/apidocs/cloud-databases-api#set-scaling-values-on-a-specified-group

parameters:
Required:  --icdId - 64 digit CRN for ICD instance
Optional:  --memory - New memory size in multiple of step size, divided across all group instances
One or more of: 
Optional:  --cpu - New cpu count in multiple of step size, divided across all group instances
Optional:  --disk - New disk size in multiple of step size, divided across all group instances


Example: 
```
go run main.go --icdId crn:v1:bluemix:public:databases-for-postgresql:us-south:a%2F4ea1882a2d3401ed1e459979941966ea:9504e257-8916-4c00-890a-23189eeebdfd:: --memory 4096
```