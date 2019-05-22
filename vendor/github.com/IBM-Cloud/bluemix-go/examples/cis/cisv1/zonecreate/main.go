package main

import (
    "flag"
    "log"
    "os"

    "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
    "github.com/IBM-Cloud/bluemix-go/session"
    "github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

    var cis_id string
    flag.StringVar(&cis_id, "cis_id", "", "CRN of the CIS service instance")

    var domain string
    flag.StringVar(&domain, "domain", "", "DNS domain name for zone")

    flag.Parse()

    if domain == "" || cis_id == "" {
        flag.Usage()
        os.Exit(1)
    }

    trace.Logger = trace.NewLogger("true")
    sess, err := session.New()
    if err != nil {
        log.Fatal(err)
    }

    cisClient, err := cisv1.New(sess)
    if err != nil {
        log.Fatal(err)
    }
    zonesAPI := cisClient.Zones()
    params := cisv1.ZoneBody{Name: domain}
    myZonePtr, err := zonesAPI.CreateZone(cis_id, params)
       
    if err != nil {
        log.Fatal(err)
    }
    myZone := *myZonePtr
    zoneId := myZone.Id
    myZonePtr, err = zonesAPI.GetZone(cis_id, zoneId)

    if err != nil {
        log.Fatal(err)
    }

}