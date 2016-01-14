// Credits: https://gist.github.com/ijt/950790
package main

import (
    "fmt"
    "http"
    "io/ioutil"
    "os"
    )

func main() {
    response, err := http.Get("http://golang.org/")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}
