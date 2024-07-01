package main

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

var (
	Version = "v1"
	Date    = ""
)

func main() {
	fmt.Println("app was built on ", Date)
	// a := api.New(Version)
	// a.Start(a.Engine)

	// sanitize input (email)
	email := "test.user@gmail.com"
	fmt.Println(isEmailValid(email))
	if em, err := mail.ParseAddress(email); err != nil {
		fmt.Println("the email is invalid")
	} else {
		fmt.Println(em, "the provided email address")
	}

	// sanitize html files
	htmlInput := `<div><script>console.log("script inside html file")</script></div>`
	policy := bluemonday.UGCPolicy()
	sanitizeInput := policy.Sanitize(htmlInput)
	fmt.Println("the sanitized html input file", sanitizeInput)
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
