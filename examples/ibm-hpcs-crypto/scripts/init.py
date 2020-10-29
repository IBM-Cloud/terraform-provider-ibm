import json
import pexpect
import os
from packages import object_storage as object_storage
from packages import hpcs as hpcs
from packages import custom as custom
import re

# --------------------------------------------
# GET INPUTS From Null resource
# --------------------------------------------
inputfile = os.environ.get("INPUT_FILE","")
tke_files_path=os.environ.get("CLOUDTKEFILES","")
hpcs_guid = os.environ.get("HPCS_GUID","")

# print(pexpect.run('ibmcloud iam oauth-tokens'))

if inputfile == "":
    print("[ERROR] Unable to read file or Provided file is empty")
try:
    data= json.loads(inputfile)
except Exception as error:
    print("[ERROR] Unable to read data from  input file", error)
else:

    # --------------------------------------------
    # Declare JSON input key-values
    # --------------------------------------------
    print(data)
    admin_name=data["admin_name"]
    admin_password=data["admin_password"]
    threshold_value = data["threshold"]
    rev_threshold_value = data["rev_threshold"]
    random_mk=[]
    custom_mk=[]
    if "random_mk" in data.keys():
        random_mk=data["random_mk"]
    if "custom_mk" in data.keys():
        custom_mk=data["custom_mk"]
    # ----------------------------------------------------------------------------------------
    # Create custom directory in the output path provided inorder to avoid misplacement of data
    # ----------------------------------------------------------------------------------------
    resultDir = custom.custom_tke_files_path(tke_files_path,hpcs_guid)
    os.environ['CLOUDTKEFILES'] = resultDir
    os.system("echo [INFO] TKE Files will be located at $CLOUDTKEFILES")

    # -----------------------------------------------------------------------------------
    # List Crypto units and format the output to get guid-crypto_unit_num key-val pair
    # -----------------------------------------------------------------------------------
    cu_list= hpcs.list_crypto_units()

    cu_num_dict = custom.conv_cu_list_dict(cu_list)

    cu_num = custom.get_cu_num(hpcs_guid,cu_num_dict)

    # --------------------------------------------
    # Add Crypto unit
    # --------------------------------------------
    hpcs.crypto_unit_add(cu_num)

    # --------------------------------------------
    # List signkeys 
    # --------------------------------------------
    sig_key_list = hpcs.list_sig_keys()

    # ---------------------------------------------------------
    # Add signkeys and format output to to get key_num
    # ---------------------------------------------------------

    added_sigkey= hpcs.sigkey_add(admin_name,admin_password)

    sig_key_list = hpcs.list_sig_keys()

    key_num = custom.get_keynum(sig_key_list)

    # --------------------------------------------
    # Select signkeys 
    # --------------------------------------------
    selected_sigkey =  hpcs.sigkey_select(key_num,admin_password)

    # --------------------------------------------
    # Load Crypto unit administrator
    # --------------------------------------------

    added_admin = hpcs.admin_add(key_num,admin_password)

    hpcs.list_admins()
    # --------------------------------------------
    # Set Threshold for administrator
    # --------------------------------------------

    thres_set= hpcs.threshold_set(threshold_value,rev_threshold_value,admin_password)

    hpcs.list_thresholds()

    # --------------------------------------------
    # Add Master keys
    # --------------------------------------------

    mk_total=list()
    if len(random_mk)>0:
        for mk in random_mk:
            mk_total.append(mk)
            description = mk["description"]
            password = mk["password"]
            mk_random = hpcs.mk_random_add(description,password)
            

    if len(custom_mk) >0:
        for mk in custom_mk:
            mk_total.append(mk)
            description = mk["description"]
            password = mk["password"]
            key = mk["key"]
            mk_custom = hpcs.mk_custom_add(description,password,key)
            
    reversed_mk_total = mk_total[::-1] 

    # ---------------------------------------------------------------------
    # List and format Master keys output to get key num and password values
    # ---------------------------------------------------------------------
    try:
        mks = hpcs.list_mks()
        mk_header=("KEYNUM","DESCRIPTION","VERIFICATION PATTERN")
        mk_list = custom.conv_str_list(mks)
        mk_list_dict = custom.conv_list_dict(mk_list,mk_header)
        mk_list_even = mk_list_dict[1::2]                       #gets even index in a list.. as the output is not in json format.. we are trying to get the keynum of the mk key..
        mknum=""
        mk_pw=list()
        for i in range(len(reversed_mk_total)):
            if i>2:
                break
            if mknum == "":
                mknum = mk_list_even[i]["KEYNUM"]
            else:
                mknum=mknum+" "+mk_list_even[i]["KEYNUM"]
            mk_pw.append(reversed_mk_total[i]["password"])

        mk_keynum = mknum[::-1]
        keypassword = mk_pw[::-1]
        key3_password =""
        key1_password = keypassword[0]
        key2_password = keypassword[1]
        if len(keypassword)>2:
            key3_password=keypassword[2]
    except Exception as error:
        print ("[ERROR] Unable to get key num and key passwords",error)
    else:
        print("[INFO] Keynum of the master keys that are to be loaded are: "+mk_keynum)

    # --------------------------------------------
    # Load Master keys
    # --------------------------------------------

    master_key_load = hpcs.mk_load(mk_keynum,admin_password,key1_password,key2_password,key3_password)

    # --------------------------------------------
    # Commit Master keys
    # --------------------------------------------

    master_key_commit = hpcs.mk_commit(admin_password)

    # --------------------------------------------
    # Set and list Master registry
    # --------------------------------------------

    master_key_set_item = hpcs.mk_setitem(admin_password)

    mk_registry = hpcs.list_mk_registry()
