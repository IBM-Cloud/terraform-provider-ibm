# IBM Function example

This example shows how IBM Cloud Functions action is triggered when documents in Cloudant NoSQL databases are changed or added.

In this example a Cloudant NoSQL service instance is created. We deploy a python app which creates a database 'databasedemo' in Cloudant  NOSQL. We bind a cloudant package using IBM Cloud Function package and create an action, trigger and rule.

When you change documents or add documents in your Cloudant database you can see logs in IBM Cloud Functions dashboard.


To run, configure your IBM Cloud provider

Running the example

For planning phase

```shell
terraform plan
```

For apply phase

```shell
terraform apply
```

For destroy

```shell
terraform destroy
```
