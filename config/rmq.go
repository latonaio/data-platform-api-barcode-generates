package config

import (
	"fmt"
	"os"
	"strings"
)

func newRMQ() *RMQ {
	return &RMQ{
		user:                os.Getenv("RMQ_USER"),
		pass:                os.Getenv("RMQ_PASS"),
		addr:                os.Getenv("RMQ_ADDRESS"),
		port:                os.Getenv("RMQ_PORT"),
		vhost:               os.Getenv("RMQ_VHOST"),
		queueFrom:           os.Getenv("RMQ_QUEUE_FROM"),
		queueToSQL:          getEnvStrings("RMQ_QUEUE_TO_SQL"),
		queueToResponse:     os.Getenv("NESTJS_DATA_CONNECTION_REQUEST_CONTROL_MANAGER_CONSUME"),
		sessionControlQueue: os.Getenv("RMQ_SESSION_CONTROL_QUEUE"),
	}
}

type RMQ struct {
	user  string
	pass  string
	addr  string
	port  string
	vhost string

	queueFrom       string
	queueToSQL      []string
	queueToResponse string

	sessionControlQueue string
}

func (c *RMQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.user, c.pass, c.addr, c.port, c.vhost)
}

func (c *RMQ) QueueFrom() string {
	return c.queueFrom
}
func (c *RMQ) QueueToSQL() []string {
	return c.queueToSQL
}

func (c *RMQ) QueueToResponse() string {
	return c.queueToResponse
}

func (c *RMQ) SessionControlQueue() string {
	return c.sessionControlQueue
}

func getEnvStrings(key string) []string {
	rawVal := os.Getenv(key)
	rawVal = strings.ReplaceAll(rawVal, "\\ ", "$THIS_SECTION_IS_SPACE")
	rawVal = strings.ReplaceAll(rawVal, " ", "")
	rawVal = strings.ReplaceAll(rawVal, "$THIS_SECTION_IS_SPACE", " ")
	val := strings.Split(rawVal, ",")
	return val
}
