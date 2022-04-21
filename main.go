package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostPort := strings.TrimPrefix(r.URL.Path, "/prometheus/targets/add/")
	hostColonPort := strings.Replace(hostPort, "_", ":", -1)
	err := os.WriteFile("/u01/data/sd_configs/redis/"+hostPort+".json", []byte("[{\"labels\":{\"job\":\"redis\",\"env\":\"redis\"},\"targets\":[\""+hostColonPort+"\"]}]"), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Fprintf(w, hostColonPort)
}

func main() {
	http.HandleFunc("/prometheus/targets/add/", handler)
	http.ListenAndServe(":9091", nil)
}
