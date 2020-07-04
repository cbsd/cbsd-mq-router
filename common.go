package main

import "time"

type Comment struct {
	Command string
	JobID uint64
	Date time.Time
	CommandArgs map[string]interface{}
}

type CommentProtocol interface {
	Decode(encodedComment []byte) (*Comment, error)
	Encode(comment *Comment) ([]byte, error)
}

type CommentProcessor interface {
	DoProcess(comment *Comment) error
}

type CbsdTask struct {
	DskGuid		string
	ErrCode		int
	Guid		string
	Message		string
	Progress	int
	Vnc		string
}
