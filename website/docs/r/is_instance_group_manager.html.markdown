---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager"
description: |-
  Manages IBM VPC instance group manager.
---

# ibm_is_instance_group_manager
Create, update, or delete an instance group manager on VPC of an instance group. For more information, about instance group manager, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Basic Instance Group Manager
The following example creates a basic instance group manager.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

data "ibm_is_image" "ubuntu" {
  name = "ibm-ubuntu-20-04-6-minimal-amd64-6"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = file("~/.ssh/id_ed25519.pub")
}

resource "ibm_is_instance_template" "example" {
  name    = "example-template"
  image   = data.ibm_is_image.ubuntu.id
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_group" "example" {
  name              = "example-group"
  instance_template = ibm_is_instance_template.example.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.example.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    delete = "15m"
    update = "10m"
  }
}

resource "ibm_is_instance_group_manager" "example" {
  name                 = "example-ig-manager"
  aggregation_window   = 120
  instance_group       = ibm_is_instance_group.example.id
  cooldown             = 300
  manager_type         = "autoscale"
  enable_manager       = true
  max_membership_count = 2
  min_membership_count = 1
}

resource "ibm_is_instance_group_manager" "scheduled" {
  name           = "example-instance-group-manager"
  instance_group = ibm_is_instance_group.example.id
  manager_type   = "scheduled"
  enable_manager = true
}
```

### Example with Load Balancer, Autoscaling, and Scheduled Managers
An example to set up an instance group with both autoscaling and scheduled managers working together with a load balancer, addresses the need of both automatic scaling based on metrics and scheduled scaling for predictable workload patterns.

```terraform

resource "ibm_is_instance_template" "webserver" {
  name    = "${var.prefix}-webserver-template"
  image   = data.ibm_is_image.ubuntu.id
  profile = "cx2-2x4"

  primary_network_interface {
    subnet          = ibm_is_subnet.example.id
    security_groups = [ibm_is_security_group.webserver.id]
  }

  vpc  = ibm_is_vpc.example.id
  zone = ibm_is_subnet.example.zone
  keys = [ibm_is_ssh_key.example.id]

  # User data for web server setup
  user_data = base64encode(<<-EOF
    #!/bin/bash
    apt-get update
    apt-get install -y nginx
    systemctl start nginx
    systemctl enable nginx
    
    # Configure nginx for health checks
    echo "server { listen 80; location /health { return 200 'OK'; } }" > /etc/nginx/sites-available/health
    ln -s /etc/nginx/sites-available/health /etc/nginx/sites-enabled/
    systemctl reload nginx
  EOF
  )
}

# Load Balancer for the instance group
resource "ibm_is_lb" "webapp_lb" {
  name           = "${var.prefix}-load-balancer"
  subnets        = [ibm_is_subnet.example.id]
  type           = "private"
  
  # Enable logging for monitoring (optional)
  logging = true
}

# Backend pool for the load balancer
resource "ibm_is_lb_pool" "webapp_backend_pool" {
  name           = "${var.prefix}-backend-pool"
  lb             = ibm_is_lb.webapp_lb.id
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
  health_monitor_url     = "/health"
}

# Load balancer listener
resource "ibm_is_lb_listener" "webapp_listener" {
  lb           = ibm_is_lb.webapp_lb.id
  port         = 80
  protocol     = "http"
  default_pool = ibm_is_lb_pool.webapp_backend_pool.id
}

# Security group for web servers
resource "ibm_is_security_group" "webserver" {
  name           = "${var.prefix}-webserver-sg"
  vpc            = ibm_is_vpc.example.id
}

