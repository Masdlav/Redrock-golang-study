package main

import "fmt"

func Receiver(v interface{})  {
	switch v.(type) {
	case int:
		fmt.Println("这是一个int")
	case string:
		fmt.Println("这是一个string")
	case bool:
		fmt.Println("这是一个bool")
	}
}

func main()  {
	Receiver(2)

}
