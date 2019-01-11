# Enable VLAN spanning for the IaaS account. 
# Required to enable connectivity and routing between data centers for Security Groups

resource "ibm_network_vlan_spanning" "spanning" {
  "vlan_spanning" = "on"
}
