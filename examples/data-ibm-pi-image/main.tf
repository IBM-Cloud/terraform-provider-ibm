data "ibm_pi_image" "imagedata"
{
name="${var.imagename}"
powerinstanceid="${var.powerinstanceid}"

}



output "state"
{
value="${data.ibm_pi_image.imagedata.state}"
}

output "imageid"
{
value="${data.ibm_pi_image.imagedata.imageid}"
}

