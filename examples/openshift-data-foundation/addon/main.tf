resource "null_resource" "customResourceGroup" {

    provisioner "local-exec" {
       
        when = create
        command = "sh ./createcrd.sh"
      
    }

    provisioner "local-exec" {
      
       when = destroy
       command = "cd ocscluster && terraform destroy --auto-approve -var-file input.tfvars"

    }

    depends_on = [
      null_resource.addOn
    ]

   
}


resource "null_resource" "addOn" {

    provisioner "local-exec" {

        when = create
        command = "sh ./createaddon.sh"

    }

    provisioner "local-exec" {

        when = destroy
        command = "cd ibm_odf_addon && terraform destroy --auto-approve -var-file input.tfvars"
      
    }
 
}


resource "null_resource" "updateCRD" {

    triggers = {
        numOfOsd = var.numOfOsd
        ocsUpgrade = var.ocsUpgrade
        workerNodes = var.workerNodes
    }
    

    provisioner "local-exec" {
       
        command = "sh ./updatecrd.sh"
      
    }

    depends_on = [
      null_resource.upgradeODF
    ]

}

resource "null_resource" "upgradeODF" {

    triggers = {
      odfVersion = var.odfVersion
    }

    provisioner "local-exec" {

        command = "sh ./updateodf.sh"
      
    }
  
   depends_on = [
      null_resource.customResourceGroup, null_resource.addOn
    ]

}