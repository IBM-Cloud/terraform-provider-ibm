# Deploying and Managing Openshift Data Foundation on Satellite

This example shows how to deploy and manage the Openshift Data Foundation (ODF) on IBM Cloud Satellite based RedHat Openshift cluster.

#### Please Select the ODF Template you wish to install on your ROKS Satellite Cluster and follow the documentation.

- odf-remote -  Choose this template if you have a CSI driver installed in your cluster. For example, the azuredisk-csi-driver driver. You can use the CSI driver to dynamically provision storage volumes when deploying ODF.
    
- odf-local - Choose this template when you have local storage available to your worker nodes. If your storage volumes are visible when running lsblk, you can use these disks when deploying ODF if they are raw and unformatted.