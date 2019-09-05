
variable "pvm_instance_name"
{
description="Name of the instance"
}


variable "power_instance_id"
{
description="Power Instance associated with the account"
default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}

variable "operation"
{
description="Operation to execute on the lpar"
default="hard-reboot"
}
