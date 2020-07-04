package tools

import "log"

func FatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Database has been connected")
	}
}
func PrintlnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
func PanicErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
