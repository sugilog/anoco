package agent

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func JobWorker(
	publisher chan Job,
	reporter chan Result,
	status chan Job,
	reply chan string,
) {
	jobappend := make(chan Job)
	jobremove := make(chan Job)
	jobs := make(map[int]Job)

	for {
		select {
		case j := <-publisher:
			go func(job Job) {
				var out bytes.Buffer
				cmd := exec.Command("bash", "-c", job.Command)

				cmd.Stdout = &out
				cmd.Stderr = &out
				jobappend <- job
				err := cmd.Run()

				result := Result{
					// FIXME
					Id:     job.Id,
					Output: out.String(),
				}

				if err != nil {
					result.Status = "ERROR"
					result.Error = fmt.Sprintf("%v", err)
					log.Fatal(err)
				} else {
					result.Status = "SUCCESS"
				}

				reporter <- result
				jobremove <- job
			}(j)
		case j := <-jobappend:
			jobs[j.Id] = j
		case j := <-jobremove:
			delete(jobs, j.Id)
		case j := <-checker:
			job := jobs[j.Id]

			if job.Id != 0 {
				checked <- "WORKING"
			} else {
				checked <- "UNKNOWN"
			}
		}
	}
}
