from packages import object_storage as object_storage
import os

api_key = os.environ.get("API_KEY","")
cos_service_crn = os.environ.get("COS_SERVICE_CRN","")
endpoint = os.environ.get("ENDPOINT","")
bucket_name = os.environ.get("BUCKET","")
input_file_name = os.environ.get("INPUT_FILE_NAME","")

input_file_path=object_storage.download_file_cos(api_key,cos_service_crn,endpoint,bucket_name,input_file_name)
