# IBM Power Virtual Server in IBM Cloud

This example creates a Power Systems Virtual Server running AIX or IBM i. The
server is configured to allow incoming SSH connections through a publicly
accessible IP address and authenticated using the provided SSH key.

## Power Systems Virtual Server Resources

The following infrastructure resources will be created (Ansible modules in
parentheses):

* SSH Key (ibm_pi_key)
* Virtual Server Instance (ibm_pi_instance)

## Configuration Parameters

The following parameters can be set by the user:

* `pi_name`: Name assigned to Virtual Server Instance
* `sys_type`: The type of system on which to create the VM (s922/e880/any)
* `pi_image`: VM image name ([retrieve available images])
* `proc_type`: The type of processor mode in which the VM will run
               (shared/dedicated)
* `processors`: The number of vCPUs to assign to the VM (as visibile within the
                guest operating system)
* `memory`: The amount of memory (GB) to assign to the VM
* `pi_cloud_instance_id`: The cloud_instance_id for this account
* `ssh_public_key`: The value of the ssh public key to be authorized for SSH
                    access

## Running

### Install IBM Cloud Ansible Modules

Note: Alternate install path is to use 'examples/install_modules.yml' playbook.

1. Download IBM Cloud Ansible modules from [release page].

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

### List Available PI VM Images

1. To list available images run the 'list_pi_images' playbook. *note: Images
   are specific to a PI instance, and thus the 'pi_cloud_instance_id' var
   must be set before running this playbook.:

    ```
    ansible-playbook list_pi_images.yml
    ```

[retrieve available images]: #list-available-pi-images
[release page]:https://github.com/Mavrickk3/terraform-provider-ibm/releases
