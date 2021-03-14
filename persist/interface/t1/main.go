// 用来测试对不同网页的请求
package main

import (
	"errors"
	"log"
	"net/http"
)

type Interface interface {
	Get() error
}

type Baidu struct{}

func (Baidu) Get() error {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	log.Println(url, "success")
	return nil
}

type Bilibili struct{}

func (Bilibili) Get() error {
	url := "https://www.bilibili.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	log.Println(url, "success")
	return nil
}

func NewGet(i Interface) error {
	var err error
	switch stru := i.(type) {
	case Baidu:
		err = stru.Get()
	case Bilibili:
		err = stru.Get()
	default:
		err = errors.New("error type")
	}

	return err
}

func main() {
	var (
		err error
		ba  Baidu
		bi  Bilibili
	)

	err = NewGet(ba)
	if err != nil {
		log.Println(err)
	}

	err = NewGet(bi)
	if err != nil {
		log.Println(err)
	}

}
