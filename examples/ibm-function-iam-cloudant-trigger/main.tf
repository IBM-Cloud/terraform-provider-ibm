data "ibm_resource_group" "resource-group" {
   name = var.resource_group
}

resource "ibm_function_namespace" "namespace" {
   name                = var.namespace
   resource_group_id   = data.ibm_resource_group.resource-group.id
}

resource "null_resource" "prepare_app_zip" {
  triggers = {
    app_version = var.app_version
    git_repo    = var.git_repo
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

data "ibm_space" "spacedata" {
  space = var.space
  org   = var.org
}

resource "ibm_service_instance" "service-instance" {
  name       = var.service_instance_name
  space_guid = data.ibm_space.spacedata.id
  service    = var.service
  plan       = var.plan
  tags       = ["cluster-service", "cluster-bind"]
}

resource "ibm_service_key" "serviceKey" {
  name                  = var.service_key_name
  service_instance_guid = ibm_service_instance.service-instance.id
}

data "ibm_app_domain_shared" "domain" {
  name = "mybluemix.net"
}

resource "ibm_app_route" "route" {
  domain_guid = data.ibm_app_domain_shared.domain.id
  space_guid  = data.ibm_space.spacedata.id
  host        = var.route
}

resource "ibm_app" "app" {
  depends_on = [
    ibm_service_key.serviceKey,
    null_resource.prepare_app_zip,
  ]
  name              = var.app_name
  space_guid        = data.ibm_space.spacedata.id
  app_path          = var.app_zip
  wait_time_minutes = 10

  buildpack  = var.buildpack
  
  memory                = 256
  instances             = 2
  disk_quota            = 512
  route_guid            = [ibm_app_route.route.id]
  service_instance_guid = [ibm_service_instance.service-instance.id]
  app_version           = var.app_version
  command               = var.app_command
}


resource "ibm_function_package" "package" {
  depends_on = [ibm_function_namespace.namespace,]

  name = var.packageName
  namespace = var.namespace

  user_defined_parameters = <<EOF
        [
    {
        "key":"name",
        "value":"terraform"
    },
    {
        "key":"place",
        "value":"India"
    }
]
EOF

}

resource "ibm_function_action" "action" {
  depends_on = [ibm_function_namespace.namespace,]

  name = "${ibm_function_package.package.name}/${var.actionName}"
  namespace = var.namespace

  exec {
    kind = "nodejs:10"
    code = file("hello.js")
  }
}

resource "ibm_function_package" "boundpackage" {
  depends_on = [ibm_function_namespace.namespace,]

  name              = var.boundPackageName
  bind_package_name = "/whisk.system/cloudant"
  namespace = var.namespace

  user_defined_parameters = <<EOF
	[
    {
        "key":"username",
        "value":"${ibm_service_key.serviceKey.credentials["username"]}"
    },
    {
        "key":"password",
        "value":"${ibm_service_key.serviceKey.credentials["password"]}"
    },
   {
        "key":"host",
        "value":"${ibm_service_key.serviceKey.credentials["host"]}"
    }
    ]
EOF

}

resource "ibm_function_trigger" "trigger" {
  depends_on = [ibm_app.app,
		ibm_function_namespace.namespace,]

  name       = var.triggerName
  namespace = var.namespace

  feed {
    name = "${ibm_function_package.boundpackage.name}/changes"

    parameters = <<EOF
	[
   	{
        "key":"dbname",
        "value":"${var.dbname}"
    },
    {
        "key":"includeDocs",
        "value":"true"
    }
]
EOF

  }
}

resource "ibm_function_rule" "rule" {
  depends_on = [ibm_function_namespace.namespace,]

  name         = var.ruleName
  namespace = var.namespace
  trigger_name = ibm_function_trigger.trigger.name
  action_name  = ibm_function_action.action.name
}

