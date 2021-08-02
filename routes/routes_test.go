package routes

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal("Request to 'GET / ' filed!")
	}
	fmt.Println(req)

}
