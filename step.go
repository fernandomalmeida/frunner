package frunner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Step struct {
	Name     string   `yaml:"name"`
	Dir      string   `yaml:"-"`
	Image    string   `yaml:"image"`
	Commands []string `yaml:"commands"`
}

func (s *Step) Run() error {
	args := []string{
		"run",
		"--mount",
		fmt.Sprintf("type=bind,src=%s,target=/opt/workdir", s.Dir),
		"-w",
		"/opt/workdir",
		"--rm",
		s.Image,
	}

	for _, command := range s.Commands {
		splittedCommand := strings.Split(command, " ")
		cmdArgs := append(args, splittedCommand...)
		err := s.command("podman", cmdArgs...).Run()
		if err != nil {
			return fmt.Errorf("error on running command [%s]: %s", command, err)
		}
	}

	return nil
}

func (s *Step) command(name string, args ...string) *exec.Cmd {
	log.Printf("Running command: %s %s on dir: %s\n", name, args, s.Dir)
	cmd := exec.Command(name, args...)
	cmd.Dir = s.Dir
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd
}
