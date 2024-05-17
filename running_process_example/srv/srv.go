package srv

import "fmt"

func ThisIsSrvPublicFunction() {
	fmt.Println("this is public function from srv package of srv.go file")
}

func init() {
	fmt.Println("run init function from srv/srv.go")
}
