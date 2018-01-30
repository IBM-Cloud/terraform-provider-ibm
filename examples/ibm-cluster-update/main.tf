provider "ibm" {}

data "ibm_org" "org" {
  org = "${var.org}"
}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

data "ibm_account" "account" {
  org_guid = "${data.ibm_org.org.id}"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.name}"
  org_guid     = "${data.ibm_org.org.id}"
  space_guid   = "${data.ibm_space.space.id}"
  account_guid = "${data.ibm_account.account.id}"
}

resource "null_resource" "drain_cluster" {
  provisioner "local-exec" {
    command = <<EOF
        export KUBECONFIG="${data.ibm_container_cluster_config.cluster_config.config_file_path}"
        for i in `kubectl get nodes -o name`
        do
            node=`echo $i | awk -F'/' '{print $2}'`
            echo "drain node  $node"
            kubectl drain $node --force --timeout 60s --ignore-daemonsets --delete-local-data
        done
        EOF
  }
}

resource "ibm_container_cluster" "cluster" {
  name         = "${var.cluster_name}${random_id.name.hex}"
  datacenter   = "${var.datacenter}"
  org_guid     = "${data.ibm_org.org.id}"
  space_guid   = "${data.ibm_space.space.id}"
  account_guid = "${data.ibm_account.account.id}"
  no_subnet    = true
  subnet_id    = ["${var.subnet_id}"]
  kube_version = "1.8.6"

  workers = [{
    name   = "worker1"
    version = "1.8.6"
  },
    {
      name   = "worker2"
      version = "1.8.6"
    },
    {
      name   = "worker3"
      version = "1.8.6"
    },
  ]

  machine_type    = "${var.machine_type}"
  isolation       = "${var.isolation}"
  public_vlan_id  = "${var.public_vlan_id}"
  private_vlan_id = "${var.private_vlan_id}"
}

resource "null_resource" "uncordon_cluster" {
  provisioner "local-exec" {
    command = <<EOF
        export KUBECONFIG="${data.ibm_container_cluster_config.cluster_config.config_file_path}"
        for i in `kubectl get nodes -o name`
        do
            node=`echo $i | awk -F'/' '{print $2}'`
            echo "drain node  $node"
            kubectl uncordon $node
        done
        EOF
  }
  depends_on = ["ibm_container_cluster.cluster"]
}
