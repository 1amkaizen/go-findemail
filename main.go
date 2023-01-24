package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func init() {
	flag.Usage = func() {
		h := []string{
			"-d  for domain",
			"-fn for FirstName",
			"-ln for LastName\n",
			"Example   findemail -d gmail.com -fn john -ln doe\n",
		}
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}

}

type User struct {
	Email string `json:"email"`
}

//anymailfinder api
//h8mrDtt5x5Nkoryw9v8eH9kF

func skrapp() {
	var d string
	flag.StringVar(&d, "d", "Domain", "Masukkan Domain")
	var fn string
	flag.StringVar(&fn, "fn", "FirstName", "Masukkan FirstName")
	var ln string
	flag.StringVar(&ln, "ln", "LastName", "Masukkan LastName")

	flag.Parse()

	url := fmt.Sprintf("https://api.skrapp.io/api/v2/find?firstName=%s&lastName=%s&domain=%s", fn, ln, d)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("X-Access-Key", "103144543q5B67JkdTP2OYRhaRendXPpboO9QVHNMM")

	r, _ := client.Do(req)
	defer r.Body.Close()
	io.Copy(os.Stdout, r.Body)
	fmt.Println("\n")

}

func main() {
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	skrapp()
}
