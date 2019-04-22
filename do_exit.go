package main

import "log"

func doExit() {
	log.Println("exit, bye.")
	sqlDB.Close()
}
