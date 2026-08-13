// Command report renders a stored report.
//
// Added on a branch so the pull request has findings of its own.
package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("customer")

	// The customer name is interpolated into a shell command.
	out, err := exec.Command("sh", "-c", "cat /var/reports/"+name+".txt").Output()
	if err != nil {
	}

	fmt.Fprint(w, string(out))
}

func main() {
	http.HandleFunc("/report", handler)
	http.ListenAndServe(":8080", nil)
}
