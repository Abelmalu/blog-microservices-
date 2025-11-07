package main

import (
	"log"
	"net/http"
)

const listenAddr = ":8000"
const imageFilename = "hello.jpg" // Assumes this file is in the same directory

// This handler specifically serves one file.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile automatically handles:
	// 1. Finding the file on the local filesystem.
	// 2. Reading the contents (the raw bytes).
	// 3. Setting the correct Content-Type header (e.g., image/png).
	// 4. Writing the file data to the response body.
	http.ServeFile(w, r, imageFilename)
}

func main() {
	// Register the handler to serve the image only when the client requests the root path.
	http.HandleFunc("/", imageHandler)

	log.Printf("✅ Single Image Server started on http://localhost%s", listenAddr)
	log.Printf("   Access image at: http://localhost%s/", listenAddr)

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
