// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSnapshotDatasource_basic(t *testing.T) {
	var snapshot string
	snpName := "data.ibm_is_snapshot.ds_snapshot"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	// 	publicKey := strings.TrimSpace(`
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	// `)
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDCZNT0p7DKLKoXshjO0xHuPVpxAHRYhTkVI+yklxPoF/1UuIy+8cczEUnEoIZ81hHVnyY8Ltr75Xr2/yI1g+a6t+xwgdp74f5+R69AuIuywsoLdt8mda2iFryrPs5eebhSBo9c2dGya/uv1pR8o/ED9xrIB4QQ+vrT8uuss3GIP6lR5w9vJqWgeB4I4TqGu4giWLkeMyBKVVUn/UoAqoa7uD7deFXkmN73j227AyCV+lx8MBFQZKvnsglJxgSnjwlQP2iWCyLGvDV1j0lldXAvjkxYRl8H9ilrOk5KrCpqVtlsXcXqqBTeCEBpI0KOHIAl/JJ+YFc3+YTQq0waGB82CaqQCnJEy6bBpdPz4tge16/2BehY0fY5AG8sqb7Or91ilsE0nnuR3Tzk5l9YS8ftqyjYAoRbBy3P7AF433qbCgoclndp5Mcy5EsEJ//kme0u4n8ZF20ZS8oqKK1QIfOio8U2er/SSF3UAHhT2n/hKaeGeXv0kFwCsExjP1iZsHk= bhaveshshrivastav@Bhaveshs-MacBook-Pro.local
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSnapshotExists("ibm_is_snapshot.testacc_snapshot", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", name1),
					// resource.TestCheckResourceAttrSet(snpName, "delatable"), // commented as it is deprecated
					resource.TestCheckResourceAttrSet(snpName, "href"),
					resource.TestCheckResourceAttrSet(snpName, "crn"),
					resource.TestCheckResourceAttrSet(snpName, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(snpName, "encryption"),
					resource.TestCheckResourceAttrSet(snpName, "captured_at"),
				),
			},
		},
	})
}

func testDSCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, sname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           					= "%s"
		vpc             				= ibm_is_vpc.testacc_vpc.id
		zone            				= "%s"
		total_ipv4_address_count 		= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    	= "%s"
		image   	= "%s"
		profile 	= "%s"
		primary_network_interface {
		  subnet    = ibm_is_subnet.testacc_subnet.id
		}
		vpc  		= ibm_is_vpc.testacc_vpc.id
		zone 		= "%s"
		keys 		= [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet 	= ibm_is_subnet.testacc_subnet.id
		  name   	= "eth1"
		}
	  }
	resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
		}
	data "ibm_is_snapshot" "ds_snapshot" {
		depends_on 	= [ibm_is_snapshot.testacc_snapshot]
		name 		= "%s"
	}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, sname, sname)
}
