package main

import (
	"fmt"
)

// define a custom type
type ApplicationStatus int 
type JobStatus int 
type HTTPMethod int 



// database state
const (
	Applied ApplicationStatus = iota
	Interviewing
	Approved
	Rejected
)

// http method
const (
	GET HTTPMethod = iota 
	POST 
	PUT 
	DELETE
)


// job status queue
const (
	Queued JobStatus = iota
	Running
	Success 
	Failed 
	Retrying
)




func main(){
	var jobStatus JobStatus = Queued
	fmt.Println(jobStatus) // value
	fmt.Println(int(jobStatus)) // iota value


	var httpMethod HTTPMethod = GET;
	fmt.Println(httpMethod)
}


