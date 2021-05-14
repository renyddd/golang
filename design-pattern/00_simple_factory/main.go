package main

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	api := NewAPI(2)
	fmt.Println(api.Say("Tom"))

	api2 := NewAPI(1)
	fmt.Println(api2.Say("Bob"))

	api4 := NewAPI(4)
	fmt.Println(api4.Say("anonmous"))
}
