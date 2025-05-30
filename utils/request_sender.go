package utils

import (
	"log"
	"net/http"
	"time"
)

type RequestSender interface {
	Start(fn DoRequestCallback)
}

type DoRequestCallback = func(c http.Client, reqID int)
type RequestSenderMode = int

const (
	ParallelMode RequestSenderMode = iota
	SequentialMode
)

type RequestSenderConfig struct {
	NumOfRequest int
	Mode         RequestSenderMode
	Jitter       time.Duration
}

type requestSender struct {
	config RequestSenderConfig
}

func NewRequestSender(config RequestSenderConfig) *requestSender {
	rs := &requestSender{
		config: config,
	}

	rs.applyConfig()

	return rs
}

func (r *requestSender) Start(fn DoRequestCallback) {
	c := http.Client{}
	for i := range r.config.NumOfRequest {
		if r.config.Mode == ParallelMode {
			go fn(c, i)
		} else {
			fn(c, i)
		}

		// Wait a while before send new request
		time.Sleep(r.config.Jitter)
	}

	log.Println("[INFO] all request sent")
	log.Println("[INFO] press Ctrl + c to stop the app")
}

// applyConfig will set default value for config when it missing
func (r *requestSender) applyConfig() {
	if r.config.NumOfRequest == 0 {
		r.config.NumOfRequest = 10
	}

	if r.config.Jitter == 0 {
		r.config.Jitter = time.Second
	}
}
