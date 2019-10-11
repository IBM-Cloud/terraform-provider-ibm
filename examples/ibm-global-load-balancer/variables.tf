variable "Domain_Name" {
         type = "string"
	 description = "Enter domain name"
         default = "MyCis.com"
} 

variable "Pool_Name" {
         type = "string"
         description = "Enter pool name"
         default = "OriginPool.com"
} 

variable "origin1" {
         type = "string"
         description = "Enter origin-1 name"
         default = "Pool-1"
} 

variable "origin2" {
         type = "string"
	description = "Enter origin-2 name"
         default = "Pool-2"
} 

variable "num_origins"  {
	description = "Enter number of origins"
         default = 1
} 

