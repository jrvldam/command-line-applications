package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("add new task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("list tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		want := task + "\n"

		if want != string(out) {
			t.Errorf("got %s, want %s", string(out), want)
		}
	})

	t.Run("complete task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		cmd.Run()

		cmdL := exec.Command(cmdPath, "-list")
		out, err := cmdL.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		want := ""

		if want != string(out) {
			t.Errorf("got %s, want %s", string(out), want)
		}
	})

	t.Run("no flags passed", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatal("error expected")
		}

		if string(out) != "Invalid option\n" {
			t.Errorf("got %s, want %s", string(out), "Invalid Option\n")
		}
	})
}
