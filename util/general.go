package util

import "log"

//Recover - catch panic
func Recover() {
	if r := recover(); r != nil {
		log.Println("PANIC:", r)
	}
}
