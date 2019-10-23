variable "network_type"
{
  description="Type of network to create -vlan/pub-vlan"
}

variable "network_name"{
  description="Name of the network to be created"
}

variable "network_dns"{
  description="Value of the DNS Servers"
}

variable "network_CIDR"{
  description="Value of the Network CDIR"
}



variable "cloudinstanceid"
{
description="Power Instance associated with the account"
#default="d16705bd-7f1a-48c9-9e0e-1c17b71e7331"
#default="d7d4b40a-b06f-40a9-aec0-7ccd144cfa7b"
default="7c548b9d-504b-4a50-98a2-5d914c0bacef"
}

