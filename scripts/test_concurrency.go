package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func patchPost(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	url := fmt.Sprintf("http://localhost:3000/v1/posts/%d", id)

	jsonData := []byte(`{"content":"pathched content ohai guys"}`)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("request error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go patchPost(&wg, 4)
	}

	wg.Wait()
}
