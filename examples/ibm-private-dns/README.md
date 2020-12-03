# Private DNS resource example

This example shows how to private dns zone, permitted network, records and GLB monitor.

This sample configuration will create the vpc network that needs to be associated, private dns instance, zone under private dns instance which will get mapped to vpc network , p-dns record and GLB monitor.

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
## PDNS Resources

`IBM CLOUD PDNS GLB Monitor`
```hcl
resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "test-pdns-glb-monitor"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "test monitor description"
		interval=63
		retries=3
		timeout=8
		port=8080
		type="HTTP"
		expected_codes= "200"
		path="/health"
		method="GET"
		expected_body="alive"
		headers{
			name="headerName"
			value=["example","abc"]
		}	
  }
```

## PDNS Data Sources

`PDNS GLB Monitor`
```hcl
data "ibm_dns_glb_monitors" "test1" {
		instance_id = ibm_resource_instance.test-pdns-instance.guid		
}
```

  
