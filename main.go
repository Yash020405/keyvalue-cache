// main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
)

// Cache stores key-value pairs in memory.
type Cache struct {
    store map[string]string
    mutex sync.RWMutex
}

// NewCache creates a new Cache.
func NewCache() *Cache {
    return &Cache{
        store: make(map[string]string),
    }
}

// Put adds or updates a key-value pair.
func (c *Cache) Put(key, value string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.store[key] = value
}

// Get retrieves the value associated with the key.
func (c *Cache) Get(key string) (string, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    val, ok := c.store[key]
    return val, ok
}

// PutRequest represents the JSON payload for /put.
type PutRequest struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

// Response is used to send back JSON responses.
type Response struct {
    Status  string `json:"status"`
    Message string `json:"message,omitempty"`
    Key     string `json:"key,omitempty"`
    Value   string `json:"value,omitempty"`
}

func main() {
    cache := NewCache()

    // POST /put endpoint
    http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        var req PutRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        // Validate key and value lengths (max 256 ASCII characters)
        if len(req.Key) == 0 || len(req.Key) > 256 || len(req.Value) > 256 {
            resp := Response{Status: "ERROR", Message: "Key or value length violation"}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(resp)
            return
        }
        cache.Put(req.Key, req.Value)
        resp := Response{Status: "OK", Message: "Key inserted/updated successfully."}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
    })

    // GET /get endpoint
    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        key := r.URL.Query().Get("key")
        if len(key) == 0 {
            resp := Response{Status: "ERROR", Message: "Key parameter missing"}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(resp)
            return
        }
        value, found := cache.Get(key)
        if !found {
            resp := Response{Status: "ERROR", Message: "Key not found."}
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(resp)
            return
        }
        resp := Response{Status: "OK", Key: key, Value: value}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
    })

    log.Println("Starting server on port 7171")
    if err := http.ListenAndServe(":7171", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
