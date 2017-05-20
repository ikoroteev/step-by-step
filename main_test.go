package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedGreeting := "URL.Path = \"/\""
	testingGreeting := strings.Trim(string(greeting), " \n")
	if testingGreeting != expectedGreeting {
		t.Fatalf("Wrong greeting '%s', expected '%s'", testingGreeting, expectedGreeting)
	}
}

func TestHandlerFewCases(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	var testCases = []string{"/abc", "/f", "///","/As"}

	for _, testCase := range testCases {
		res, err := http.Get(ts.URL + testCase)
		if err != nil {
			t.Fatal(err)
		}

		greeting, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			t.Fatal(err)
		}

		expectedGreeting := fmt.Sprintf("URL.Path = \"%s\"", testCase)
		testingGreeting := strings.Trim(string(greeting), " \n")
		if testingGreeting != expectedGreeting {
			t.Fatalf("Wrong greeting '%s', expected '%s'", testingGreeting, expectedGreeting)
		}
	}

}

// TODO: add one more test together
