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

	var icdId string
	var memory int
	var cpu int
	var disk int
	flag.StringVar(&icdId, "icdId", "", "CRN of the IBM Cloud Database service instance")
	flag.IntVar(&memory, "memory", 0, "Memory size in increments of memory scaling factor")
	flag.IntVar(&cpu, "cpu", 0, "Number of CPUs in increments of cpu scaling factor")
	flag.IntVar(&disk, "disk", 0, "Disk size in increments of disk scaling factor")
	flag.Parse()

	if icdId == "" {
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

	groupAPI := icdClient.Groups()
	params := icdv4.GroupReq{}
	if memory != 0 {
		memoryReq := icdv4.MemoryReq{AllocationMb: memory}
		params.GroupBdy.Memory = &memoryReq
	}
	if cpu != 0 {
		cpuReq := icdv4.CpuReq{AllocationCount: cpu}
		params.GroupBdy.Cpu = &cpuReq
	}
	if disk != 0 {
		diskReq := icdv4.DiskReq{AllocationMb: disk}
		params.GroupBdy.Disk = &diskReq
	}
	task, err := groupAPI.UpdateGroup(icdId, "member", params)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for {
		innerTask, err := taskAPI.GetTask(task.Id)
		if err != nil {
			log.Fatal(err)
		}
		count = count + 1
		log.Printf("Task : %v     %v\n", count, innerTask.Status)

		// Querying status after completion returns ''. So 'completed' or '' as completed
		if innerTask.Status != "running" {
			break
		}
	}

	group, err := groupAPI.GetGroups(icdId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Groups :", group)

}
