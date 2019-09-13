variable "volumesize"
{
  description="Size of the volume to be created"
}

variable "volumename"{
  description="Name of the volume to be created"
}

variable "volumetype"{
  description="Type of the volume - ssd/shared"
}

variable "volumeshareable"{
  description="Is the volume to be shared or not"
}


variable "powerinstanceid"
{
description="Power Instance associated with the account"
default="d16705bd-7f1a-48c9-9e0e-1c17b71e7331"

}

