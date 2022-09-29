package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func eatMemory(memSlice []int64) {

	for {
		memSlice = append(memSlice, make([]int64, 125_000)...)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.Printf(
			"\nTotalAlloc = %v\n Sys = %v\n HeapAlloc=%v\n",
			m.TotalAlloc/1000/1000,
			m.Sys/1000/1000,
			m.HeapAlloc/1000/1000,
		)
		time.Sleep(time.Second)
	}
}

func main() {

	memSlice := []int64{}

	// m := http.NewServeMux()

	go eatMemory(memSlice)

	// m.HandleFunc("…", pprof.Index)

	// m.HandleFunc("/getMemory", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, Otter!"))
	// 	var m runtime.MemStats
	// 	runtime.ReadMemStats(&m)
	// 	log.Printf("TotalAlloc = %v\nSys = %v\n", m.TotalAlloc/1024/1024, m.Sys/1024/1024)
	// 	w.Write([]byte(fmt.Sprintf("Memory Alloc: %v", m.TotalAlloc/1024/1024)))
	// })

	// go func() {
	log.Println(http.ListenAndServe(":6060", http.DefaultServeMux))
	// }()

	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: m,
	// }

	// log.Fatal(s.ListenAndServe())
}
