package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//One function in GO may return two values
	//The _ is a blank identifier, meaning it will not be used
	res, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(res.Body) //Uppercase .Body because is from another package (is public/exported func)
	res.Body.Close()
	fmt.Printf("%s", body)
}
