provider "ibm" {}


resource "ibm_storage_block" "bs_endurance" {
        type = "Endurance"
        datacenter = "dal13"
        capacity = 20
        iops = 0.25
        snapshot_capacity = 10
        os_format_type = "Linux"
        notes = "endurance notes"
}
resource "ibm_storage_block" "bs_performance" {
        type = "Performance"
        datacenter = "dal13"
        capacity = 20
        iops = 100
        os_format_type = "Linux"
        notes = "performance notes"
}
