package main

import (
	"flag"
	"log"

	"github.com/kr/pretty"

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

	// err = session.SignIn(*user, *pass)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := session.GetJSON("/san-francisco")
	// if err != nil {
	// 	panic(err)
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	err = session.SetLocation(caviar.Address{
		City:          "San Francisco",
		PostalCode:    "94110",
		State:         "California",
		StreetAddress: "1301 Valencia Street",
	})
	if err != nil {
		panic(err)
	}

	listing, err := session.Merchants()
	if err != nil {
		panic(err)
	}

	pretty.Println(listing)
}
