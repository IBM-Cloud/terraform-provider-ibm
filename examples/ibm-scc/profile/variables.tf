variable "ibm_scc_control_library_id" {
	description = "The id of the control library"
	type        = string
	default     = ""
}

variable "ibm_scc_profile_name" {
	description = "The name of the profile"
	type        = string
	default     = "scc_demo_profile"
}

variable "ibm_scc_profile_description" {
	description = "The description of the profile"
	type        = string
	default     = "This profile as a demo using Terraform"
}

variable "ibm_scc_profile_attachment_name" {
	description = "The name of the profile"
	type        = string
	default     = "scc_demo_profile_attachment"
}

variable "ibm_scc_profile_attachment_desc" {
	description = "The description of the profile attachment"
	type        = string
	default     = "This description of the profile attachnment made by Terraform"
}
