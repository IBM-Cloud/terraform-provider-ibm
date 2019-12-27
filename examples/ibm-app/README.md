# IBM Application example

This example shows how to deploy an application in the IBM PaaS

In the variables.tf you would find `git_repo` which is git url of a Cloud Foundry application repository.
You must provide valid values for the variables `org` and `space`.

When you perform `terraform apply` the provisioner  will download the code from the `git_repo` and zip it at
location specified by variable `app_zip`.

The example provisions a cloudant db service instance, create routes and assigns that route and service instance to the application.

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

To remove the stack wait for few minutes and test the stack by launching a browser with cluster url.

```shell
terraform destroy
```
