package main

import (
    "net/http"
    "sync"
    "time"
)

// RequestStats is used to track request statistics for each IP address.
type RequestStats struct {
    SuccessCount int
    FailCount    int
    LastRequest  time.Time
}

// Global variables for storing request data and synchronization.
var (
    requestMutex sync.Mutex
    requestStats = make(map[string]*RequestStats)
)

// handleSendRawTransactionOptional handles the sendRawTransactionOptional requests.
func handleSendRawTransactionOptional(w http.ResponseWriter, r *http.Request) {
    ip := r.RemoteAddr

    // Check if requests from this IP should be limited.
    if shouldLimitRequest(ip) {
        http.Error(w, "Request rate limited", http.StatusTooManyRequests)
        return
    }

    // Mock transaction processing logic (replace with actual logic).
    success := processMockTransaction()

    // Update the stats for this IP based on the transaction result.
    updateStats(ip, success)

    // Send an OK response.
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Transaction processed"))
}

// processMockTransaction simulates a transaction processing logic.
func processMockTransaction() bool {
    // Include the actual transaction processing logic here.
    // Assuming all requests are processed successfully for this example.
    return true 
}

// updateStats updates the request statistics for the given IP.
func updateStats(ip string, success bool) {
    requestMutex.Lock()
    defer requestMutex.Unlock()

    stats, exists := requestStats[ip]
    if !exists {
        stats = &RequestStats{}
        requestStats[ip] = stats
    }

    if success {
        stats.SuccessCount++
    } else {
        stats.FailCount++
    }
    stats.LastRequest = time.Now()
}

// shouldLimitRequest decides whether requests from a given IP should be limited.
func shouldLimitRequest(ip string) bool {
    requestMutex.Lock()
    defer requestMutex.Unlock()

    stats, exists := requestStats[ip]
    if !exists || time.Since(stats.LastRequest) > 1*time.Minute {
        // Do not limit new IPs or IPs with no requests in the last minute.
        return false
    }

    total := stats.SuccessCount + stats.FailCount
    if total == 0 {
        return false
    }
    successRate := float64(stats.SuccessCount) / float64(total)
    return successRate < 0.5
}

// main starts the HTTP server.
func main() {
    http.HandleFunc("/sendRawTransactionOptional", handleSendRawTransactionOptional)
    http.ListenAndServe(":8080", nil)
}
