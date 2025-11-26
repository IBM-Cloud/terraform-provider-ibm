// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISLBListenerPolicy_basic(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))

	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	priority2 := "2"
	actionPool := "forward_to_pool"
	actionListener := "forward_to_listener"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),

					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, actionPool, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),

					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, actionListener, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),

					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerPolicyConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname2, priority2, action),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority2),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPoolPolicy(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))

	priority1 := "1"
	protocol := "http"
	port := "8080"
	actionPool := "forward_to_pool"
	// actionListener := "forward_to_listener"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPoolPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, actionPool, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),

					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerListenerPolicy(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))

	priority1 := "1"
	protocol := "http"
	port := "8080"
	protocol1 := "tcp"
	port1 := "8081"
	actionlistener := "forward_to_listener"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerListenerPolicyConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, port1, protocol1, lblistenerpolicyname1, actionlistener, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),

					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPolicyRedirect_basic(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))

	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "redirect"
	priority2 := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyRedirectConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerPolicyRedirectConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname2, priority2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority2),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPolicyReject_basic(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("tflblisuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblisuat-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))
	lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandIntRange(10, 100))

	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "reject"
	priority2 := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyRejectConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},

			{
				Config: testAccCheckIBMISLBListenerPolicyRejectConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port, protocol, lblistenerpolicyname2, priority2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority2),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPolicyHttpRedirect_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	lbpolicyname := fmt.Sprintf("tflblispol%d", acctest.RandIntRange(10, 100))
	protocol1 := "https"
	port1 := "9086"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_status_code", "302"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_uri", "/example?doc=geta"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_status_code", "303"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_uri", "/example?doc=updated"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectConfigRemove(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckNoResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_uri"),
					resource.TestCheckNoResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target_https_redirect_listener"),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPolicyHttpRedirectNew_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	lbpolicyname := fmt.Sprintf("tflblispol%d", acctest.RandIntRange(10, 100))
	protocol1 := "https"
	port1 := "9086"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.http_status_code", "302"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.uri", "/example?doc=get"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfigUpdate(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.http_status_code", "301"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.uri", "/example?doc=getupdated"),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfigRemoveUri(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.http_status_code", "301"),
					resource.TestCheckResourceAttr("ibm_is_lb_listener_policy.lb_listener_policy", "target.0.uri", ""),
				),
			},
		},
	})
}

func TestAccIBMISLBListenerPolicyParameterizedRedirectNew_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflblis-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflblis-subnet-%d", acctest.RandIntRange(10, 100))
	lbname := fmt.Sprintf("tflblis%d", acctest.RandIntRange(10, 100))
	lbpolicyname := fmt.Sprintf("tflblispol%d", acctest.RandIntRange(10, 100))
	protocol1 := "https"
	port1 := "9086"
	url := "https://{host}:8080/{port}/{host}/{path}"
	urlUpdate := "{protocol}://test.{host}:80/{path}"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISLBListenerPolicyParameterizedRedirectNewConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname, url),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.http_status_code", "302"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.url", url),
				),
			},
			{
				Config: testAccCheckIBMISLBListenerPolicyParameterizedRedirectNewConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, lbname, port1, protocol1, lbpolicyname, urlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerExists("ibm_is_lb_listener_policy.lb_listener_policy", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.http_status_code", "302"),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.lb_listener_policy", "target.0.url", urlUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerPolicyDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_listener_policy" {
			continue
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		policyID := parts[2]

		getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &lbListenerID,
			ID:             &policyID,
		}

		policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		if err == nil {
			return fmt.Errorf("LBLIstenerPolicy still exists: %s %v", rs.Primary.ID, policy)
		}
	}

	return nil
}

func testAccCheckIBMISLBListenerPolicyExists(n string, policyID string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		policyID := parts[2]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &lbListenerID,
			ID:             &policyID,
		}

		policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		if err != nil {
			return err
		}

		policyID = *policy.ID

		return nil
	}
}

func testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_LB" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb           = ibm_is_lb.testacc_LB.id
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
		port         = %s
		protocol     = "%s"
	  }
	  resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	  }

	  resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		action    = "%s"
		priority  = %s
		name      = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id

}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPoolPolicyConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_LB" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb           = ibm_is_lb.testacc_LB.id
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
		port         = %s
		protocol     = "%s"
	  }
	  resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	  }

	  resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		action    = "%s"
		priority  = %s
		name      = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id

}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerListenerPolicyConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, port1, protocol1, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_LB" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb           = ibm_is_lb.testacc_LB.id
		port         = %s
		protocol     = "%s"
	  }

	  resource "ibm_is_lb_listener" "testacc_lb_listener1" {
 		 lb       = ibm_is_lb.testacc_LB.id
		 port     = %s
		 protocol = "%s"
      }

	  resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		action    = "%s"
		priority  = %s
		name      = "%s"
		target_id = ibm_is_lb_listener.testacc_lb_listener1.listener_id

}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, port1, protocol1, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, priority, action string) string {

	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  resource "ibm_is_lb" "testacc_LB" {
		name    = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	  }
	  resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb           = ibm_is_lb.testacc_LB.id
		default_pool = ibm_is_lb_pool.testacc_pool.pool_id
		port         = %s
		protocol     = "%s"
	  }
	  resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "test"
		lb             = ibm_is_lb.testacc_LB.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	  }

	  resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
		lb        = ibm_is_lb.testacc_LB.id
		listener  = ibm_is_lb_listener.testacc_lb_listener.listener_id
		action    = "%s"
		priority  = %s
		name      = "%s"
		target_id = ibm_is_lb_pool.testacc_pool.pool_id
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyRedirectConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
		target_http_status_code = 302
  		target_url              = "https://www.redirect.com"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyRedirectConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "redirect"
		priority = %s
		name = "%s"
		target_http_status_code = 302
  		target_url              = "https://www.redirect.com"

}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyRejectConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyRejectConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "reject"
		priority = %s
		name = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyHttpsRedirectConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
			}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	  
	resource "ibm_is_lb_listener" "lb_listener2"{
	lb       = ibm_is_lb.testacc_LB.id
	port     = "9087"
	protocol = "http"
	https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
	https_redirect_status_code = 301
	https_redirect_uri = "/example?doc=geta" 
	}
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
	name = "%s"
	lb = ibm_is_lb.testacc_LB.id
	listener = ibm_is_lb_listener.lb_listener2.listener_id
	action = "https_redirect"
	target_id = ibm_is_lb_pool.testacc_pool.pool_id
	priority = 2
	target_http_status_code = 302
	target_url = "https://www.redirect.com"
	target_https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
	target_https_redirect_status_code = 302
	target_https_redirect_uri = "/example?doc=geta"
	}
	
	
	resource "ibm_is_lb_pool" "testacc_pool" {
	name           = "test-pool-1"
	lb             = ibm_is_lb.testacc_LB.id
	algorithm      = "round_robin"
	protocol       = "https"
	health_delay   = 60
	health_retries = 5
	health_timeout = 30
	health_type    = "https"
	proxy_protocol = "v1"
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, lbpolicyname)

}
func testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
		https_redirect {
			http_status_code = 301
			listener {
				id = ibm_is_lb_listener.lb_listener1.listener_id
			}
			uri = "/example?doc=get"
		  }
	}
	resource "ibm_is_lb_listener" "lb_listener3"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9088"
		protocol = "https"
		certificate_instance="%s"
	}
	
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
		name = "%s"
		lb = ibm_is_lb.testacc_LB.id
		listener = ibm_is_lb_listener.lb_listener2.listener_id
		action = "https_redirect"
		target {
			http_status_code = 302
			listener {
				id = ibm_is_lb_listener.lb_listener1.listener_id
			}
			uri = "/example?doc=get"
		}
		priority = 2
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, acc.LbListerenerCertificateInstance, lbpolicyname)

}
func testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
		https_redirect {
			http_status_code = 301
			listener {
				id = ibm_is_lb_listener.lb_listener1.listener_id
			}
			uri = "/example?doc=get"
		  } 
	}
	resource "ibm_is_lb_listener" "lb_listener3"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9088"
		protocol = "https"
		certificate_instance="%s"
	}
	
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
		name = "%s"
		lb = ibm_is_lb.testacc_LB.id
		listener = ibm_is_lb_listener.lb_listener2.listener_id
		action = "https_redirect"
		target {
			http_status_code = 301
			listener {
				id = ibm_is_lb_listener.lb_listener3.listener_id
			}
			uri = "/example?doc=getupdated"
		}
		priority = 2
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, acc.LbListerenerCertificateInstance, lbpolicyname)

}
func testAccCheckIBMISLBListenerPolicyHttpsRedirectNewConfigRemoveUri(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	resource "ibm_is_lb_listener" "lb_listener2"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9087"
		protocol = "http"
	
		https_redirect {
			http_status_code = 301
			listener {
				id = ibm_is_lb_listener.lb_listener1.listener_id
			}
			uri = "/example?doc=get"
		  }
	}
	resource "ibm_is_lb_listener" "lb_listener3"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9088"
		protocol = "https"
		certificate_instance="%s"
	}
	
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
		name = "%s"
		lb = ibm_is_lb.testacc_LB.id
		listener = ibm_is_lb_listener.lb_listener2.listener_id
		action = "https_redirect"
		target {
			http_status_code = 301
			listener {
				id = ibm_is_lb_listener.lb_listener3.listener_id
			}
		}
		priority = 2
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, acc.LbListerenerCertificateInstance, lbpolicyname)

}

