package learngo

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)



type Event struct {
	ID string 
	Event string 
	Data string 
	Retry int
}



func (e Event) format() string {
	var s string;
	if e.ID != "" {
		s += fmt.Sprintf("id: %s\n", e.ID)
	}
	if e.Event != "" {
		s += fmt.Sprintf("event: %s\n", e.Event)
	}
	if e.Retry > 0{
		s += fmt.Sprintf("retry: %d\n", e.Retry)
	}

	s += fmt.Sprintf("data: %s\n\n", e.Data)
	return s 
}



// manages client connections and fans out events to all subscribers
type Broker struct {
	clients map[chan Event]bool
	mu sync.RWMutex // for thread-safety 
}

func NewBroker() *Broker {
	return &Broker {
		clients: make(map[chan Event]bool),
	}
}


func (b* Broker) Subscribe() chan Event {
	ch := make(chan Event, 10)
	b.mu.Lock()
	b.clients[ch] = true
	b.mu.Unlock()
	return ch
}


func (b *Broker) Unsubscribe(ch chan Event){
	b.mu.Lock()
	delete(b.clients, ch)
	close(ch)
	b.mu.Unlock()
}


func (b *Broker) Publish  (event Event){
	b.mu.RLock()
	defer b.mu.RUnlock()
	for ch := range b.clients {
		select {
		case ch <- event:
		default: // drop the event rather than blocking the entire publish loop
		}
	}
}


func handleSSE(broker *Broker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming not supported", http.StatusInternalServerError)
		}
		ch := broker.Subscribe()
		defer broker.Unsubscribe(ch)
		for {
			select {
			case event := <- ch:
				fmt.Fprintf(w, event.format())
				flusher.Flush()
			case <-r.Context().Done():
				return
			}
		}
	}
}




var symbols = []string{"APPL", "GOOG", "MSFT", "AMZN", "TSLA"};

func startTicker(broker *Broker){
	counter := 0
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			counter++;
			symbol := symbols[rand.Intn(len(symbols))]
			price := 100 + rand.Float64() * 200
			broker.Publish(Event{
				ID: fmt.Sprintf("%d", counter),
				Event: "stock-update",
				Data: fmt.Sprintf(`{"symbol":"%s","price":"%.2f","time":"%s"}`, symbol, price, time.Now().Format(time.RFC3339)),
			})
		}
	}()
}
 

func main() {
	broker := NewBroker()
	startTicker(broker)


	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/events", handleSSE(broker))


	log.Println("Server started on :8080")
	log.Println("Dashboard: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}