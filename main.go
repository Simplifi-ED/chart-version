package main

import (
	"fmt"
	"github.com/spf13/cobra"
	// "gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "chart-version",
		Short: "Change the version of a Helm chart",
		Run:   changeChartVersion,
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func changeChartVersion(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: chart-version [chart-path1 chart-path2 ...] [new-version]")
		os.Exit(1)
	}

	chartPaths := args[:len(args)-1]
	newVersion := args[len(args)-1]

	for _, chartPath := range chartPaths {
		chartFile := filepath.Join(chartPath, "Chart.yaml")
		chartData, err := os.ReadFile(chartFile)
		if err != nil {
			fmt.Printf("Failed to read Chart.yaml: %v\n", err)
			os.Exit(1)
		}

		chartLines := strings.Split(string(chartData), "\n")
		for i, line := range chartLines {
			if strings.HasPrefix(line, "version:") {
				chartLines[i] = "version: " + newVersion
				break
			}
		}

		updatedChartData := []byte(strings.Join(chartLines, "\n"))
		err = os.WriteFile(chartFile, updatedChartData, 0644)
		if err != nil {
			fmt.Printf("Failed to write updated Chart.yaml: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully updated version to %s in Chart.yaml\n", newVersion)
	}
}
