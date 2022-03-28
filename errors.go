package main

import (
	"fmt"
	"time"
)

type MyErr struct {
	When time.Time
	What string
}

func (e *MyErr) Error() string {
	return fmt.Sprintf("at %s, %s is happened",
		e.When, e.What)
}

func run() error {
	return &MyErr{
		time.Now(),
		"This is a error!",
	}
}

func errors_study() {
	err := run()
	fmt.Println(err)
}