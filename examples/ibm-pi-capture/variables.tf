variable "instancename"
{
  description="Instance Name to exported"
}

variable "capturename"{
  description="Capture Name for the instance. Make sure it's unique"
}

variable "capturedestination"{
  description="Destination of the capture image-catalog/cloud-storage/both"
}


variable "powerinstanceid"
{
description="Power Instance associated with the account"
#default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
default="d16705bd-7f1a-48c9-9e0e-1c17b71e7331"
}

