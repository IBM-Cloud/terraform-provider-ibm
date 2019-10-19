data "ibm_pi_images" "imagedata"
{
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "images"
{
value="${data.ibm_pi_images.imagedata.image_info}"
}
