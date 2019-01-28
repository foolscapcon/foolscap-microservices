package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"io/ioutil"
	"net/http"
	"./tito"

)
func check(e error) {
	if e != nil {
		panic(e)
	}
}


const (
	secret_file = ".tito_secret.yaml"
	path = "/"
)

func main() {
	dat, err := ioutil.ReadFile(secret_file)
	check(err)
	m := make(map[string]string)

	err = yaml.Unmarshal([]byte(dat), &m)
	check(err)
	
	fmt.Printf("--- m:\n%v\n\n", m)	
	hook, _ := tito.New(tito.Options.Secret(m["security_token"]))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, tito.RegistrationCompleted)
		if err != nil {
			if err == tito.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case tito.RegistrationCompletedPayload:
			release := payload.(tito.RegistrationCompletedPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		}
	})
	http.ListenAndServe(":80", nil)
}
