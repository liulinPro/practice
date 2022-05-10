// Author: LiuLin
// Date: 2022/5/9

package main

import (
	"fmt"
	"time"
)

func main() {
	now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Add(1*time.Hour).Format("2006-01-02 15:04:05"))
	parse, _ := time.Parse("2006-01-02 15:04:00", "2022-05-09 10:34:45")
	fmt.Println(now)
	fmt.Println(parse)
	fmt.Println(parse.Before(now))
}

func main1() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	//now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2019-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println("now:", now.Format(format), "\na:", a.Format(format), "\nb:", b.Format(format))
	fmt.Println("---    a > now  >  b   -----------")
	fmt.Println("now  a   After: ", now.After(a))
	fmt.Println("now  a   Before:", now.Before(a))
	fmt.Println("now  b   After:", now.After(b))
	fmt.Println("now  b   Before:", now.Before(b))
	fmt.Println("a  b   After:", a.After(b))
	fmt.Println("a  b   Before:", a.Before(b))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())
}
