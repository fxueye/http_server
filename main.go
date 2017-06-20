package main

//281431280
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port int
}

func main() {
	config := Config{}
	bytes, err := ioutil.ReadFile("./config.json")
	err = json.Unmarshal(bytes, &config)
	wd, err := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(wd)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server start on:%d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
