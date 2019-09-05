variable "servername"
{
  description="Name of the server"
}


variable "powerinstanceid"
{
description="Power Instance associated with the account"
default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}


variable "memory"{
  description="Memory Of the Power VM Instance"
 default="2"
}

variable "processors"{
  description="Processor Count on the server"
  default="1"
}

variable "proctype"{
  description="Processor Type for the LPAR - shared/dedicated"
default="shared"
}


variable "sshkeyname"
{
description="Key Name to be passed"
default="bramcolo"
}

variable "volumename"
{
description="Volume Name to be created"
}

variable "volumesize"
{
description="Volume Size to be created"
default="40"
}

variable "volumetype"
{
description="Type of volume to be created - ssd/shared"
default="ssd"
}

variable "shareable"{
description="Should the volume be shared or not true/false"
default="true"
}

variable "networks" {
  default=["POWERCFN","POWERBACKUP","POWERADMIN"]
}

variable "systemtype"{
  description = "Systemtype of the server"
  default="s922"
}

variable "migratable"{
  description = "Server can be migrated"
default="true"
}

variable "imagename"{
  description = "Name of the image"
  default="7200-03-03"
}

variable "replicationpolicy"
{
description="Replication Policy of the vm"
default="none"
}
variable "replicants"
{
description="Number of replicants"
default=1
}
