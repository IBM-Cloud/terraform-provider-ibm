# IBM Cloud Object Storage example

This example shows how to setup cloud object storage on classic/vpc-cluster including helm initialization.

In this example, a cloud object store is instantiated and attached to existing cluster.

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

The below steps needs to be executed to verify the cloud object storage attachement to the cluster.

```shell
export KUBECONFIG=<kube config file path>

cat deployment-example.yaml | sed 's/CLAIM_NAME/<name of the physical volume claimed>/g' | kubectl apply -f -

PVC_TEST_POD=`kubectl get pods --all-namespaces | grep "pvc-config" | awk '{print$2}'`

kubectl exec -it $PVC_TEST_POD sh

cd /tmp/test; echo "this is test data" > test.text

```

verfiy the test.text file in the configured cos bucket in the IBM cloud UI.

