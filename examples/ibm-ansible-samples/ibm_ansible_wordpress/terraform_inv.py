#!/usr/bin/env python

# Terraform-Ansible dynamic inventory for IBM Cloud
# Copyright (c) 2018, IBM UK
# steve_strutt@uk.ibm.com
ti_version = '0.6'
#
# 01-10-2018 - 0.5 - Added support for Cloud Load Balancer 
# 05-10-2018 - 0.6 - Added support for IBM Cloud Resource Instances and internet-svcs
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Can be used alongside static inventory files in the same directory 
#
# This inventory script expects to find Terraform tags of the form 
# group: host_group associated with each tf instance to define the 
# host group membership for Ansible. Multiple group tags are allowed per host
#   
# terraform_inv.ini file in the same directory as this script, points to the 
# location of the terraform.tfstate file to be inventoried
# [TFSTATE]
# TFSTATE_FILE = /usr/share/terraform/ibm/Demoapp2x/terraform.tfstate
# 
# Validate correct execution: 
#   With supplied test files - './terraform_inv.py -t ../tr_test_files/terraform.tfstate'
#   With ini file './terraform.py'
# Successful execution returns groups with lists of hosts and _meta/hostvars with a detailed
# host listing.
# Validate successful operation with ansible:
#   With - 'ansible-inventory -i inventory --list'

# Resources imported into Ansible
# ibm_compute_vm_instance  - tagged in Terraform with host group
# ibm_lbaas - tagged in Terraform with host group 'cloudloadbalancers'


import json
import configparser
import os
from os import getenv
from collections import defaultdict
from argparse import ArgumentParser


def parse_params():
    parser = ArgumentParser('IBM Cloud Terraform inventory')
    parser.add_argument('--list', action='store_true', default=True, help='List Terraform hosts')
    parser.add_argument('--tfstate', '-t', action='store', dest='tfstate', help='Terraform state file in current or specified directory (terraform.tfstate default)')
    parser.add_argument('--version', '-v', action='store_true', help='Show version')
    args = parser.parse_args()
    # read location of terrafrom state file from ini if it exists 
    if not args.tfstate:
        dirpath = os.getcwd()
        print ()
        config = configparser.ConfigParser()
        ini_file = 'terraform_inv.ini'
        try:
            # attempt to open ini file first. Only proceed if found
            # assume execution from the ansible playbook directory
            filepath = dirpath + "/inventory/" + ini_file
            open(filepath) 
            
        except FileNotFoundError:
            try:
                # If file is not found it may be because command is executed
                # in inventory directory
                filepath = dirpath + "/" + ini_file
                open(filepath) 
            
            except FileNotFoundError:
                raise Exception("Unable to find or open specified ini file")
            else:
                config.read(filepath)
        else: 
            config.read(filepath)

        config.read(filepath)
        tf_file = config['TFSTATE']['TFSTATE_FILE'] 
        tf_file = os.path.expanduser(tf_file)
        args.tfstate = tf_file
    return args


def get_tfstate(filename):
    return json.load(open(filename))

def parse_state(tf_source, prefix, sep='.'):
    for key, value in list(tf_source.items()):
        try:
            curprefix, rest = key.split(sep, 1)
        except ValueError:
            continue
        if curprefix != prefix or rest == '#':
            continue

        yield rest, value


def parse_attributes(tf_source, prefix, sep='.'):
    attributes = defaultdict(dict)
    for key, value in parse_state(tf_source, prefix, sep):
        index, key = key.split(sep, 1)
        attributes[index][key] = value

    return list(attributes.values())


def parse_dict(tf_source, prefix, sep='.'):
    return dict(parse_state(tf_source, prefix, sep))

def parse_list(tf_source, prefix, sep='.'):
    return [value for _, value in parse_state(tf_source, prefix, sep)]

