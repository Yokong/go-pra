package main

import (
	"TestApp"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"
)

var comm *tars.Communicator

func main() {
	comm = tars.NewCommunicator()
	obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 19000 -t 60000"
	app := new(TestApp.Hello)
	comm.StringToProxy(obj, app)
	req := "Hi Yoko!"
	var out string
	ret, err := app.TestHello(req, &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret, out)
}
