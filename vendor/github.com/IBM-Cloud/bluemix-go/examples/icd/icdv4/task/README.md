# Cloud Database Task status information 

Get information about a task and its status. Tasks themselves are persistent so old tasks may be consulted as well as active/running tasks. ICD  resource instance must be created first via UI or using the `resource/service-instance` example to create an icd instance. Any of the ICD instance types are supported, Postgres, ElasticSearch, RabbitMQ, Etcd, etc. The 64 digit CRN of the icd instance should be supplied as the icdId. Environment variable IC_API_KEY must be set with API key and IC_REGION set to the ICD deployment region. 

Most ICD operations are performed as tasks and return task object that includes the task ID. See scalingupdate for an example of usage of this command to query running tasks. 

Details of the API function implemented can be found in the IBM CLoud API docs: 
https://console.bluemix.net/apidocs/cloud-databases-api#get-information-about-a-task

parameters:
Required:  --taskId - Task ID returned as task.Id from calls to scalingupdate etc. 

Example: 
```
go run main.go --taskId crn:v1:bluemix:public:databases-for-postgresql:us-south:a%2F4ea1882a2d3401ed1e459979941966ea:9504e257-8916-4c00-890a-23189eeebdfd:task:efd741ac-c32a-4b5c-94d5-2471e7534a4e
```