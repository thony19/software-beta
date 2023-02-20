package main

// import (
//     "fmt"
//     "net/http"
//     "sync"
//     "time"

//     _ "net/http/pprof"
// )

// // Some function that does work
// func hardWork(wg *sync.WaitGroup) {
//     defer wg.Done()
//     fmt.Printf("Start: %v\n", time.Now())

//     // Memory
//     a := []string{}
//     for i := 0; i < 500000; i++ {
//         a = append(a, "aaaa")
//     }

//     // Blocking
//     time.Sleep(2 * time.Second)
//     fmt.Printf("End: %v\n", time.Now())
// }

// func main() {
//     var wg sync.WaitGroup

//     // Server for pprof
//     go func() {
//         fmt.Println(http.ListenAndServe("localhost:6060", nil))
//     }()
//     wg.Add(1) // pprof - so we won't exit prematurely
//     wg.Add(1) // for the hardWork
//     go hardWork(&wg)
//     wg.Wait()
// }
import ( "encoding/json"
    "math/rand" 
    "net/http" 
    _ "net/http/pprof" 
    "time" ) 
    
func main () { 
    // http.HandleFunc( "/log" , logHandler)
    // http.ListenAndServe( ":8080" , nil )


    if err := profiler.Start(profiler.Config{
        Service:        "gophersapi-service",
        ServiceVersion: "1.0",
        ProjectID:      "gophersapi", // optional on GCP
     }); err != nil {
        log.Fatalf("Cannot start the profiler: %v", err) 
     }
}
