package main

import (
	"time"

	"github.com/laidingqing/feichong/helpers"
)

func main() {
	log := helpers.NewLogger()
	year, month, _ := time.Now().Date()
	var orderMonth = int(month)
	var orderYear = int(year)
	var m = 0
	var next = false
	for i := 0; i < 12; i++ {
		log.Log("end", orderMonth+m, "orderMonth", orderMonth, "m", m, "orderYear", orderYear)
		if orderMonth < 12 {
			orderMonth++
			m = i
		} else {
			orderMonth = 1
			if next == false {
				next = true
			}
			m++
		}
		if next {
			orderYear = year + 1
		}
	}

}