func testAccCheckIBMISLBListenerPolicyParameterizedRedirectNewConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname, url string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
		type = "private"
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "http"
	}
	
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
		name = "%s"
		lb = ibm_is_lb.testacc_LB.id
		listener = ibm_is_lb_listener.lb_listener1.listener_id
		action = "redirect"
		target {
			http_status_code = 302
			url = "%s"
		}
		priority = 2
	}`, vpcname, subnetname, zone, cidr, lbname, lbpolicyname, url)

}
func testAccCheckIBMISLBListenerPolicyHttpsRedirectConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
			}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	  
	resource "ibm_is_lb_listener" "lb_listener2"{
	lb       = ibm_is_lb.testacc_LB.id
	port     = "9087"
	protocol = "http"
	https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
	https_redirect_status_code = 301
	https_redirect_uri = "/example?doc=geta" 
	}
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
	name = "%s"
	lb = ibm_is_lb.testacc_LB.id
	listener = ibm_is_lb_listener.lb_listener2.listener_id
	action = "https_redirect"
	target_id = ibm_is_lb_pool.testacc_pool.pool_id
	priority = 2
	target_http_status_code = 302
	target_url = "https://www.redirect.com"
	target_https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
	target_https_redirect_status_code = 303
	target_https_redirect_uri = "/example?doc=updated"
	}
	
	
	resource "ibm_is_lb_pool" "testacc_pool" {
	name           = "test-pool-1"
	lb             = ibm_is_lb.testacc_LB.id
	algorithm      = "round_robin"
	protocol       = "https"
	health_delay   = 60
	health_retries = 5
	health_timeout = 30
	health_type    = "https"
	proxy_protocol = "v1"
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, lbpolicyname)

}

func testAccCheckIBMISLBListenerPolicyHttpsRedirectConfigRemove(vpcname, subnetname, zone, cidr, lbname, port, protocol, lbpolicyname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
			}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "lb_listener1"{
		lb       = ibm_is_lb.testacc_LB.id
		port     = "9086"
		protocol = "https"
		certificate_instance="%s"
	}
	  
	resource "ibm_is_lb_listener" "lb_listener2"{
	lb       = ibm_is_lb.testacc_LB.id
	port     = "9087"
	protocol = "http"
	https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
	https_redirect_status_code = 301
	https_redirect_uri = "/example?doc=geta" 
	}
	resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
	name = "%s"
	lb = ibm_is_lb.testacc_LB.id
	listener = ibm_is_lb_listener.lb_listener2.listener_id
	action = "forward"
	target_id = ibm_is_lb_pool.testacc_pool.pool_id
	priority = 2
	target_http_status_code = 302
	target_url = "https://www.redirect.com"
	}
	
	
	resource "ibm_is_lb_pool" "testacc_pool" {
	name           = "test-pool-1"
	lb             = ibm_is_lb.testacc_LB.id
	algorithm      = "round_robin"
	protocol       = "https"
	health_delay   = 60
	health_retries = 5
	health_timeout = 30
	health_type    = "https"
	proxy_protocol = "v1"
	}`, vpcname, subnetname, zone, cidr, lbname, acc.LbListerenerCertificateInstance, lbpolicyname)

}
