variable "hostip"{
  default = ["169.60.225.4","169.60.225.2","169.60.225.5","169.60.225.8","169.60.225.14","169.60.225.10","169.60.225.11"]
}

variable "passwords"{
  default = ["bTjkrct5","xTKZ32Ln","mfaFTp7f"]
}

variable "path"{
  default = "/var/folders/k9/s119gk81351fj53h3fm0pvgc0000gn/T/register-host_brp4rbq2021qlc4nvc0g_517939670.sh"
}

variable "ipcount"{
  default = "7"
}

variable "private_ssh_key"{
  default = "id_rsa"
}
variable "module_depends_on" {
}