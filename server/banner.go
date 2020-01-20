package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func banner() {
	log.Printf("\n" + "\n" +
		"  _______   ______          .___  ___.   ______   .__   __.  __  .___________.  ______   .______      \n" +
		" /  _____| /  __  \\         |   \\/   |  /  __  \\  |  \\ |  | |  | |           | /  __  \\  |   _  \\     \n" +
		"|  |  __  |  |  |  |  ______|  \\  /  | |  |  |  | |   \\|  | |  | `---|  |----`|  |  |  | |  |_)  |    \n" +
		"|  | |_ | |  |  |  | |______|  |\\/|  | |  |  |  | |  . `  | |  |     |  |     |  |  |  | |      /     \n" +
		"|  |__| | |  `--'  |        |  |  |  | |  `--'  | |  |\\   | |  |     |  |     |  `--'  | |  |\\  \\----.\n" +
		" \\______|  \\______/         |__|  |__|  \\______/  |__| \\__| |__|     |__|      \\______/  | _| `._____|\n" +
		"                                                                                                      ")
	fmt.Println("\n")
	time.Sleep(time.Millisecond * 2)
}
