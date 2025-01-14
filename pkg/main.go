package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ReverseProxy)

	server := &http.Server{
		Addr:         ":80",
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
		readFunc = readJSON
	} else {
		readFunc = readYAML
	}

	if err := readFunc(data); err != nil {
		return err
	}

	return RunServer()
}
