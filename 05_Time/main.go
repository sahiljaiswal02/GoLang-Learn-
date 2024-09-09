package main

import "fmt"
       "time"

func main() {
	fmt.Println("Welcome to the time study of golang")  

	presentTime := time.Now()
	fmt.Println(presentTime) // 2020-08-10 23:23:00.000000 +0530 IST m=+0.000984201

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))  // 10-08-2020 23:23:00 Monday

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate) //	2020-08-10 23:23:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday")) // 10-08-2020 23:23:00 Monday
}