package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func (p *Person) String() string {
	return p.Name
}

func appendTime(b []byte, t time.Time) []byte {
	const days = "SunMonTueWedThuFriSat"
	const months = "JanFebMarAprMayJunJulAugSepOctNovDec"

	t = t.UTC()
	yy, mm, dd := t.Date()
	hh, mn, ss := t.Clock()
	day := days[3*t.Weekday():]
	mon := months[3*(mm-1):]

	return append(b,
		day[0], day[1], day[2], ',', ' ',
		byte('0'+dd/10), byte('0'+dd%10), ' ',
		mon[0], mon[1], mon[2], ' ',
		byte('0'+yy/1000), byte('0'+(yy/100)%10), byte('0'+(yy/10)%10), byte('0'+yy%10), ' ',
		byte('0'+hh/10), byte('0'+hh%10), ':',
		byte('0'+mn/10), byte('0'+mn%10), ':',
		byte('0'+ss/10), byte('0'+ss%10), ' ',
		'G', 'M', 'T')
}

const (
	days   = "SunMonTueWedThuFriSat"
	months = "JanFebMarAprMayJunJulAugSepOctNovDec"
)

func reverseStr(str string) string {
	a := []byte(str)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func reverseStrV2(str string) string {
	a := []byte(str)
	var b []byte
	for i := len(str) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return string(b)
}

func main() {
	// t := time.Now()
	// yy, mm, dd := t.Date()
	// hh, mn, ss := t.Clock()
	// day := days[3*t.Weekday():]
	// mon := months[3*(mm-1)]
	// fmt.Println(yy, mm, dd)
	// fmt.Println(hh, mn, ss)
	// fmt.Println(day)
	// fmt.Println(mon)
	// fmt.Println(yy, dd)

	a := "yoko"
	b := reverseStrV2(a)
	fmt.Println(b)
}
