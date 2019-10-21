data "ibm_power_image" "powerimages"{
  name="${var.image_name}"
}

output "state"
{
value="${data.ibm_power_image.powerimages.state}"
}
