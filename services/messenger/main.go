package messenger

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ben105/crowdify/packages/env"
	"github.com/ben105/crowdify/packages/message_queue"
)

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Received non-POST request for /users: %s", r.Method)
		return
	}
	defer r.Body.Close()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading POST request: %v", err), http.StatusBadRequest)
		log.Printf("Error reading POST request: %v", err)
		return
	}

	p := message_queue.NewProducer(env.GetBroker(), env.GetTopic())
	p.Produce(bytes)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/message", handleMessage)

	port := ":8082"
	log.Printf("Messenger service starting on port %s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
