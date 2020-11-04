from packages import object_storage as object_storage
from packages import custom as custom
import time
import shutil
import os

api_key = os.environ.get("API_KEY","")
cos_service_crn = os.environ.get("COS_SERVICE_CRN","")
endpoint = os.environ.get("ENDPOINT","")
bucket_name = os.environ.get("BUCKET","")

tke_files_path = os.environ.get("CLOUDTKEFILES","")
hpcs_guid = os.environ.get("HPCS_GUID","")

resultDir = custom.custom_tke_files_path(tke_files_path,hpcs_guid)

zip_file_name =str(time.strftime("%y%m%d%S%H%M%S"))+"_"+hpcs_guid+"_tkefiles"
zip_file_path = shutil.make_archive(zip_file_name, 'zip', resultDir)


object_storage.upload_file_cos(api_key,cos_service_crn,endpoint,bucket_name,zip_file_path,zip_file_name)
