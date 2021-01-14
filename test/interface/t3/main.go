package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
)

// TODO: 待修改返回值接口
type dataResInterface interface {
	implement()
}

type client struct {
	rcli *resty.Client
}

func NewClient() *client {
	return &client{resty.New()}
}

type a1 struct {
	Message string `json:"message"`
}
func (*a1) implement() {}

type a2 struct {
	List1 string `json:"list1"`
	List2 string `json:"list2"`
}
func (*a2) implement() {}

func (c *client) Get(url string) (dataResInterface, error) {
	var err error

	switch url {
	case "a1":
		t := &a1{}
		_, err = c.rcli.R().
			SetResult(&t).
			Get("http://localhost:8090/a1")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		//log.Printf("Get success %v\n", t)
		return t, nil
	case "a2":
		t := new(a2)
		_, err = c.rcli.R().
			SetResult(&t).
			Get("http://localhost:8090/a2")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return t, nil
	}
	return nil, err
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go StartServer(&wg)

	client := NewClient()
	s := []string{"a1", "a2"}

	for _, v := range s {
		res, err := client.Get(v)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res, err)
	}
	wg.Wait()
}

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	r := gin.Default()
	r.GET("/a1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/a2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"list1": "thisislist1",
			"list2": "lsisllllll2",
		})
	})

	r.Run("0.0.0.0:8090")
}
