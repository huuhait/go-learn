package main

import (
	"os"
	"time"

	"github.com/huuhait/go-learn/config"
	"github.com/huuhait/go-learn/jobs"
)

type Job interface {
	Process()
}

func GetJob(job string) Job {
	switch job {
	case "fetch_price":
		return jobs.NewFetchPrice()
	case "sleep":
		return jobs.NewSleep()
	default:
		return nil
	}
}

func main() {
	config.InitConfig()

	for index, job_name := range os.Args {
		if index == 0 {
			continue
		}

		job := GetJob(job_name)

		go job.Process() // => other
	}

	for {
		time.Sleep(5 * time.Second)
	}
}
