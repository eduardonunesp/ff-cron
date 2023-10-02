package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

var (
	appURL    string // APP_URL
	cronToken string // CRON_TOKEN
	cronSched string // CRON_SCHED
	quickRun  bool   // QUICK_RUN
)

func init() {
	if v, ok := os.LookupEnv("APP_URL"); ok {
		appURL = v
	}

	if v, ok := os.LookupEnv("CRON_TOKEN"); ok {
		cronToken = v
	}

	if v, ok := os.LookupEnv("CRON_SCHED"); ok {
		cronSched = v
	}

	if v, ok := os.LookupEnv("QUICK_RUN"); ok {
		quickRun = v == "true"
	}

	if appURL == "" || cronToken == "" || cronSched == "" {
		panic("APP_URL, CRON_TOKEN and CRON_SCHED are required")
	}
}

func cronJob() {
	url := fmt.Sprintf("%s/api/v1/cron/%s", appURL, cronToken)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintln(os.Stdout, "HTTP response status:", resp.Status)
}

func main() {
	if quickRun {
		cronJob()
		return
	}

	c := cron.New()
	c.AddFunc(cronSched, cronJob)
	c.Start()

	select {}
}
