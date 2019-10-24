package endpoint

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// go test ./... -v -cover -race
var router *mux.Router

func spyOnRequest(verb string, url string, strPayload string) (int, string) {
	responseRecorder := httptest.NewRecorder()
	bytesPayload := strings.NewReader(strPayload)
	request, e := http.NewRequest(verb, url, bytesPayload)
	if e != nil {
		panic("Oh Bum")
	}
	router.ServeHTTP(responseRecorder, request)
	return responseRecorder.Code, responseRecorder.Body.String()
}

func Example_endpointusage() {
	router = InitRouter()
	verb, url, payload := "GET", "/workshops", ""
	code, body := spyOnRequest(verb, url, payload)
	fmt.Printf("%s %s %s -> Code: %s (%d) Body %s\n", verb, url, payload, http.StatusText(code), code, body)

	// Output:
	// GET /workshops  -> Code: Accepted (202) Body
}

func Example_endpointerrors() {
	router = InitRouter()
	code, body := spyOnRequest("GET", "/workshops", "")
	fmt.Printf("GET /workshops -> Code: %s (%d) Body %s\n", http.StatusText(code), code, body)

	// Output:
	// GET /workshops  -> HTTP Status: OK(200), Body: ""
}

func Test_endpointconcurrency(t *testing.T) {
	t.Error("ToDo")
}

func Benchmark_endpoint(b *testing.B) {
	b.Error("ToDo")
}
