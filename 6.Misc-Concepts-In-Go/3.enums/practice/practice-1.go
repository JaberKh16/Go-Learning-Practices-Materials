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

// MarshalJSON for JSON encoding
func (j JobStatus) MarshalJSON() ([]byte, error)  {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON for JSON decoding (fixed method signature)
func (j *JobStatus) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	str := string(data)
	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	
	// Map string back to enum value
	switch str {
	case "Queued":
		*j = Queued
	case "Running":
		*j = Running
	case "Success":
		*j = Success
	case "Failed":
		*j = Failed
	case "Retrying":
		*j = Retrying
	default:
		return fmt.Errorf("invalid JobStatus: %s", str)
	}
	return nil
}


func main(){
	var jobStatus JobStatus = Queued
	fmt.Println(jobStatus) // value
	fmt.Println(int(jobStatus)) // iota value


	var httpMethod HTTPMethod = GET;
	fmt.Println(httpMethod)
}


