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
	var userId string
	var password string
	var count int
	flag.StringVar(&icdId, "icdId", "", "CRN of the IBM Cloud Database service instance")
	flag.StringVar(&userId, "userId", "", "User name")
	flag.StringVar(&password, "password", "", "Password")
	flag.Parse()

	if icdId == "" || userId == "" || password == "" {
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
	taskAPI := icdClient.Tasks()
	userAPI := icdClient.Users()
	params := icdv4.UserReq{
		User: icdv4.User{
			UserName: userId,
			Password: password,
		},
	}
	task, err := userAPI.CreateUser(icdId, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Task 1 Create:", task)

	count = 0
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

	log.Println(">>>>>>>>>> Update User")

	task, err = userAPI.UpdateUser(icdId, userId, params)
	if err != nil {
		log.Fatal(err)
	}
	count = 0
	for {
		innerTask, err := taskAPI.GetTask(task.Id)
		if err != nil {
			log.Fatal(err)
		}
		count = count + 1
		log.Printf("Task : %v     %v\n", count, innerTask)
		if innerTask.Status == "completed" || innerTask.Status == "failed" || innerTask.Status == "" {
			break
		}
	}

	log.Println(">>>>>>>>>> Delete User")
	task, err = userAPI.DeleteUser(icdId, userId)
	if err != nil {
		log.Fatal(err)
	}
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
	log.Println(">>>>>>>>>> End of run")
}
