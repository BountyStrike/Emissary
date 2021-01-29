package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

func TestMainApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	webhook := "webhook:=" + ts.URL + "§data:={'chat_id': 'xxxx'}"
	old := os.Args
	os.Args = []string{"main", "-in", webhook, "-m", "hacker"}
	main()
	// Reset so next tests will work
	os.Args = old

	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", 200, res.StatusCode)
	}
}

func TestMainWithoutWebhook(t *testing.T) {

	// This is a hack in order to test the application without
	// failing the test when e.g. os.Exit(1) is executed.
	// By doing this I have complete control over the exit status code.
	if os.Getenv("EXEC") != "" {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		webhook := ts.URL + "§data:={'chat_id': 'xxxx'}"
		os.Args = []string{"main", "-in", webhook, "-m", "hacker"}
		main()

		err = res.Body.Close()
		if err != nil {
			t.Fatal(err)
		}

		if res.Request.Method == "" {
			t.Errorf("Expected no request to server, got a %s request", res.Request.Method)
		}

		panic("Should never reach this")
	}

	// Run the test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestMainWithoutWebhook")
	cmd.Env = append(os.Environ(), "EXEC=TRUE")
	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	e, ok := err.(*exec.ExitError)

	if ok != true {
		t.Errorf("Was not true, is %t", ok)
	}

	expectedErrorString := "exit status 1"
	if e.Error() != expectedErrorString {
		t.Errorf("Was not expectedSTring, got: %s", e.Error())
	}

}
