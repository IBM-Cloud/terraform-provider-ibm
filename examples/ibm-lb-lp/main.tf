resource "ibm_is_lb" "lb2" {
  name    = "myloadbalancer11"
  subnets = ["0737-e19d5734-48aa-4ca5-80b7-0c42a10886e5"]
}


resource "ibm_is_lb_listener" "lb_listener2"{
	//lb       = "r006-3e0e345b-8d93-4ad0-800b-17fd04d89817"
	lb       = ibm_is_lb.lb2.id
  	port     = "9086"
  	protocol = "http"
}
resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
	//load_balancer_id = "r006-3e0e345b-8d93-4ad0-800b-17fd04d89817"
	load_balancer_id = ibm_is_lb.lb2.id
	listener_id = ibm_is_lb_listener.lb_listener2.listener_id
	action = "redirect"
	target_http_status_code = 302
	target_url = "https://www.google.com" 
	priority = 9
	name = "mylis113" 
	//target_id = "r006-beafdff0-4fe0-4db4-8f0c-b0b4ad828712" 
//	target_id = "r006-d7a6d59b-041f-49ca-8a43-3664688bba01" 
	rules{
		condition = "contains"
		type = "header"
		field = "1"
		value = "2"
	}		  
}

