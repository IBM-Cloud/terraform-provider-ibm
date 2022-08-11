# Global IP

The Global ip example launches a web server, install nginx. It also configures the firewall set of rules to allow access to certain ip addresses/ports from specific internet addresses while denying traffic from other sources.
Global IP's provide IP flexibility by allowing users to shift workloads between servers (even in different datacenters)

To run, configure your IBMCLOUD provider

* Pass the public key while running terraform.

Running the example
```shell
terraform apply -var 'ssh_public_key=<public_key_value>'
```

After the apply is done then type the global IP from outputs in your browser and see the nginx welcome page
