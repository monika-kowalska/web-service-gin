package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFindCampains(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/campaigns", ts.URL))

	fmt.Println(resp)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestFindCampain(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/campaigns/1", ts.URL))

	fmt.Println(resp)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
