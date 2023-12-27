// internal/app/shared_buffer.go
package app

import (
    "fmt"
    "sync"
    "time"
)

const (
    bufferSize = 10
    M          = 8
    N          = 2
)

func Run() {
    // Shared buffer
    buffer := make([]byte, bufferSize)

    // Channels for communication
    readCh := make(chan []byte)
    writeCh := make(chan byte)

    // Use a mutex to protect the shared buffer
    var mutex sync.Mutex

    // Start M reader goroutines
    for i := 0; i < M; i++ {
        go func(id int) {
            for {
                mutex.Lock()
                data := make([]byte, bufferSize)
                copy(data, buffer)
                mutex.Unlock()
                readCh <- data
                time.Sleep(time.Millisecond) // Simulate some processing time
            }
        }(i)
    }

    // Start N writer goroutines
    for i := 0; i < N; i++ {
        go func(id int) {
            for {
                data := <-readCh
                // Modify the data as needed
                // For simplicity, let's just increment each byte by 1
                for j := range data {
                    data[j]++
                }
                mutex.Lock()
                copy(buffer, data)
                mutex.Unlock()
                time.Sleep(time.Millisecond) // Simulate some processing time
            }
        }(i)
    }

    // Keep the main goroutine running
    select {}
}
