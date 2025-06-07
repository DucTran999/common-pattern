package utils

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/bxcodec/faker"
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
	if r.config.Mode == ParallelMode {
		r.sendParallel(fn)
	} else {
		r.sendSequential(fn)
	}

	log.Println("[INFO] all request sent")
	log.Println("[INFO] press Ctrl + C to stop the app")
}

func (r *requestSender) sendParallel(fn DoRequestCallback) {
	wg := sync.WaitGroup{}

	for i := range r.config.NumOfRequest {
		c := newHTTPClientWithSourceIP(faker.IPV4)
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fn(*c, idx)
		}(i)

		// Wait a while before send new request
		time.Sleep(r.config.Jitter)
	}

	wg.Wait()
}

func (r *requestSender) sendSequential(fn DoRequestCallback) {
	c := http.Client{}

	for i := range r.config.NumOfRequest {
		fn(c, i)

		// Wait a while before send new request
		time.Sleep(r.config.Jitter)
	}
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

func newHTTPClientWithSourceIP(sourceIP string) *http.Client {
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP: net.ParseIP(sourceIP),
		},
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
	}

	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		DisableKeepAlives:   false,
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}
}
