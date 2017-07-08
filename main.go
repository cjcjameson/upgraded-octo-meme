package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ScottLikesThisGPDBNode() (bool, error) {
	return true, nil
}

func main() {
	cmd := exec.Command("/usr/local/gpdb/bin/pg_ctl", "status", "-w", "-D",
		"/Users/pivotal/workspace/gpdb/gpAux/gpdemo/datadirs/dbfast_mirror1/demoDataDir0")
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	nodeAppreciation, err := ScottLikesThisGPDBNode()
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
