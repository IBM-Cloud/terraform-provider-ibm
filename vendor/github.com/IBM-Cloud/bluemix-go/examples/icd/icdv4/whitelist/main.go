package main

import (
	"flag"
	"log"
	"net/url"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var icdId string
	var ipAddress string
	var description string
	flag.StringVar(&icdId, "icdId", "", "CRN of the IBM Cloud Database service instance")
	flag.StringVar(&ipAddress, "ipAddress", "", "IpAddress in CIDR notation")
	flag.StringVar(&description, "description", "", "IP address description")
	flag.Parse()

	if icdId == "" {
		flag.Usage()
		os.Exit(1)
	}

	icdId = url.PathEscape(icdId)
	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	icdClient, err := icdv4.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	whitelistAPI := icdClient.Whitelists()
	taskAPI := icdClient.Tasks()

	whitelistReq := icdv4.WhitelistReq{
		WhitelistEntry: icdv4.WhitelistEntry{
			Address:     ipAddress,
			Description: description,
		},
	}
	task, err := whitelistAPI.CreateWhitelist(icdId, whitelistReq)
	log.Println(">>>>>>>> Task 1 Create Hitelist :", task)

	count := 0
	for {
		innerTask, err := taskAPI.GetTask(task.Id)
		if err != nil {
			log.Fatal(err)
		}
		count = count + 1
		log.Printf("Task : %v     %v\n", count, innerTask)
		if innerTask.Status != "running" {
			break
		}
	}

	whitelist, err := whitelistAPI.GetWhitelist(icdId)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Whitelist :", whitelist)

	task, err = whitelistAPI.DeleteWhitelist(icdId, ipAddress)
	log.Println(">>>>>>>> Task 2 Delete Whitelist :", task)

	count = 0
	for {
		innerTask, err := taskAPI.GetTask(task.Id)
		if err != nil {
			log.Fatal(err)
		}
		count = count + 1
		log.Printf("Task : %v     %v\n", count, innerTask)
		// Querying status after completion returns ''. So 'completed' or '' as completed
		if innerTask.Status != "running" {
			break
		}
	}

}
