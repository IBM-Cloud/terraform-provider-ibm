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
    var dns_name string
    flag.StringVar(&dns_name, "dns_name", "", "dns_name for zone")


    flag.Parse()

    if zone_id == "" || cis_id == "" || dns_name == "" {
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
    dnsAPI := cisClient.Dns()

   
    log.Println(">>>>>>>>>  Dns create")
    params := cisv1.DnsBody{
                        Name: "www.example.com",
                        DnsType: "A",
                        Content: "192.168.127.127",
                    }
    myDnsPtr, err := dnsAPI.CreateDns(cis_id, zone_id, params)
       
    if err != nil {
        log.Fatal(err)
    }

    myDns := *myDnsPtr
    dnsId := myDns.Id
    log.Println("Dns create :", myDns)

    log.Println(">>>>>>>>>  Dns read")
    myDnsPtr, err = dnsAPI.GetDns(cis_id, zone_id, dnsId)

    if err != nil {
        log.Fatal(err)
    }

     myDns = *myDnsPtr
                
    log.Println("Dns Details by ID:", myDns)

    log.Println(">>>>>>>>>  Dns delete")
    err = dnsAPI.DeleteDns(cis_id, zone_id, dnsId)
    if err != nil {
        log.Fatal(err)
    }
}