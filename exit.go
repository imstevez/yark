package main

import "log"

func exit() {
	log.Println("exit, bye.")
	sqlDB.Close()
}
