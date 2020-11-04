from ibm_botocore.client import Config
import ibm_boto3
import os

def download_file_cos(api_key,service_crn,endpoint,bucket,input_file):  
    cos = ibm_boto3.client(service_name='s3',
    ibm_api_key_id=api_key,
    ibm_service_instance_id=service_crn,
    ibm_auth_endpoint="https://iam.cloud.ibm.com/identity/token",
    config=Config(signature_version='oauth'),
    endpoint_url=endpoint)

    try:
        _=cos.download_file(Bucket=bucket,Key=input_file,Filename=input_file)
    except Exception as e:
        print(Exception, e)
    else:
        print('File Downloaded at: '+os.path.abspath(input_file))
        return input_file


def upload_file_cos(api_key,service_crn,endpoint,bucket,zip_file_path,output_zip): 
    cos = ibm_boto3.client(service_name='s3',
    ibm_api_key_id=api_key,
    ibm_service_instance_id=service_crn,
    ibm_auth_endpoint="https://iam.cloud.ibm.com/identity/token",
    config=Config(signature_version='oauth'),
    endpoint_url=endpoint)

    try:
        _=cos.upload_file(Filename=zip_file_path, Bucket=bucket,Key=output_zip)
    except Exception as e:
        print(Exception, e)
    else:
        print('File Uploaded as: "%s" in cos_bucket "%s" ' % (output_zip, bucket))
        os.remove(zip_file_path)
