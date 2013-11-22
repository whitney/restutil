package restutil

import (
  "fmt"
  "encoding/json"
  "net/http"
)

func JsonErr(w http.ResponseWriter, msg string, code int) {
  w.Header().Set("Content-Type", "application/json") 
  w.WriteHeader(code)
  fmt.Fprintln(w, fmt.Sprintf("{\"msg\": \"%s\"}", msg))
}

func JsonSucc(w http.ResponseWriter, json string, code int) {
  w.Header().Set("Content-Type", "application/json") 
  w.WriteHeader(code)
  fmt.Fprintln(w, json)
}

func JsonWrapStr(s string) string {
  m := make(map[string]string)
  m["msg"] = s
  js, err := JsonWrapMap(m)
  if err != nil {
    panic(err)
  }
  return js
}

func JsonWrapMap(m map[string]string) (s string, err error) {
  json, err := json.Marshal(m)
  if err != nil {
    return s, err
  }
  return string(json),nil
}
