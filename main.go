package main

import(
  "net/http"
  "log"
  "io/ioutil"
//  "encoding/json"
//  "fmt"
//  "reflect"
)

type pr_label struct {
  Name string              `json:"name"`
  Number int
}

func main() {
 println("hello!")
 url := "https://api.github.com/repos/maruyama-ke/hello/issues"
 resp, err := http.Get(url)
 if err != nil {
   log.Fatal(err)
 }
 defer resp.Body.Close()
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
   log.Fatal(err)
 }
 var issue_list = string(body)
 println(issue_list)
}
