package main

import (
    "net/url"
    "net/http"
    "encoding/json"
)

type Webservice struct { }

func (WhoisAPI Webservice) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {

       // get IP
       IP := values.Get("IP")
       
       // get redis connection
       client := GetConnection()
       defer client.Close()
       
       // get cache data
       data := client.Get(IP);
       
       // no cache 
       if len(data) == 0 {
         query := GetQuery(IP)
         cache, _ := json.Marshal(query)
         client.Set(IP, string(cache))
         return 200, query , GetHeaders()
       }

      // use cache
       var result = new(QueryReponse)       
       json.Unmarshal([]byte(data), &result)
       return 200, result, GetHeaders()
}