class TerraformInventory:
    def __init__(self):
        self.args = parse_params()
        if self.args.version:
            print(ti_version)
        elif self.args.list:
            print(self.list_all())

    def list_all(self):
        #tf_hosts = []
        hosts_vars = {}
        attributes = {}
        groups = {}
        groups_json = {}
        inv_output = {}
        group_hosts = defaultdict(list)
        for name, attributes, groups in self.get_tf_instances():
            #tf_hosts.append(name)
            hosts_vars[name] = attributes
            for group in list(groups):
                #print(group)
                group_hosts[group].append(name)
                #print(group_hosts.items())

        for group in group_hosts:
            inv_output[group] = {'hosts': group_hosts[group]}
        inv_output["_meta"] = {'hostvars': hosts_vars} 
        return json.dumps(inv_output, indent=2)    
        #return json.dumps({'all': {'hosts': hosts}, '_meta': {'hostvars': hosts_vars}}, indent=2)

    def get_tf_instances(self):
        tfstate = get_tfstate(self.args.tfstate)
        for module in tfstate['modules']:
            for resource in module['resources'].values():

                if resource['type'] == 'ibm_compute_vm_instance':
                    tf_attrib = resource['primary']['attributes']
                    name = tf_attrib['hostname']
                    group = []

                    attributes = {
                        'id': tf_attrib['id'],
                        'image': tf_attrib['os_reference_code'],
                        'ipv4_address': tf_attrib['ipv4_address'],
                        #'metadata': json.loads(tf_attrib.get('user_metadata', '{}')),
                        'metadata': tf_attrib['user_metadata'],
                        'region': tf_attrib['datacenter'],
                        'ram': tf_attrib['memory'],
                        'cpu': tf_attrib['cores'],
                        #'ssh_keys': parse_list(tf_attrib, 'ssh_key_ids'),
                        'public_ipv4': tf_attrib['ipv4_address'],
                        'private_ipv4': tf_attrib['ipv4_address_private'],
                        'ansible_host': tf_attrib['ipv4_address_private'],
                        'ansible_ssh_user': 'root',
                        'provider': 'ibm',
                        'tags': parse_list(tf_attrib, 'tags'),
                    }
                    
                        
                    #print (attributes["tags"])
                    #tag of form group: xxxxxxx is used to define ansible host group
                    for value in list(attributes["tags"]):
                        try:
                            curprefix, rest = value.split(":", 1)
                        except ValueError:
                            continue
                        if curprefix != "group" :
                            continue
                        group.append(rest)

                    yield name, attributes, group


                if resource['type'] == 'ibm_lbaas':
                    #provider = 'ibm'
                    tf_attrib = resource['primary']['attributes']
                    name = tf_attrib['name']
                    group = []

                    attributes = {
                        'id': tf_attrib['id'],
                        'vip': tf_attrib['vip'],                       
                        'region': tf_attrib['datacenter'],                       
                        'provider': 'ibm',
                        'tags': parse_list(tf_attrib, 'tags'),
                    }

                    # cloudloadbalancer's do not support tagging. So force group
                    group.append('cloudloadbalancer')
                
                    yield name, attributes, group


                # IBM Cloud services - Resource Controller instances
                if resource['type'] == 'ibm_resource_instance':
                    #provider = 'ibm'
                    tf_attrib = resource['primary']['attributes']
                    name = tf_attrib['name']
                    group = []

                    attributes = {
                        'id': tf_attrib['id'],
                        'plan': tf_attrib['plan'],
                        'service': tf_attrib['service'],
                        'status': tf_attrib['status'],
                        'resource_group': tf_attrib['resource_group_id'],                       
                        'region': tf_attrib['location'],                       
                        'provider': 'ibm',
                        'tags': parse_list(tf_attrib, 'tags'),
                    }

                    if tf_attrib['service'] == 'internet-svcs':
                        group.append('internet-svcs')

                    # Add resouce to IBM service resource group
                    group.append('ibmcloudresources')
                    #tag of form group: xxxxxxx is used to define ansible host group
                    for value in list(attributes["tags"]):
                        try:
                            curprefix, rest = value.split(":", 1)
                        except ValueError:
                            continue
                        if curprefix != "group" :
                            continue
                        group.append(rest)

                    yield name, attributes, group


                # Repeat this section for each additional resource type to be added
                #
                # if resource['type'] == 'ibm_lbaas':
                #     #provider = 'ibm'
                #     tf_attrib = resource['primary']['attributes']
                #     name = tf_attrib['name']
                #     group = []

                #     attributes = {
                #         'id': tf_attrib['id'],
                #         'vip': tf_attrib['vip'],                       
                #         'region': tf_attrib['datacenter'],                       
                #         'provider': 'ibm',
                #         'tags': parse_list(tf_attrib, 'tags'),
                #     }

                #     for value in list(attributes["tags"]):
                #         try:
                #             curprefix, rest = value.split(":", 1)
                #         except ValueError:
                #             continue
                #         if curprefix != "group" :
                #             continue
                #         group.append(rest)   



                else:    
                    continue        
             
                yield name, attributes, group


if __name__ == '__main__':
    TerraformInventory()