package main

import (
	"fmt"
	"sync"
	"time"
)

type Bucket struct {
	tokens int
	lastTime time.Time
	mu sync.Mutex
} 

const (
	maxTokens  = 5
	refillInterval = time.Minute
)

func NewBucket() *Bucket {
	return &Bucket{
		tokens: maxTokens,
		lastTime: time.Now(),
	}
}

func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	currentTime := time.Now()
	elapsedTime := currentTime.Sub(b.lastTime)

	tokensToAdd := int(elapsedTime / refillInterval)
	if tokensToAdd > 0 {
		b.tokens = min(maxTokens, b.tokens + tokensToAdd)
		b.lastTime = currentTime
	}

	if b.tokens > 0 {
		b.tokens -= 1
		return true
	} else {
		return false
	}

}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main () {
	bucket := NewBucket()

	for i := 0; i < 10; i++ {
		if bucket.Allow() {
			fmt.Printf("Request %d allowed\n", i + 1)
		} else {
			fmt.Printf("Request %d denied\n", i + 1)
		}
		time.Sleep(10 * time.Second)
	}
}


