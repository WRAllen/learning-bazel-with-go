package pkg

import (
	"fmt"
	"net/http"

	"github.com/wrallen/sampleBazel2/pkg/log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! 🌍")
}

func Run() {
	http.HandleFunc("/", helloHandler)

	Port := "8080"
	log.LogRecord(Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
	if err != nil {
		panic(err)
	}
}