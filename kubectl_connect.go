package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "kubectl connect"}

var srcPod string
var destPod string
var port int
var namespace string
var timeout int

func main() {
	rootCmd.PersistentFlags().StringVarP(&srcPod, "source", "s", "", "Source pod name")
	rootCmd.PersistentFlags().StringVarP(&destPod, "destination", "d", "", "Destination pod name")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "Namespace")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 22, "Port")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 10, "Timeout in seconds")

	rootCmd.AddCommand(versionCmd)
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if srcPod == "" || destPod == "" {
			fmt.Println("Both Source and destination pods are required.")
			os.Exit(1)
		}

		fmt.Printf("Checking connectivity from %s to %s on port %d in namepsace %s....\n", srcPod, destPod, port, namespace)

		start := time.Now()
		err := checkConnectivity(srcPod, destPod, port, namespace, timeout)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("Connectivity between %s and %s is OK.  Took %s.\n", srcPod, destPod, elapsed)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kubectl-connect",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubectl-connect v1.0.0")
	},
}

func getPodIP(namespace, destPod string) (string, error) {
	cmdStr := fmt.Sprintf("kubectl get pod -n %s %s -o jsonpath='{.status.podIP}'", namespace, destPod)
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func checkConnectivity(srcPod, destPod string, port int, namespace string, timeout int) error {
	destIP, err := getPodIP(namespace, destPod)
	if err != nil {
		return err
	}

	cmdStr := fmt.Sprintf("kubectl exec -it -n %s %s -- /bin/bash -c 'timeout %d bash -c \"</dev/tcp/%s/%d\"'", namespace, srcPod, timeout, destIP, port)
	cmd := exec.Command("bash", "-c", cmdStr)

	var output, stderr bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		fmt.Printf("FAIL!\nConnectivity between %s and %s on port %d failed. - Output:\n%s\n", srcPod, destPod, port, output)
		return err
	}

	fmt.Printf("SUCCESS!\n")
	return nil
}
