# Download Cluster Config example

This example shows how to download the kubernetes cluster configuration.

This downloads the configuration of given cluster at the destination path specified.

Example: 

```
go run main.go -org example.com -space test -clustername myCluster -path /home/myCluster/config
```