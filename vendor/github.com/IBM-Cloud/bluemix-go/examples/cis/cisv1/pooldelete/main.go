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

    var pool_id string
    flag.StringVar(&pool_id, "pool_id", "", "Id for pool")

    flag.Parse()

    if pool_id == "" || cis_id == "" {
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
    poolsAPI := cisClient.Pools()

    log.Println(">>>>>>>>>  Pool read")
    myPoolPtr, err := poolsAPI.GetPool(cis_id, pool_id)

    if err != nil {
        log.Fatal(err)
    }

    myPool := *myPoolPtr
                
    log.Println("Pool Details by ID:", myPool)

    log.Println(">>>>>>>>>  Pool delete")
    err = poolsAPI.DeletePool(cis_id, pool_id)
    if err != nil {
        log.Fatal(err)
    }

}