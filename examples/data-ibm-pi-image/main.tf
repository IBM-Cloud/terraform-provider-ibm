data "ibm_pi_image" "imagedata"
{
pi_image_name="${var.imagename}"
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "state"
{
value="${data.ibm_pi_image.imagedata.state}"
}

output "imageid"
{
value="${data.ibm_pi_image.imagedata.imageid}"
}

