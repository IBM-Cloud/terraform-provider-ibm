## Version 1.0



data "ibm_pi_instance" "instance"
{ 
instancename="${var.instancename}"
pi_cloud_instance_id="${var.powerinstanceid}"

}


resource "ibm_pi_capture" "picapture"{

  pi_instance_name = "${data.ibm_pi_instance.instance.id}"
  pi_capture_name = "${var.capturename}"
  pi_capture_destination = "${var.capturedestination}"
  pi_cloud_instance_id="${var.powerinstanceid}"

}
