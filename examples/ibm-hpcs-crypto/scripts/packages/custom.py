import re
import time
import os

# This module is mainly to customise the output as the outputs of tke plugin are not in usable format
# converts raw cli output to list
def conv_str_list(cli_output):
    ansi_escape = re.compile(r'\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])')
    result = ansi_escape.sub('', cli_output)
    result1= result.split("\r\n")
    return result1

# converts list obtained from conv_str_list to dictionary
def conv_list_dict(list,keys):
    key_dict=[]
    for key in list[2:]:
        if key != "" and key.find("selected as the current signature key")== -1:
            d= dict(zip(keys,tuple(key.split())))
            key_dict.append(d)
    reversed_key_dict = key_dict[::-1] 
    return reversed_key_dict

#converts cryptounit list cli output to map of guid:cu_num
def conv_cu_list_dict(cu_list):
    if cu_list == "":
        print ("[ERROR:] There are no valid crypto units in the account")
        return
    x = re.search("Resource group:((.*[\\n?]+)*)", cu_list)
    if not x:
        print("[ERROR] Unable to fetch crupto units details")
        return
    else:
        cu_list_header=("CRYPTO UNIT NUM","SELECTED","LOCATION")
        cu_dict=dict()
        found=False
        hpcs_cu_list = x.group(1).split("\r\n\r\n")
        for cu in hpcs_cu_list:
            y = re.search("SERVICE INSTANCE:[\\s?]+(.*)",cu)
            if not y:
                continue
            found=True
            guid=y.group(1).strip()
            reaesc = re.compile(r'\x1b[^m]*m')
            hpcs_id = reaesc.sub('', guid)

            r=conv_str_list(cu)

            d= conv_list_dict(r,cu_list_header)
            cu_number=""
            for dic in reversed(d):
                if cu_number == "":
                    cu_number = dic["CRYPTO UNIT NUM"]
                else:
                    cu_number=cu_number+" "+dic["CRYPTO UNIT NUM"]
            cu_dict[hpcs_id]=cu_number
        if not found:
            print("[ERROR] No Crypto Units Found")
            return
        print("[INFO] Available Crypto unit number on this account is",cu_dict)
        return cu_dict

# Create a directory for TKEFiles with guid name
def ComputeTKEFilesDir(dir, name):
    resultDirPrefix = name
    resultDirSuffix = "_tkefiles"
    resultDir = os.path.join(dir, resultDirPrefix+resultDirSuffix)
    return resultDir

def custom_tke_files_path(output_path,hpcs_id):
    if not os.path.isdir(output_path):
        print("[ERROR] Path: '%s' to download the TKEFiles doesn't exist" % output_path)
        return
    if output_path == "":
        output_path=os.curdir
    resultDir = ComputeTKEFilesDir(output_path, hpcs_id)
    try: 
        os.makedirs(resultDir, exist_ok = True) 
        print("Directory '%s' created successfully" % resultDir) 
    except OSError as error: 
        print("Directory '%s' can not be created|already exists" % resultDir,error) 
    else:
        return resultDir

def get_cu_num(hpcs_guid,cu_num_dict):
    try:
        for k,v in cu_num_dict.items():
            if  k==hpcs_guid:
                cu_num = v
                break
    except AttributeError as error:
        print("[ERROR] Unable to fetch details from dictionary",error) 
        return
    else:
        print ("[INFO] Selected Crypto Units are: "+cu_num)
        return cu_num

def get_keynum(sig_key_list):
    try:
        sklist  =   sig_key_list.split("\r\n\r\n")
        sig_key_header=("KEYNUM","DESCRIPTION","SUBJECT KEY IDENTIFIER")
        for sk in sklist:
            reaesc = re.compile(r'\x1b[^m]*m')
            sk = reaesc.sub('', sk)

            k = re.search("KEYNUM[\\s?]+DESCRIPTION[\\s?]+SUBJECT[\\s?]+KEY[\\s?]+IDENTIFIER",sk)
            if not k:
                continue
            l = conv_str_list(sk)
            d = conv_list_dict(l,sig_key_header)
            keynum = d[0]["KEYNUM"]
    except Exception as error:
        print("[ERROR] Unable get keynum of the signature key",error)
        return
    else:
        print("[INFO] Key num of the added signature key is:"+keynum)
        return keynum
