resource "ibm_is_lb" "lb2"{
  name    = "mylb"
  subnets = ["35860fed-c911-4936-8c94-f0d8577dbe5b"]
}

resource "ibm_is_lb_listener" "lb_listener2"{
        lb       = ibm_is_lb.lb2.id
        port     = "9086"
        protocol = "http"
}
resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
        load_balancer_id = ibm_is_lb.lb2.id
        listener_id = ibm_is_lb_listener.lb_listener2.listener_id
        action = "redirect"
        priority = 2
        name = "mylistener8"
	target_http_status_code = 302
	target_url = "https://www.google.com"
        //target_id = "r006-beafdff0-4fe0-4db4-8f0c-b0b4ad828712"
	rules{
                condition = "contains"
                type = "header"
                field = "1"
                value = "2"
        }
}
