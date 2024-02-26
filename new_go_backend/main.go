package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define routes
    http.HandleFunc("/nocache", NoCacheHandler)
    http.HandleFunc("/publiccache", PublicCacheHandler)
    http.HandleFunc("/privatecache", PrivateCacheHandler)

    // Start server
    fmt.Println("Server is listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

// Handler for "/nocache" route with Cache-Control: no-store header
func NoCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-store")
    fmt.Fprintf(w, "Response with Cache-Control: no-store")
}

// Handler for "/publiccache" route with Cache-Control: public header
func PublicCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "public, max-age=3600")
    fmt.Fprintf(w, "Response with Cache-Control: public, max-age=3600")
}

// Handler for "/privatecache" route with Cache-Control: private header
func PrivateCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "private, max-age=3600")
    fmt.Fprintf(w, "Response with Cache-Control: private, max-age=3600")
}
