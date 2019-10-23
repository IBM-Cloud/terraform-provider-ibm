package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var taskId string
	flag.StringVar(&taskId, "taskId", "", "ID of IBM Cloud Database task")
	flag.Parse()

	if taskId == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	icdClient, err := icdv4.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	taskAPI := icdClient.Tasks()
	task, err := taskAPI.GetTask(taskId)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Task :", task)

}
