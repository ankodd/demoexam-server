package models

type Status string

const (
	Waiting  Status = "waiting"
	Working  Status = "working"
	Accepted Status = "done"
)
