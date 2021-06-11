package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockCheckWebsites(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockCheckWebsites, websites)
	assert.Equal(t, got, want)
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i :=0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i:=0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}