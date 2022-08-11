variable "namespace" {
  default = ""
}

variable "resource_group" {
 default = "Default"
}

variable "packageName" {
  default = "utils"
}

variable "actionName" {
  default = "hello"
}

variable "boundPackageName" {
  default = "mycloudant"
}

variable "triggerName" {
  default = "myCloudantTrigger"
}

variable "ruleName" {
  default = "cloudantRule"
}

variable "dbname" {
  default = "databasedemo"
}

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
  default = "Lite"
}

variable "service_instance_name" {
  default = "mycloudantdb"
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
  default = "my-app-cloudant"
}

variable "app_name" {
  default = "myapp"
}

variable "app_command" {
  default = "python app.py"
}

variable "buildpack" {
  default = ""
}

