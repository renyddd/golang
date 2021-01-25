package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
	"log"
	"sync"
)

// 根据 url 的不同来返回不同的结构体
// 一个 New 带结果的函数用以返回带结果的实例和 error

type Interface interface {
	get() (Interface, error)
}

type a1 struct {
	Message string `json:"message"`
}

func NewA1() *a1 { return &a1{} }
func (*a1) get() (Interface, error) {
	var err error
	url := "http://localhost:8090/ping"
	client := resty.New()
	resStru := NewA1()
	_, err = client.R().
		SetResult(resStru).
		Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//log.Println("in a1 get fun:", resp, resStru)
	return resStru, nil
}

type l1 struct {
	List1 string `json:"list1"`
	List2 string `json:"list2"`
}

func NewL1() *l1 { return &l1{} }
func (*l1) get() (Interface, error) {
	var err error
	url := "http://localhost:8090/list"
	client := resty.New()
	resStru := NewL1()
	_, err = client.R().
		SetResult(&resStru).
		Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resStru, err
}

func GetData(url string) (Interface, error) {
	var (
		err  error
		stru Interface
	)
	switch url {
	case "ping":
		stru, err = NewA1().get()
		if err != nil {
			log.Println("err on get:", err)
			return nil, err
		}
		return stru, nil
	case "list":
		stru, err = NewL1().get()
		if err != nil {
			log.Println("err on get:", err)
			return nil, err
		}
		return stru, nil
	}

	return nil, errors.New("wrong input argument")
}

func main() {
	var (
		err error
		res Interface
		wg  sync.WaitGroup
	)
	wg.Add(1)
	go StartServer(&wg)

	gets := []string{"ping", "list"}
	for _, v := range gets {
		res, err = GetData(v)
		log.Printf("In main %t, %v, %v", res, res, err)
	}

	wg.Wait()
}

func StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"list1": "thisislist1",
			"list2": "lsisllllll2",
		})
	})

	r.Run("0.0.0.0:8090")
}