resource "ibm_is_security_group_rule" "webserver_http" {
  group     = ibm_is_security_group.webserver.id
  direction = "inbound"
  remote    = "0.0.0.0/0"
  tcp {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_security_group_rule" "webserver_ssh" {
  group     = ibm_is_security_group.webserver.id
  direction = "inbound"
  remote    = "0.0.0.0/0"
  tcp {
    port_min = 22
    port_max = 22
  }
}

# Instance Group with Load Balancer integration
resource "ibm_is_instance_group" "webapp_ig" {
  name              = "${var.prefix}-instance-group"
  instance_template = ibm_is_instance_template.webserver.id
  instance_count    = 3
  subnets           = [ibm_is_subnet.example.id]
  # Load balancer integration
  application_port     = 80
  load_balancer        = ibm_is_lb.webapp_lb.id
  load_balancer_pool   = element(split("/", ibm_is_lb_pool.webapp_backend_pool.id), 1)
  
  # Lifecycle management to prevent conflicts with auto-scaling
  lifecycle {
    ignore_changes = [
      instance_count,
    ]
  }
  
  timeouts {
    create = "15m"
    delete = "15m"
    update = "10m"
  }
}

# Autoscaling Manager - Handles automatic scaling based on metrics
resource "ibm_is_instance_group_manager" "webapp_autoscaler" {
  name               = "${var.prefix}-autoscale-manager"
  aggregation_window = 120
  instance_group     = ibm_is_instance_group.webapp_ig.id
  
  cooldown             = 120
  manager_type         = "autoscale"
  enable_manager       = true
  max_membership_count = 10
  min_membership_count = 2
}

# CPU-based scaling policy for the autoscaler
resource "ibm_is_instance_group_manager_policy" "webapp_cpu_policy" {
  name                   = "${var.prefix}-cpu-scaling-policy"
  instance_group         = ibm_is_instance_group.webapp_ig.id
  instance_group_manager = ibm_is_instance_group_manager.webapp_autoscaler.manager_id
  metric_type            = "cpu"
  metric_value           = 70
  policy_type            = "target"
}

# Memory-based scaling policy (optional additional policy)
resource "ibm_is_instance_group_manager_policy" "webapp_memory_policy" {
  name                   = "${var.prefix}-memory-scaling-policy"
  instance_group         = ibm_is_instance_group.webapp_ig.id
  instance_group_manager = ibm_is_instance_group_manager.webapp_autoscaler.manager_id
  metric_type            = "memory"
  metric_value           = 80
  policy_type            = "target"
}

# Scheduled Manager - Handles time-based scaling actions
resource "ibm_is_instance_group_manager" "webapp_scheduler" {
  name           = "${var.prefix}-scheduled-manager"
  instance_group = ibm_is_instance_group.webapp_ig.id
  manager_type   = "scheduled"
  enable_manager = true
}

# Scheduled action to scale down during off-peak hours (5 PM daily)
resource "ibm_is_instance_group_manager_action" "webapp_scale_down_evening" {
  name                   = "${var.prefix}-scale-down-evening"
  instance_group         = ibm_is_instance_group.webapp_ig.id
  instance_group_manager = ibm_is_instance_group_manager.webapp_scheduler.manager_id
  
  # target_manager specifies which manager's limits to modify
  target_manager = ibm_is_instance_group_manager.webapp_autoscaler.manager_id
  
  # Scale down at 5:05 PM daily (17:05 UTC)
  cron_spec = "05 17 * * *"
  
  # Adjust the autoscaler's limits for off-peak hours
  min_membership_count = 2
  max_membership_count = 4
}

# Scheduled action to scale up for peak hours (8 AM daily)
resource "ibm_is_instance_group_manager_action" "webapp_scale_up_morning" {
  name                   = "${var.prefix}-scale-up-morning"
  instance_group         = ibm_is_instance_group.webapp_ig.id
  instance_group_manager = ibm_is_instance_group_manager.webapp_scheduler.manager_id
  
  # target_manager specifies which manager's limits to modify
  target_manager = ibm_is_instance_group_manager.webapp_autoscaler.manager_id
  
  # Scale up at 8:05 AM daily (08:05 UTC)
  cron_spec = "05 08 * * *"
  
  # Adjust the autoscaler's limits for peak hours
  min_membership_count = 3
  max_membership_count = 10
}

# Weekend scale-down action (Saturday midnight)
resource "ibm_is_instance_group_manager_action" "webapp_weekend_scale_down" {
  name                   = "${var.prefix}-weekend-scale-down"
  instance_group         = ibm_is_instance_group.webapp_ig.id
  instance_group_manager = ibm_is_instance_group_manager.webapp_scheduler.manager_id
  
  target_manager = ibm_is_instance_group_manager.webapp_autoscaler.manager_id
  
  # Scale down at midnight on Saturday (00:05 UTC)
  cron_spec = "05 00 * * 6"
  
  min_membership_count = 1
  max_membership_count = 3
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `aggregation_window` - (Optional, Integer) The time window in seconds to aggregate metrics prior to evaluation.
- `cooldown` - (Optional, Integer) The duration of time in seconds to pause further scale actions after scaling has taken place.
- `enable_manager` - (Optional, Bool)  Enable or disable the instance group manager. Default value is **true**.
- `instance_group` - (Required, String) The instance group ID where instance group manager is created.
- `manager_type` - (Optional, String) The type of instance group manager. Default value is `autoscale`. Valid values are `autoscale` and `scheduled`.
- `max_membership_count`- (Required, Integer) The maximum number of members in a managed instance group.
- `min_membership_count` - (Optional, Integer) The minimum number of members in a managed instance group. Default value is `1`.
- `name` - (Optional, String) The name of the instance group manager.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `actions` - (String) List of actions of the instance group manager.
- `id` - (String) The ID in the combination of instance group ID and instance group manager ID.
- `policies` - (String) List of policies associated with the instance group manager.
- `manager_id` - (String) The ID of the instance group manager.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_instance_group_manager` resource by using `id`.
The `id` property can be formed from `instance group ID`, and `instance group manager ID`. For example:

```terraform
import {
  to = ibm_is_instance_group_manager.manager
  id = "<instance_group_id>/<instance_group_manager_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_instance_group_manager.manager <instance_group_id>/<instance_group_manager_id>
```
## Related Resources

- [`ibm_is_instance_group`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/is_instance_group) - Creates the instance group
- [`ibm_is_instance_group_manager_policy`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/is_instance_group_manager_policy) - Defines scaling policies for autoscale managers
- [`ibm_is_instance_group_manager_action`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/is_instance_group_manager_action) - Creates scheduled actions for scheduled managers
- [`ibm_is_lb`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/is_lb) - Load balancer for distributing traffic
- [`ibm_is_lb_pool`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/is_lb_pool) - Backend pool for load balancer