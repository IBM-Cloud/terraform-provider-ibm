resource "null_resource" "prepare_app_zip" {
  triggers = {
    app_version = "${var.app_version}"
    git_repo = "${var.git_repo}"
  }
  provisioner "local-exec" {
    command = <<EOF
        mkdir -p ${var.dir_to_clone}
        cd ${var.dir_to_clone}
        git init
        git remote add origin ${var.git_repo}
        git fetch
        git checkout -t origin/master
        zip -r ${var.app_zip} *
        EOF
  }
}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

data "ibm_app_domain_shared" "domain" {
  name = "mybluemix.net"
}

resource "ibm_app_route" "route" {
  domain_guid = "${data.ibm_app_domain_shared.domain.id}"
  space_guid  = "${data.ibm_space.space.id}"
  host        = "${var.route}"
}

resource "ibm_service_instance" "service" {
  name       = "${var.service_instance_name}"
  space_guid = "${data.ibm_space.space.id}"
  service    = "${var.service_offering}"
  plan       = "${var.plan}"
  tags       = ["my-service"]
}

resource "ibm_service_key" "key" {
  name = "%s"
  service_instance_guid = "${ibm_service_instance.service.id}"
}

resource "ibm_app" "app" {
  depends_on = ["ibm_service_key.key", "null_resource.prepare_app_zip"]
  name              = "${var.app_name}"
  space_guid        = "${data.ibm_space.space.id}"
  app_path          = "${var.app_zip}"
  wait_time_minutes = 10

  buildpack  = "${var.buildpack}"
  disk_quota = 512

  command               = "${var.command}"
  memory                = 256
  instances             = 2
  disk_quota            = 512
  route_guid            = ["${ibm_app_route.route.id}"]
  service_instance_guid = ["${ibm_service_instance.service.id}"]

  environment_json = {
    "somejson" = "somevalue"
  }

  app_version = "${var.app_version}"
}
