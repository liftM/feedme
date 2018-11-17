package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/liftM/feedme/hack/caviar"
)

func main() {
	user := flag.String("caviar_username", "", "Caviar username")
	pass := flag.String("caviar_password", "", "Caviar password")
	flag.Parse()

	log.Printf("user: %#v", *user)
	log.Printf("pass: %#v", *pass)

	session, err := caviar.New()
	if err != nil {
		panic(err)
	}

	err = session.SignIn(*user, *pass)
	if err != nil {
		panic(err)
	}

	res, err := session.Get("/san-francisco")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
