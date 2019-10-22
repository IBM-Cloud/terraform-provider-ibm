provider "ibm" {}

resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "dal13"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        notes = "endurance notes"
}
resource "ibm_storage_file" "fs_performance" {
        type = "Performance"
        datacenter = "dal13"
        capacity = 20
        iops = 200
        notes = "performance notes"
}
