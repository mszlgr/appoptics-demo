package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"

    "github.com/go-redis/redis"
    "github.com/appoptics/appoptics-apm-go/v1/ao"
)

func hello(w http.ResponseWriter, req *http.Request) {
  t := ao.TraceFromContext(req.Context())
  fmt.Printf("%s\n", t.LoggableTraceID())
  fmt.Fprintf(w, "hello - from golang-ao\n")
}

func fail(w http.ResponseWriter, req *http.Request) {
  http.Error(w, "internal error", http.StatusInternalServerError)
}

func remote(w http.ResponseWriter, req *http.Request) {
  t := ao.TraceFromContext(req.Context())
  fmt.Printf("%s\n", t.LoggableTraceID())

  c := &http.Client{}
  interReq, _ := http.NewRequestWithContext(req.Context(), "GET", "http://node-ao:3000/", nil)

  l := ao.BeginHTTPClientSpan(req.Context(), interReq)
  defer l.End()

  resp, e := c.Do(interReq)
  if e != nil {
    fmt.Printf("%s\n", e.Error())
    fmt.Fprintf(w, e.Error())
    return
  }
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Fprintf(w, string(body))
}

func redis_handler(w http.ResponseWriter, req *http.Request) {
  span, _ := ao.BeginSpan(req.Context(), "redis",
		"Spec", "cache",
		"KVOp", strings.ToLower("info"),
  )
  defer span.End()
  rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
    })
  info := rdb.Info("cpu")
  fmt.Fprintf(w, "%v\n", info)
}


func main() {
  http.HandleFunc("/", ao.HTTPHandler(hello))
  http.HandleFunc("/fail", ao.HTTPHandler(fail))
  http.HandleFunc("/remote", ao.HTTPHandler(remote))
  http.HandleFunc("/redis", ao.HTTPHandler(redis_handler))

  fmt.Print("Starting on port 8000")
  http.ListenAndServe(":8000", nil)
}
