data "ibm_pi_tenant" "tenantdata"
{
#tenantid="${var.tenant}"
powerinstanceid="${var.powerinstanceid}"
}

output "tenantid"
{
value="${data.ibm_pi_tenant.tenantdata.tenantid}"
}

output "tenantinstances"
{
value="${data.ibm_pi_tenant.tenantdata.cloudinstances}"
}
#
output "enabled"
{
value="${data.ibm_pi_tenant.tenantdata.enabled}"
}

#output "createdate"
#{
#value="${data.ibm_pi_tenant.tenantdata.creationdate}"
#}

output "tenantname"
{
value="${data.ibm_pi_tenant.tenantdata.tenantname}"
}
