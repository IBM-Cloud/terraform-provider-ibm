resource "null_resource" "customResourceGroup" {

    provisioner "local-exec" {

        when = create
        command = "sh ./createcrd.sh"

    }

    provisioner "local-exec" {

       when = destroy
       command = "sh ./deletecrd.sh"

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
        command = "sh ./deleteaddon.sh"

    }

}


resource "null_resource" "updateCRD" {

    triggers = {
        numOfOsd = var.numOfOsd
        ocsUpgrade = var.ocsUpgrade
        workerNodes = var.workerNodes
        workerPools = var.workerPools
        osdDevicePaths = var.osdDevicePaths
        taintNodes = var.taintNodes
        addSingleReplicaPool = var.addSingleReplicaPool
        resourceProfile = var.resourceProfile
        enableNFS = var.enableNFS
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
