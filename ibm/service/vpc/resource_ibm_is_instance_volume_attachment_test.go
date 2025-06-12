// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceVolumeAttachment_basic(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	iops1 := int64(600)
	iops2 := int64(900)

	capacity1 := int64(20)
	capacity2 := int64(22)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity1, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "600"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity1, iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "900"),
				),
			},

			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity2, iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity2)),
				),
			},
		},
	})
}
func TestAccIBMISInstanceVolumeAttachment_sdpbasic(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	volUpdateName := fmt.Sprintf("tf-vol-update-%d", acctest.RandIntRange(10, 100))
	iops1 := int64(10000)
	iops2 := int64(28000)

	capacity1 := int64(1000)
	capacity2 := int64(22000)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentSdpConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity1, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "10000"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentSdpConfig(vpcname, subnetname, sshname, publicKey, name, attName, volUpdateName, autoDelete, capacity1, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "10000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "volume_name", volUpdateName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentSdpConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity1, iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "28000"),
				),
			},

			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentSdpConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity2, iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity2)),
				),
			},
		},
	})
}
func TestAccIBMISInstanceVolumeAttachment_bandwidth(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	volUpdateName := fmt.Sprintf("tf-vol-update-%d", acctest.RandIntRange(10, 100))
	iops1 := int64(10000)

	capacity1 := int64(1000)
	bandwidth1 := int64(1800)
	bandwidth2 := int64(2400)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete, capacity1, iops1, bandwidth1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "10000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "bandwidth", fmt.Sprintf("%d", bandwidth1)),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, attName, volUpdateName, autoDelete, capacity1, iops1, bandwidth2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "iops", "10000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "volume_name", volUpdateName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "bandwidth", fmt.Sprintf("%d", bandwidth2)),
				),
			},
		},
	})
}
func TestAccIBMISInstanceVolumeAttachment_crn(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentCrnConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, autoDelete),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_volume_attachment.testacc_att", "iops"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceVolumeAttachment_userTag(t *testing.T) {
	var instanceVolAtt string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	attName := fmt.Sprintf("tf-volatt-%d", acctest.RandIntRange(10, 100))
	autoDelete := true
	volName := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	iops1 := int64(600)
	capacity1 := int64(20)
	userTag1 := "tag-0"
	userTag2 := "tag-1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentUsertagConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, userTag1, autoDelete, capacity1, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "capacity", fmt.Sprintf("%d", capacity1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "tags.0", userTag1),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentUsertagConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, userTag2, autoDelete, capacity1, iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceVolumeAttachmentExists("ibm_is_instance_volume_attachment.testacc_att", instanceVolAtt),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "name", attName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "delete_volume_on_instance_delete", fmt.Sprintf("%t", autoDelete)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_volume_attachment.testacc_att", "tags.0", userTag2),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumeAttachmentDestroy(s *terraform.State) error {

	instanceC, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_volume_attachment" {
			continue
		}
		getinsvolAttOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := instanceC.GetInstanceVolumeAttachment(getinsvolAttOptions)

		if err == nil {
			return fmt.Errorf("instance volume attachment still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func parseVolAttTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

func testAccCheckIBMISInstanceVolumeAttachmentExists(n string, instanceVolAtt string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		instanceId, id, err := parseVolAttTerraformID(rs.Primary.ID)
		instanceC, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getinsVolAttOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
			ID:         &id,
			InstanceID: &instanceId,
		}
		foundins, _, err := instanceC.GetInstanceVolumeAttachment(getinsVolAttOptions)
		if err != nil {
			return err
		}
		instanceVolAtt = *foundins.ID
		return nil
	}
}

func testAccCheckIBMISInstanceVolumeAttachmentConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName string, autoDelete bool, capacity, iops int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance = ibm_is_instance.testacc_instance.id
	
		name 			= "%s"
		profile 		= "custom"
		capacity	 	= %d
		iops			= %d
	
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
	  }
	 
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, attName, capacity, iops, autoDelete, volName)
}

func testAccCheckIBMISInstanceVolumeAttachmentSdpConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName string, autoDelete bool, capacity, iops int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance = ibm_is_instance.testacc_instance.id
	
		name 			= "%s"
		profile 		= "sdp"
		capacity	 	= %d
		iops			= %d
	
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
	  }
	 
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, attName, capacity, iops, autoDelete, volName)
}
func testAccCheckIBMISInstanceVolumeAttachmentBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName string, autoDelete bool, capacity, iops, volBandwidth int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance = ibm_is_instance.testacc_instance.id
	
		name 			= "%s"
		profile 		= "sdp"
		capacity	 	= %d
		iops			= %d
	
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
		bandwidth = %d
	  }
	 
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, attName, capacity, iops, autoDelete, volName, volBandwidth)
}

func testAccCheckIBMISInstanceVolumeAttachmentCrnConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName string, autoDelete bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	}
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	}
	resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "tf-test-snapshot"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
	}
	resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance 		= ibm_is_instance.testacc_instance.id
		name 			= "%s"
		profile 		= "general-purpose"
		snapshot_crn 	= ibm_is_snapshot.testacc_snapshot.crn
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
	}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, attName, autoDelete, volName)
}

func testAccCheckIBMISInstanceVolumeAttachmentUsertagConfig(vpcname, subnetname, sshname, publicKey, name, attName, volName, usertag string, autoDelete bool, capacity, iops int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_instance_volume_attachment" "testacc_att" {
		instance = ibm_is_instance.testacc_instance.id
	
		name 			= "%s"
		profile 		= "custom"
		capacity	 	= %d
		iops			= %d
		tags = ["%s"]
	
		delete_volume_on_instance_delete = %t
		volume_name = "%s"
	  }
	 
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, attName, capacity, iops, usertag, autoDelete, volName)
}
