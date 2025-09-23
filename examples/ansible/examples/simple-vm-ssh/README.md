# IBM Cloud Ansible: VPC Virtual Server Instance

This example creates a Virtual Server Instance (VSI) inside of a Virtual
Private Cloud (VPC). The VSI is configured to allow incoming SSH connections
through a publicly accessible IP address and authenticated using an SSH key
pair.

## VPC Resources

The following VPC infrastructure resources will be created (Ansible modules in
parentheses):

* VPC (ibm_is_vpc)
* Subnet (ibm_is_subnet)
* VSI (ibm_is_instance)
* Floating IP Address (ibm_is_floating_ip)
* Security Group Rule (ibm_is_security_group_rule)

## Configuration Parameters

The following parameters can be set by the user:

* `name_prefix`: Prefix used to name created resources
* `vsi_image`: VSI image name ([retrieve available images])
* `vsi_profile`: VSI profile name ([retrieve available profiles])
* `ssh_public_key`: SSH Public Key
* `ipv4_cidr_block`: IPv4 CIDR Block for VPC Subnet
* `zone`: IBM Cloud zone

## Running

### Install IBM Cloud Ansible Modules

Note: Alternate install path is to use 'examples/install_modules.yml' playbook.

1. Download IBM Cloud Ansible modules from [release page]

2. Extract module archive.

    ```
    unzip ibmcloud_ansible_modules.zip
    ```

3. Add modules and module_utils to the [Ansible search path]. E.g.:

    ```
    cp build/modules/* $HOME/.ansible/plugins/modules/.
    cp build/module_utils/* $HOME/.ansible/plugins/module_utils/.

    ```

### Set API Key and Region

1. [Obtain an IBM Cloud API key].

2. Export your API key to the `IC_API_KEY` environment variable:

    ```
    export IC_API_KEY=<YOUR_API_KEY_HERE>
    ```

    Note: Modules also support the 'ibmcloud_api_key' parameter, but it is
    recommended to only use this when encrypting your API key value.

3. Export desired IBM Cloud region to the 'IC_REGION' environment variable:

    ```
    export IC_REGION=<REGION_NAME_HERE>
    ```

    Note: Modules also support the 'ibmcloud_region' parameter.

### Create

1. To create all resources and test public SSH connection to the VM, run the
   'create' playbook:

    ```
    ansible-playbook create.yml
    ```

### Destroy

1. To destroy all resources run the 'destroy' playbook:

    ```
    ansible-playbook destroy.yml
    ```

### List Available VSI Images and Profiles

1. To destroy all resources run the 'destroy' playbook:

    ```
    ansible-playbook list_vsi_images_and_profiles.yml
    ```

[retrieve available images]: #list-available-vsi-images-and-profiles
[retrieve available profiles]: #list-available-vsi-images-and-profiles
[Ansible search path]:https://docs.ansible.com/ansible/latest/dev_guide/overview_architecture.html#ansible-search-path
[Obtain an IBM Cloud API key]:https://cloud.ibm.com/docs/iam?topic=iam-userapikey
[release page]:https://github.com/Mavrickk3/terraform-provider-ibm/releases
