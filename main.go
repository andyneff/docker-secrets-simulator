package main

import (
    "os"
    "strings"
    "io/ioutil"
    "fmt"
    "sort"
    "strconv"
)

func main() {
  // This MUST remain in alphabetical order for contains to work
  ignores := []string{ "GOLANG_VERSION", "GOPATH", "HOME", "HOSTNAME", "PATH",
                       "TERM", "no_proxy" }

  mode := os.FileMode(0644)

  if len(os.Args) > 1 {
    mode_int, _ := strconv.ParseInt(os.Args[1], 8, 64)
    mode = os.FileMode(mode_int)
  }

  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    if !contains(ignores, pair[0]) {
      contents := []byte(pair[1]);
      err := ioutil.WriteFile("/run/secrets/" + pair[0], contents, mode)
      check(err)
    }
  }
}

func contains(list []string, value string) bool {
	i := sort.SearchStrings(list, value)
	return i < len(list) && list[i] == value
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
