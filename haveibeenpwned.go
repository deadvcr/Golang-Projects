/*
This simple application uses the haveibeenpwned.com API to check if an email has been 'pwned'
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const inputdelimiter = '\n'

func main() {
	reader := bufio.NewReader(os.Stdin)
	var email string
	fmt.Println("Enter an email to check if you've been pwned")
	email, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}

	email = strings.TrimSpace(email)
	checkEmail(email)
}

func checkEmail(email string) {
	req, err := http.Get("https://haveibeenpwned.com/api/v2/breachedaccount/" + email)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	switch req.StatusCode {
	case 200:
		fmt.Println("You have been breached!")
		break
	case 404:
		fmt.Println("You're safe! You have not been breached.")
		break
	default:
		fmt.Println("Something went wrong.")
		return
	}

}
