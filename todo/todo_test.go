package todo_test

import (
	"os"
	"testing"

	"github.com/jrvldam/command-line-applications/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("got %q, want %q", l[0].Task, taskName)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("got %d, want %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("got %q, want %q", l[1].Task, tasks[2])
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	file, err := os.CreateTemp("", "todo")
	if err != nil {
		t.Fatal("fail creating temporal file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	if err := l1.Save(file.Name()); err != nil {
		t.Fatalf("error saving list to file, %s", err)
	}

	if err := l2.Get(file.Name()); err != nil {
		t.Fatalf("error getting list from file, %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match Task %q", l1[0].Task, l2[0].Task)
	}
}
