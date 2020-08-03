package main

import (
	"fmt"
	"github.com/faasflow/goflow"
	flow "github.com/faasflow/lib/goflow"
	"log"
)

// Workload function
func doSomething(data []byte, option map[string][]string) ([]byte, error) {
	log.Println(fmt.Sprintf("you said \"%s\"", string(data)))
	return []byte(fmt.Sprintf("you said \"%s\"", string(data))), nil
}

// Workload function
func loadProfile(data []byte, option map[string][]string) ([]byte, error) {
	log.Println(fmt.Sprintf("load profile \"%s\"", string(data)))
	return []byte(fmt.Sprintf("load profile \"%s\"", string(data))), nil
}
// Workload function
func getPresignedURLForImage(data []byte, option map[string][]string) ([]byte, error) {
	log.Println(fmt.Sprintf("image url \"%s\"", string(data)))
	return []byte(fmt.Sprintf("image url \"%s\"", string(data))), nil
}
// Workload function
func detectFace(data []byte, option map[string][]string) ([]byte, error) {
	log.Println(fmt.Sprintf("detect face \"%s\"", string(data)))
	return []byte(fmt.Sprintf("detect face \"%s\"", string(data))), nil
}
// Workload function
func markProfileBasedOnStatus(data []byte, option map[string][]string) ([]byte, error) {
	log.Println(fmt.Sprintf("mask profile \"%s\"", string(data)))
	return []byte(fmt.Sprintf("mask profile \"%s\"", string(data))), nil
}

// Define provide definition of the workflow
func DefineWorkflow(f *flow.Workflow, context *flow.Context) error {
	dag := f.Dag()
	dag.Node("get-kyc-image").Apply("load-profile", loadProfile).
		Apply("get-image-url", getPresignedURLForImage)
	dag.Node("face-detect").Apply("face-detect", detectFace)
	dag.Node("mark-profile").Apply("mark-profile", markProfileBasedOnStatus)
	dag.Edge("get-kyc-image", "face-detect")
	dag.Edge("face-detect", "mark-profile")
	return nil
}

func main() {
	fs := &goflow.FlowService{
		Port:                8080,
		RedisURL:            "localhost:6379",
		WorkerConcurrency:   5,
	}
	fs.Start("myflow", DefineWorkflow)
}