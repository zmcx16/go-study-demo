package demo

import (
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testResponse = "KONOSUBA!"
)

func newTestHTTPServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(NewRateLimiter(1))
	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Length", strconv.Itoa(len(testResponse)))
		c.String(200, testResponse)
	})
	return router
}

func randomListener() (net.Listener, string) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	return listener, listener.Addr().String()
}

func TestRateLimiter_ExceedLimit_ShouldReturnHTTPStatusCode429(t *testing.T) {
	// setup
	listener, bindString := randomListener()
	httpURL := "http://" + bindString

	client := &http.Client{}
	req, _ := http.NewRequest("GET", httpURL, nil)

	r := newTestHTTPServer()
	go r.RunListener(listener)
	defer listener.Close()

	// act
	firstResp, firstErr := client.Do(req)
	SecondResp, SecondErr := client.Do(req)

	// assert
	assert.Nil(t, firstErr)
	assert.Equal(t, 200, firstResp.StatusCode)
	assert.Nil(t, SecondErr)
	assert.Equal(t, 429, SecondResp.StatusCode)
}

func TestRateLimiter_NotExceedLimit_ShouldReturnHTTPStatusCode200(t *testing.T) {
	// setup
	listener, bindString := randomListener()
	httpURL := "http://" + bindString

	client := &http.Client{}
	req, _ := http.NewRequest("GET", httpURL, nil)

	r := newTestHTTPServer()
	go r.RunListener(listener)
	defer listener.Close()

	// act
	firstResp, firstErr := client.Do(req)
	time.Sleep(time.Second)
	SecondResp, SecondErr := client.Do(req)

	// assert
	assert.Nil(t, firstErr)
	assert.Equal(t, 200, firstResp.StatusCode)
	assert.Nil(t, SecondErr)
	assert.Equal(t, 200, SecondResp.StatusCode)
}
