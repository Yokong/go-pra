package main

import (
	"TestApp"

	"github.com/TarsCloud/TarsGo/tars"
)

type HelloImp struct{}

func (imp *HelloImp) Test() (int32, error) {
	return 0, nil
}

func (imp *HelloImp) TestHello(in string, out *string) (int32, error) {
	*out = in
	return 0, nil
}

func main() {
	imp := new(HelloImp)
	app := new(TestApp.Hello)
	cfg := tars.GetServerConfig()
	app.AddServant(imp, cfg.App+"."+cfg.Server+".HelloObj")
	tars.Run()
}
