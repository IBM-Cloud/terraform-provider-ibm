variable "space" {
  default = "space"
}

variable "org" {
  default = "org"
}

variable "service" {
  default = "cloudantNoSQLDB"
}

variable "plan" {
  default = "standard"
}

variable "service_instance_name" {
  default = "mycloudantdb183"
}

variable "service_key_name" {
  default = "mycloudantdbkey"
}

variable "app_version" {
  default = "1"
}

variable "git_repo" {
  default = "https://github.com/hkantare/cf-cloudant-python.git"
}

variable "dir_to_clone" {
  default = "/tmp/my_cf_code"
}

variable "app_zip" {
  default = "/tmp/myzip.zip"
}

variable "route" {
  default = "my-app-cloudant182"
}

variable "app_name" {
  default = "myapp178"
}

variable "app_command" {
  default = "python app.py"
}

variable "buildpack" {
  default = ""
}

