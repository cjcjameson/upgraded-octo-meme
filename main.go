package main

import (
	"github.com/cjcjameson/pg_ctl"
	"fmt"
	"os"
	"strings"
)

func DoesScottLikeThisGPDBNode(status pg_ctl.Status) (bool, error) {
	fmt.Println(status.RawStdOut)
	bigEnoughPid := status.Pid > 30000
	isQuiescent := strings.Contains(status.PsPostgres, "quiescent")
	return status.IsServerRunning && bigEnoughPid && isQuiescent, nil
}

func main() {
	controller := pg_ctl.NewController("/Users/pivotal/workspace/gpdb/gpAux/gpdemo/datadirs/dbfast1/demoDataDir0")

	status, err := controller.Status()
	if err != nil {
		fmt.Println(err)
		os.Exit(status.ErrorCode)
	}

	nodeAppreciation, err := DoesScottLikeThisGPDBNode(status)
	if err != nil {
		fmt.Println("Unable to determine Scott's preference")
		fmt.Println(err)
		os.Exit(1)
	}

	if nodeAppreciation {
		fmt.Println("Scott Kahler is a big fan of this segment")
	} else {
		fmt.Println("Scott Kahler isn't wild about how this postgres process is running")
	}

	os.Exit(0)
}
