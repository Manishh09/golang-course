package main

import "fmt"

type notifier interface {
	notify() error
}

type email struct {
	mail string
}

type sms struct {
	text string
}

type discord struct {
	message string
}

func (e email) notify() error {
	fmt.Println(e.mail)
	return nil
}
func (s sms) notify() error {
	fmt.Println(s.text)
	return nil
}
func (d discord) notify() error {
	fmt.Println(d.message)
	return nil
}
