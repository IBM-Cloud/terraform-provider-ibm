data "ibm_pi_images" "imagedata"
{
pi_cloud_instance_id="${var.powerinstanceid}"
pi_image_name="${var.imagename}"
}



output "images"
{
value="${data.ibm_pi_images.imagedata.image_info.0.image_id}"
}
