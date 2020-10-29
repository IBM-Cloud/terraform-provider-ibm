import pexpect
import re
import time
import os


def list_crypto_units():
    print("########### ibmcloud tke cryptounits ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounits', encoding='utf-8')
    out= child.readlines()
    result = "".join(out)
    print (result)
    return result

def crypto_unit_add(cu_num):
    print("########### ibmcloud tke cryptounit-add ########### \n")
    if cu_num == "":
        print ("[ERROR] Selected crypto unit number or hpcs id is invalid")
        return
    child = pexpect.spawn('ibmcloud tke cryptounit-add', encoding='utf-8')
    child.expect ('Enter a list of CRYPTO UNIT NUM to add, separated by spaces:')
    child.sendline (cu_num)
    out= child.readlines()
    selected_crypto = "".join(out)
    print (selected_crypto)

def list_sig_keys():
    print("########### ibmcloud tke sigkeys ########### \n")
    child = pexpect.spawn('ibmcloud tke sigkeys', encoding='utf-8')
    out= child.readlines()
    sig_key_list = "".join(out)
    print (sig_key_list)
    return sig_key_list

def sigkey_add(admin_name,admin_password):
    print("########### ibmcloud tke sigkey-add ########### \n")
    child = pexpect.spawn('ibmcloud tke sigkey-add', encoding='utf-8')
    child.expect ('Enter an administrator name to be associated with the signature key:')
    child.sendline (admin_name)
    child.expect ('Enter a password to protect the signature key:')
    child.sendline (admin_password)
    child.expect ('Re-enter the password to confirm:')
    child.sendline (admin_password)
    out= child.readlines()
    added_sigkey = "".join(out)
    print (added_sigkey)
    return added_sigkey

def sigkey_select(keynum,admin_password):
    print("########### ibmcloud tke sigkey-sel ########### \n")
    child = pexpect.spawn('ibmcloud tke sigkey-sel',encoding='utf-8')
    child.expect ('Enter the KEYNUM values to select as current signature keys, separated by spaces:')
    child.sendline (keynum)
    child.expect ('Enter the password for KEYNUM '+keynum+':')
    child.sendline (admin_password)
    out= child.readlines()
    selected_sigkey = "".join(out)
    print (selected_sigkey)
    return selected_sigkey

def admin_add(key_num,admin_password):
    print("########### ibmcloud tke cryptounit-admin-add ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-admin-add',encoding='utf-8')
    child.timeout=300
    child.expect ('Enter the KEYNUM of the administrator signature key you wish to load:')
    child.sendline (key_num)
    child.timeout=300
    child.expect ('Enter the password for the administrator signature key file:')
    child.sendline (admin_password)
    child.timeout=300
    out= child.readlines()
    added_admin = "".join(out)
    print (added_admin)
    return added_admin

def list_admins():
    print("########### ibmcloud tke cryptounit-admins ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-admins', encoding='utf-8')
    out= child.readlines()
    admin_list = "".join(out)
    print (admin_list)

def threshold_set(threshold_value,rev_threshold_value,admin_password):
    print("########### ibmcloud tke cryptounit-thrhld-set ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-thrhld-set',encoding='utf-8')
    child.expect ('Enter the new signature threshold value:')
    child.sendline (threshold_value)
    child.timeout=300
    child.expect ('Enter the new revocation signature threshold value:')
    child.sendline (rev_threshold_value)
    child.timeout=300
    child.expect ('Enter the password for the signature key identified by:')
    child.sendline (admin_password)
    child.timeout=300
    out= child.readlines()
    thres_set = "".join(out)
    print (thres_set)
    return thres_set

def list_thresholds():
    print("########### ibmcloud tke cryptounit-thrhlds ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-thrhlds', encoding='utf-8')
    out= child.readlines()
    thresholds = "".join(out)
    print (thresholds)

def mk_random_add(description,password):
    print("########### ibmcloud tke mk-add --random ########### \n")
    child = pexpect.spawn('ibmcloud tke mk-add --random', encoding='utf-8')
    child.timeout=300
    child.expect ('Enter a description for the key part:')
    child.sendline (description)
    child.timeout=300
    child.expect ('Enter a password to protect the key part:')
    child.sendline (password)
    child.timeout=300
    child.expect ('Re-enter the password to confirm:')
    child.sendline (password)
    child.timeout=300
    out= child.readlines()
    mk_key = "".join(out)
    print (mk_key)
    return mk_key

def mk_custom_add(description,password,key):
    print("########### ibmcloud tke mk-add --value ########### \n")
    child = pexpect.spawn('ibmcloud tke mk-add --value', encoding='utf-8')
    child.timeout=300
    child.expect ('Enter the key part value:')
    child.sendline (key)
    child.timeout=300
    child.expect ('Enter a description for the key part:')
    child.sendline (description)
    child.timeout=300
    child.expect ('Enter a password to protect the key part:')
    child.sendline (password)
    child.timeout=300
    child.expect ('Re-enter the password to confirm:')
    child.sendline (password)
    child.timeout=300
    out= child.readlines()
    mk_key = "".join(out)
    print (mk_key)
    return mk_key

def list_mks():
    child = pexpect.spawn('ibmcloud tke mks', encoding='utf-8')
    out= child.readlines()
    mks = "".join(out)
    print (mks)
    return mks

def mk_load(mk_keynum,admin_password,key1_password,key2_password,key3_password):
    print("########### ibmcloud tke cryptounit-mk-load ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-mk-load', encoding='utf-8')
    child.expect ('Enter the KEYNUM values of the master key parts you wish to load.')
    child.sendline (mk_keynum)
    child.expect ('Enter the password for the signature key identified by:')
    child.sendline (admin_password)
    child.timeout=300
    child.expect ('Enter the password for key file 1')
    child.sendline (key1_password)
    child.timeout=300
    child.expect ('Enter the password for key file 2')
    child.sendline (key2_password)
    if key3_password != "":
        child.timeout=300
        child.expect ('Enter the password for key file 3')
        child.sendline (key3_password)
        child.timeout=300

    child.timeout=300
    out= child.readlines()
    mk_load = "".join(out)
    print (mk_load)
    return mk_load

def mk_commit(admin_password):
    print("########### ibmcloud tke cryptounit-mk-commit ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-mk-commit', encoding='utf-8')
    child.timeout=300
    child.expect ('Enter the password for the signature key identified by:')
    child.sendline (admin_password)
    child.timeout=300
    out= child.readlines()
    mk_commit = "".join(out)
    print (mk_commit)
    return mk_commit

def mk_setitem(admin_password):
    print("########### ibmcloud tke cryptounit-mk-setimm ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-mk-setimm', encoding='utf-8')
    child.timeout=300
    child.expect ('Do you want to continue?')
    child.sendline ('y')
    child.timeout=300
    child.expect ('Enter the password for the signature key identified by:')
    child.sendline (admin_password)
    child.timeout=300
    out= child.readlines()
    mk_set = "".join(out)
    print (mk_set)
    return mk_set

def list_mk_registry():
    print("########### ibmcloud tke cryptounit-mk ########### \n")
    child = pexpect.spawn('ibmcloud tke cryptounit-mk', encoding='utf-8')
    out= child.readlines()
    mk_registry = "".join(out)
    print (mk_registry)
    return mk_registry
    