package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "time"
)

func main() {
    // Define routes
    http.HandleFunc("/nocache", func(w http.ResponseWriter, r *http.Request) {
        NoCacheHandler(w, r)
    })
    http.HandleFunc("/publiccache", func(w http.ResponseWriter, r *http.Request) {
        PublicCacheHandler(w, r)
    })
    http.HandleFunc("/privatecache", func(w http.ResponseWriter, r *http.Request) {
        PrivateCacheHandler(w, r)
    })
    http.HandleFunc("/getresponse", func(w http.ResponseWriter, r *http.Request) {
        getresponseWithoutHeaders(w, r)
    })

    // Start server
    fmt.Println("Server is listening on port 8000...")
    http.ListenAndServe(":8000", nil)
}

// Handler for "/nocache" route with Cache-Control: no-store header
func NoCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-store")
    fmt.Fprintf(w, "Response with Cache-Control: no-store")
    sleepBeforeRespond()
}

// Handler for "/publiccache" route with Cache-Control: public header
func PublicCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "public, max-age=3600")
    fmt.Fprintf(w, "Response with Cache-Control: public, max-age=3600")
    sleepBeforeRespond()
}

// Handler for "/privatecache" route with Cache-Control: private header
func PrivateCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "private, max-age=3600")
    fmt.Fprintf(w, "Response with Cache-Control: private, max-age=3600")
    sleepBeforeRespond()
}

func getresponseWithoutHeaders(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Response without cache control headers")
    sleepBeforeRespond()
}

func sleepBeforeRespond() {
    sleepTimeStr := os.Getenv("SLEEP_TIME")
    sleepTime, err := strconv.Atoi(sleepTimeStr)
    if err != nil {
        sleepTime = 15
    }
    time.Sleep(time.Duration(sleepTime) * time.Second)
}
