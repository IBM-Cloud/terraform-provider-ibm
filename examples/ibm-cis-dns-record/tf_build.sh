set -x
cd ../../
make
cp ~/work/bin/terraform-provider-ibm ~/.terraform.d/plugins/
cd $1
terraform init
terraform apply -auto-approve
