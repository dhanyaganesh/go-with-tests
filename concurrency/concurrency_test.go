package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat:://sadfjlaksjdfkjnads..sdfkjhs" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsystem.com",
		"waat:://sadfjlaksjdfkjnads..sdfkjhs",
	}

	want := map[string]bool{
		"http://google.com":					true,
		"http://blog.gypsystem.com":			true,
		"waat:://sadfjlaksjdfkjnads..sdfkjhs":	false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}