package main

import (
	"fmt"
	"time"

	"github.com/dutchcoders/go-clamd"
)

var (
	Msk, _ = time.LoadLocation("Europe/Moscow")
	MyDateOnly = "02.01.2006"
	MyRfc  = "2006-01-02T15:04:05"
	MyRfcMili  = "2006-01-02T15:04:05.000"
)

func ParseInLocation(timeVal string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(MyRfc, timeVal, loc)
}


func main() {
	// s := "/documents/file_5"
	
	c := clamd.NewClamd("/tmp/clamd.socket")

	// reader := bytes.NewReader(clamd.EICAR)
	fmt.Printf("0\n")
	tt, err := c.ScanFile("./files/Snake_River_(5mb).jpg")
	must(err)
	fmt.Printf("tt: %v\n", tt)



	fmt.Printf("conn: %v\n", "s")
}

func must(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}