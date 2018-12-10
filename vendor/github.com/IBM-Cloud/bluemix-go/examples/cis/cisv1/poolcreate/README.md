# Pool Create/Get example

This example shows how to create and get a pool. CIS resource instance must be created first via UI or using the `resource/service-instance` example to create an `internet-svcs` instance. 64 digit CRN should be supplied as cis_id variable for zone create/delete. Environment variable BM_API_KEY must be set with API key. 

Paired with Get/Delete example. The Pool created in this example can be used as input to the DNS/Monitor/GLB examples. The pool_id can be retrived from the id field in the last line of the command output.
{"result": {"description": "", "created_on": "2018-12-10T09:35:47.269909Z", "modified_on": "2018-12-10T09:35:47.269909Z", "id": `"00f0c9cad99646eed248e1f126a1c1ac"`,

Example: 

```
go run main.go -cis_id="crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:d268d835-3ef5-4049-8526-296ff08020a0::" 
```