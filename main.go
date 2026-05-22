package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/akshaya-960/hermetic-build/pkg/environment"
	"github.com/akshaya-960/hermetic-build/pkg/utils"
)

type Config struct {
	Image        string   `json:"image"`
	OutputBinary string   `json:"output_binary"`
	Env          []string `json:"env"`
}

func main() {
	// 1. Read Configuration
	configFile, _ := os.ReadFile("build.json")
	var cfg Config
	json.Unmarshal(configFile, &cfg)

	// 2. Initialize Sandbox with Config
	sandbox := environment.New(cfg.Image)
	sandbox.Env = cfg.Env

	// 3. Execute Build
	outputPath := "/out/" + cfg.OutputBinary
	buildCmd := []string{"go", "build", "-o", outputPath, "."}
	
	fmt.Println("Starting Hermetic Build using build.json...")
	err := sandbox.RunAndExtract(buildCmd)
	
	if err == nil {
		hash, _ := utils.GenerateSHA256("bin/" + cfg.OutputBinary)
		fmt.Printf("Build success! Hash: %s\n", hash)
	}
}