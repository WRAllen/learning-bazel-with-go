package pkg

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func helloHandler(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request", zap.String("method", r.Method), zap.String("url", r.URL.String()))
		fmt.Fprintln(w, "Hello, World! 🌍")
	}
}

func Run() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Cannot initialize zap logger: %v", err))
	}
	defer logger.Sync()

	port := "8080"
	logger.Info("Starting server", zap.String("port", port))

	http.HandleFunc("/", helloHandler(logger))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
