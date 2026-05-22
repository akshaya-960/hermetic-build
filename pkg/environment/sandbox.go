package environment

import (
	"fmt"
	"os"
	"os/exec"
)

type Sandbox struct {
	Image   string
	WorkDir string
	Env     []string
}

func New(image string) *Sandbox {
	wd, _ := os.Getwd()
	return &Sandbox{
		Image:   image,
		WorkDir: wd,
		Env:     []string{}, // Default empty env
	}
}

// RunAndExtract executes a command with injected env vars and saves outputs
func (s *Sandbox) RunAndExtract(cmdArgs []string) error {
	os.MkdirAll("bin", 0755)
	binDir := s.WorkDir + "/bin"

	// Construct the podman run arguments
	args := []string{"run", "--rm", "--network", "none"}

	// Inject environment variables
	for _, env := range s.Env {
		args = append(args, "-e", env)
	}

	// Add volume mappings and workdir
	args = append(args,
		"-v", fmt.Sprintf("%s:/src:ro", s.WorkDir),
		"-v", fmt.Sprintf("%s:/out", binDir),
		"-w", "/src",
		s.Image,
	)

	// Append the command
	args = append(args, cmdArgs...)

	command := exec.Command("podman", args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}