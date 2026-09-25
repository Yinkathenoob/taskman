package main

import "testing"

func TestMarkDone(t *testing.T) {
	// Arrange: a known set of tasks
	tasks := []Task{
		{ID: 1, Title: "first", Done: false},
		{ID: 2, Title: "second", Done: false},
	}

	// Act: mark task 2 done
	result := markDone(tasks, 2)

	// Assert: it should have returned true
	if !result {
		t.Errorf("markDone(tasks, 2) = false; want true")
	}

	// Assert: task 2's Done should now be true
	if !tasks[1].Done {
		t.Errorf("task 2 Done = false; want true")
	}

	// Assert: task 1 should be untouched
	if tasks[0].Done {
		t.Errorf("task 1 Done = true; want false (should be unchanged)")
	}
}

func TestMarkDoneNotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "first", Done: false},
	}

	// Act: try to mark a task that doesn't exist
	result := markDone(tasks, 99)

	// Assert: should return false
	if result {
		t.Errorf("markDone(tasks, 99) = true; want false (no such task)")
	}
}

func TestAddTask(t *testing.T) {
	// Arrange: start with an empty task list and a next ID of 1
	tasks := []Task{}
	nextID := 1

	// Act: add a new task
	tasks, nextID = addTask(tasks, "new task", nextID)

	// Assert: the task list should have one task with the correct ID and title
	if len(tasks) != 1 {
		t.Errorf("len(tasks) = %d; want 1", len(tasks))
	}
	if tasks[0].ID != 1 {
		t.Errorf("tasks[0].ID = %d; want 1", tasks[0].ID)
	}
	if tasks[0].Title != "new task" {
		t.Errorf("tasks[0].Title = %q; want %q", tasks[0].Title, "new task")
	}
	if tasks[0].Done {
		t.Errorf("tasks[0].Done = true; want false")
	}

	// Assert: nextID should have incremented
	if nextID != 2 {
		t.Errorf("nextID = %d; want 2", nextID)
	}
}
