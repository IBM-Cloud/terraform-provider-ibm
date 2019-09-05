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
default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}

