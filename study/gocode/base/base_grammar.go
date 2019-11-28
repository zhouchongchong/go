package base

import (
	"log"
	"time"
)

func test1()  {
	x, y := 1, 2

	defer func(a int) {
		println("defer x, y = ", a, y)
	}(x)

	x += 100
	y += 200
	println(x, y)
}

func test22() {
	//a := 20
	//var add1 *int  = &a

	//println(&a)
	//println(a)
	//var add2   = &add1


	//println(*add1)
	//println(*add2)
	//test(add1)

	//fmt.Printf("%T \n",a)
	//option := newOption()
	//option.port = 8085
	//fmt.Printf("%T \n",option)
	//fmt.Printf("%v \n",option.port)
	//intstr := getInt()
	////fmt.Printf("intstr %T \n",intstr)
	//
	//println(*intstr)
	println("test: ",test11())
}

func test11() (z int)  {
	defer func() {
		println("defer :", z)
		a := 100 + z
		println("defer a:", a)

	}()

	return 100
}

func getInt() *int {
	a :=20
	return &a
}

func server(option *serverOption)  {

}
type serverOption struct {
	address string
	port	int
	path 	string
	timeout time.Duration
	log 	*log.Logger
}

func newOption() serverOption  {
	return serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 10,
		log:     nil,
	}
}
func test(p *int)  {
	//println(p)
}