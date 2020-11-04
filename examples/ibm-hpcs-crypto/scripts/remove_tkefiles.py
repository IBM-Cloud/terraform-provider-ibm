from packages import custom as custom
import os
import shutil



tke_files_path=os.environ.get("CLOUDTKEFILES","")
hpcs_guid = os.environ.get("HPCS_GUID","")
inputfile = os.environ.get("INPUT_FILE_NAME","")
if inputfile != "":
    os.remove(inputfile)
resultDir = custom.custom_tke_files_path(tke_files_path,hpcs_guid)
try:
    shutil.rmtree(resultDir)
except Exception as error:
    print ("[ERROR] Failed to delete tke files")
else:
    print ("[INFO] Succesfully deleted all the CLOUDTKEFILES")

