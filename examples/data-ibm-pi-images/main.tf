data "ibm_pi_images" "imagedata"
{
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "state"
{
value="${data.ibm_pi_images.imagedata.image_info}"
}
