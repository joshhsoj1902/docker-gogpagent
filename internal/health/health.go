package health

import (
    "net/http"
) 

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
  }