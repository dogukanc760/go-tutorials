package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic get unix time
	fmt.Printf("Şuanki Unix zamanı=> %v\n", time.Now().Unix())
	// time.Sleep(1 * time.Second)
	// fmt.Printf("Şuanki Unix zamanı=> %v\n", time.Now().Unix())

	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Custom date=> %v\n", t)
	fmt.Println("********************************************************************************************************************************")
	now := time.Now()
	fmt.Printf("Şuanki gerçel zaman => %v\n", now)
	fmt.Println("********************************************************************************************************************************")
	fmt.Printf("Şuanki ay=> %v\nŞuanki Gün=> %v\nHaftanın Günü=> %v\n", now.Month(), now.Day(), now.Weekday())
	fmt.Println("********************************************************************************************************************************")
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Yarın => %v, %v, %v, %v \n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
	fmt.Printf("********************************************************************************************************************************\n")
	longFormat := "Monday, January 2, 2018"
	fmt.Printf("Tomorrow is => %v\n", tomorrow.Format(longFormat))
	fmt.Println("********************************************************************************************************************************")
	shortFormat := "1/12/2020"
	fmt.Println("Tomorrow is => %v\n", tomorrow.Format(shortFormat))
	fmt.Println("**********************************************Farklı Tarih Operasyonları********************************************************\n")

	x := fmt.Println

	x(now)

	x("********************************************************************************************************************************************")
	xTime := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)
	x(xTime.Before(now))
	x(xTime.After(now))
	x(xTime.Equal(now))
	x("********************************************************************************************************************************************")

	// 2 tarih arasında ki fark
	diff := now.Sub(xTime)
	x(diff)
	x(diff.Hours())
	x(diff.Minutes())

}
