package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	fmt.Println(args)
	readFileContent(args[1])

	// return

	// email := email{
	// 	mail: "Hi there",
	// }

	// //showNotification(email)

	// sms := sms{
	// 	text: "Hello there",
	// }

	// //showNotification(sms)

	// disc := discord{
	// 	message: "Hi there, Im a discord bot",
	// }

	// notifierSlc := []notifier{email, sms, disc}

	// for _, n := range notifierSlc {
	// 	n.notify()
	// }

	// sq := square{
	// 	sideLength: 1.5,
	// }

	// tr := triangle{
	// 	base:   5.5,
	// 	height: 6.5,
	// }

	// shapeSlc := []shape{sq, tr}

	// for _, s := range shapeSlc {
	// 	fmt.Println(s.area())
	// }

}

func readFileContent(name string) {

	bs, err := os.Open(name)

	if err != nil {
		fmt.Errorf("Error while reading the file content")
	}

	data, err := io.Copy(os.Stdout, bs)

	fmt.Println(data)
}
