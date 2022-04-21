package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func add(w http.ResponseWriter, r *http.Request) {
	hostPort := strings.TrimPrefix(r.URL.Path, "/prometheus/targets/add/")
	hostColonPort := strings.Replace(hostPort, "_", ":", -1)
	err := os.WriteFile("/u01/data/sd_configs/redis/"+hostPort+".json", []byte("[{\"labels\":{\"job\":\"redis\",\"env\":\"redis\"},\"targets\":[\""+hostColonPort+"\"]}]"), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Fprintf(w, hostPort)
}

func remove(w http.ResponseWriter, r *http.Request) {
	hostPort := strings.TrimPrefix(r.URL.Path, "/prometheus/targets/remove/")
	err := os.Remove("/u01/data/sd_configs/redis/" + hostPort + ".json")
	if err != nil {
		fmt.Printf("Unable to remove file: %v", err)
	}
	fmt.Fprintf(w, hostPort)
}

func main() {
	http.HandleFunc("/prometheus/targets/add/", add)
	http.HandleFunc("/prometheus/targets/remove/", add)
	http.ListenAndServe(":9091", nil)
}
