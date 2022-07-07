package lab

import (
	"fmt"
	"net/http"
)

func WebAssemblyLab() {

	println("Web Assembly Lab")

	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("./lab/webassembly_lab/.")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
