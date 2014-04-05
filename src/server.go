package main

import (
    "github.com/dougblack/sleepy"
)

func main() {
    apiResource := new(Webservice)
    api := sleepy.NewAPI()
    api.AddResource(apiResource, "/WhoisAPI")
    api.Start("127.0.0.1", 80)
}
