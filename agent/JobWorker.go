package agent

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func JobWorker(worker chan Job, reporter chan Result) {
	for {
		select {
		case j := <-worker:
			go func(job Job) {
				var out bytes.Buffer
				cmd := exec.Command("bash", "-c", job.Command)

				cmd.Stdout = &out
				cmd.Stderr = &out
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
			}(j)
		}
	}
}
