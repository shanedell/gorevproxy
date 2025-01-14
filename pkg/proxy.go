package pkg

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
)

func ReverseProxy(w http.ResponseWriter, r *http.Request) {
	serverNames := []string{}
	serverIndexNames := map[string]int{}

	for i, server := range config.Servers {
		serverNames = append(serverNames, server.Name)
		serverIndexNames[server.Name] = i
	}

	if slices.Contains(serverNames, r.Host) {
		server := config.Servers[serverIndexNames[r.Host]]

		path := r.URL.Path

		for _, location := range server.Locations {
			// if paths match or the path is not specified or is "/" allowed forwarding
			if path == location.Path || slices.Contains([]string{"", "/"}, location.Path) {
				target, err := url.Parse(fmt.Sprintf("%s://%s:%s%s", location.To.Schema, location.To.Host, location.To.Port, location.Path))
				if err != nil {
					http.Error(w, "Error parsing target URL", http.StatusInternalServerError)
					return
				}

				proxy := httputil.NewSingleHostReverseProxy(target)

				proxy.ServeHTTP(w, r)
				return
			}
		}
	}

	http.Error(w, "Host or path not found", http.StatusNotFound)
}
