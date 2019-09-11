package main

import (
	"TestApp"
	"fmt"
	"testing"

	"github.com/TarsCloud/TarsGo/tars"
)

func BenchmarkTarClient(b *testing.B) {
	comm = tars.NewCommunicator()
	obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 19002 -t 60000"
	app := new(TestApp.Hello)
	comm.StringToProxy(obj, app)
	b.ResetTimer()
	req := "Hi Yoko!"
	for i := 0; i < b.N; i++ {
		var out string
		ret, err := app.TestHello(req, &out)
		fmt.Println(ret, err)
	}
}
