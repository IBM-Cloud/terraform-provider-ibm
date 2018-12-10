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

    var zone_id string
    flag.StringVar(&zone_id, "zone_id", "", "zone_id for zone")

    flag.Parse()

    if zone_id == "" || cis_id == "" {
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

    log.Println(">>>>>>>>>  Zone read")
    myZonePtr, err := zonesAPI.GetZone(cis_id, zone_id)

    if err != nil {
        log.Fatal(err)
    }

    myZone := *myZonePtr
                
    log.Println("Zone Details by ID:", myZone)

    log.Println(">>>>>>>>>  Zone delete")
    err = zonesAPI.DeleteZone(cis_id, zone_id)
    if err != nil {
        log.Fatal(err)
    }

}