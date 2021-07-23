package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://localhost:9090/api/videos"
  method := "POST"

  payload := strings.NewReader(`

{
    "title": "Golang  Go Gin Framework Crash Course 02 | Middlewares 101: Authorization, Logging and Debugging",
    "description": "In this video we are going to start working with Middlewares providing Authorization, Logging and Debugging capabilities to our API using Golang's Gin HTTP Framework.",
    "url": "https://www.youtube.com/embed/sDJLQMZzzM4?list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w",
    "author" : {
        "firstname" : "pragmatic",
        "lastname" : "reviews",
        "email" : "invokerx@163.com",
        "age" : 30
    }
}

// {
//     "title": "Golang Gin Framework Crash Course 03 | Data Binding and Validation",
//     "description": "In this video we are going to take a look at Data Binding and Validation using Golang's Gin HTTP Framework.",
//     "url": "https://youtu.be/hP56ZzQt7Ag",
//     "author" : {
//         "firstname" : "pragmatic",
//         "lastname" : "reviews",
//         "email" : "invokerx@163.com",
//         "age" : 30
//     }
// }
// {
//     "title": "Golang / Go Gin Framework Crash Course 04 | HTML, Templates and Multi-Route Grouping",
//     "description": "In this video we are going to take a look at HTML, Templates and Multi-Route Grouping using Golang's Gin HTTP Framework.",
//     "url": "https://youtu.be/OUimJ0y8lzI?list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w",
//     "author" : {
//         "firstname" : "pragmatic",
//         "lastname" : "reviews",
//         "email" : "invokerx@163.com",
//         "age" : 30
//     }
// }

`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}