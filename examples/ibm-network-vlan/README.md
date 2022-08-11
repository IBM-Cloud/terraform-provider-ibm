# IBM Cloud vlan type Resource Example

The example launches vlan type Resource.
To run, configure your IBM Cloud provider

## Get up and running

* Pass the public key while running terraform.

* Planning phase

```shell
terraform plan -var 'ssh_public_key=<public_key_value>'
```

* Apply phase

```shell
terraform apply -var 'ssh_public_key=<public_key_value>'
```

* Destroy

```shell
terraform destroy -var 'ssh_public_key=<public_key_value>'
```
