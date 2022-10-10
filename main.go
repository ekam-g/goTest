package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	for {
		var err *bool = null()
		if err != nil {
			switch *err {
			case true:
				{
					log.Println("cool")
				}
			case false:
				{
					log.Println("bad one")
				}
			}
		} else {
			log.Println("value is null")
		}
		time.Sleep(2 * time.Second)
	}
}

func null() *bool {
	val := true
	v := rand.Intn(3)
	if v == 2 {
		return &val
	} else if v == 1 {
		val = false
		return &val
	} else {
		log.Println(v)
		return nil
	}
}
