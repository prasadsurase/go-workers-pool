package main

import (
	"fmt"
	"time"
)

// EmailProcessor holds the email address
type EmailProcessor struct {
	email string
}

// Process method mimicks sending an email.
func (ep *EmailProcessor) Process() {
	fmt.Printf("Sending email to %s\n", ep.email)
	time.Sleep(1 * time.Second)
}
