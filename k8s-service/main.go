package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type sliceValue []string

func (s *sliceValue) Set(val string) error {
	values := strings.Split(val, ",")
	*s = append(*s, values...)
	return nil
}

func (s *sliceValue) String() string {
	return strings.Join(*s, ",")
}

var (
	serviceList      sliceValue
	service          string
	code             int
	clientTimeout    int
	serverHandleTime int
)

func init() {
	flag.StringVar(&service, "service", "servicea", "service name")
	flag.Var(&serviceList, "serviceList", "servicea,serviceb,servicec")
	flag.IntVar(&code, "code", 200, "response return code")
	flag.IntVar(&clientTimeout, "clientTimeout", 10, "client time out")
	flag.IntVar(&serverHandleTime, "serverHandleTime", 0, "simulate server processing time")
	flag.Parse()
}

func main() {
	serviceIndex := 0
	for index := range serviceList {
		if serviceList[index] == service {
			serviceIndex = index
			break
		}
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Duration(serverHandleTime) * time.Second)

		if len(serviceList) == serviceIndex+1 {
			c.String(code, os.Getenv("POD_NAME"))
			return
		}
		nextServiceName := serviceList[serviceIndex+1]

		req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/", nextServiceName), nil)
		resp, err := (&http.Client{
			Timeout: time.Duration(clientTimeout) * time.Second,
		}).Do(req)
		if err != nil {
			fmt.Println("do request failed", err)
			c.String(code, os.Getenv("POD_NAME"))
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)
		c.String(code, os.Getenv("POD_NAME")+"-->"+string(body))
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "Received a POST request")
	})

	r.Run(":8080")
}
