package gorevproxy

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shanedell/goutils/env"
)

func Healthcheck(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]any{
		"status": 200,
	}); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode healthcheck json: %s", err), http.StatusBadRequest)
		return
	}
}

func RunServer(proxyFunc func(w http.ResponseWriter, r *http.Request)) error {
	port := env.GetString("PROXY_PORT", "80")

	mux := http.NewServeMux()
	mux.HandleFunc("/", proxyFunc)
	mux.HandleFunc("/healthcheck", Healthcheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Reverse proxy running on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("server failed: %v", err)
	}

	return nil
}

func Run(args *ProxyArgs) error {
	file, err := os.Open(args.ConfigFile)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if args.ReadJSON {
		readFunc = ReadJSON
	} else {
		readFunc = ReadYAML
	}

	if err := readFunc(data); err != nil {
		return err
	}

	return RunServer(ReverseProxy)
}
