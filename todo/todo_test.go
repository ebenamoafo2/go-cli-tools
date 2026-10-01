package todo

import (
	"path/filepath"
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("task name should be %s, got %s", taskName, l[0].Task)
	}

}

func TestComplete(t *testing.T) {
	l := List{}
	l.Add("New Task")

	if l[0].Done {
		t.Fatal("New task should not be completed.")
	}
	if !l[0].CompletedAt.IsZero() {
		t.Fatal("New task should not have a CompletedAt time.")
	}

	if err := l.Complete(1); err != nil {
		t.Fatalf("Complete returned an error: %v", err)
	}

	if !l[0].Done {
		t.Error("Task should be done after Complete(1).")
	}
	if l[0].CompletedAt.IsZero() {
		t.Error("CompletedAt should be set after Complete(1).")
	}
}

func TestDelete(t *testing.T) {
	l := List{}

	tasks := []string{
		"New Task1",
		"New Task2",
		"New Task3",
		"New Task4",
	}
	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}
	err := l.Delete(2)
	if err != nil {
		t.Fatalf("Delete returned an error: %v", err)
	}
	if len(l) != 3 {
		t.Errorf("Expected 3 items, got %d", len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	taskName := "New Task"
	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf := filepath.Join(t.TempDir(), "json")

	if err := l1.Save(tf); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if len(l2) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(l2))
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
