module "satellite-storage-assignment"{
  source = "./modules/assignment"

  assignment_name = var.assignment_name
  cluster = var.cluster
  config = var.config
  controller = var.controller
}