# Zones Create/Get example

This example shows how to create and get a zone. CIS resource instance must be created first via UI or using the `resource/service-instance` example to create an `internet-svcs` instance. 64 digit CRN should be supplied as cis_id variable for zone create/delete. Environment variable BM_API_KEY must be set with API key. 

Paired with Get/Delete example. The Zone created in this example can be used as input to the DNS/Pool/Monitor/GLB examples. The zone_id for other examples can be retrived from the id field in the last line of the command output.
{"result": {"id": `"b6e1169a4a9fef8d8ff984fee4a4eb20"`, "name": "example.com"

Example: 

```
go run main.go -cis_id="crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:d268d835-3ef5-4049-8526-296ff08020a0::" -domain="example.com"
```