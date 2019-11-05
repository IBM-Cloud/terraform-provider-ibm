data "ibm_pi_images" "imagedata"
{ 
pi_cloud_instance_id="${var.powerinstanceid}"
pi_image_name="${var.imagename}"
}



resource "ibm_pi_image" "imagedata"
{
pi_image_name="${var.imagename}"
pi_image_id="${data.ibm_pi_images.imagedata.image_info.0.image_id}"
pi_cloud_instance_id="${var.powerinstanceid}"
}


output "image_id"
{
value="${ibm_pi_image.imagedata.id}"
}

