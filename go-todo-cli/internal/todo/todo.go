package todo

import (
	"errors"
	"sync"
)

type Todo struct {
	ID   string
	Task string
	Done bool
}

type TodoList struct {
	todos map[string]*Todo
	mu    sync.Mutex
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos: make(map[string]*Todo),
	}
}

func (tl *TodoList) Add(todo *Todo) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.todos[todo.ID] = todo
}

func (tl *TodoList) Get(id string) (*Todo, error) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	todo, exists := tl.todos[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (tl *TodoList) Delete(id string) error {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	if _, exists := tl.todos[id]; !exists {
		return errors.New("todo not found")
	}
	delete(tl.todos, id)
	return nil
}

func (tl *TodoList) List() []*Todo {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	todos := make([]*Todo, 0, len(tl.todos))
	for _, todo := range tl.todos {
		todos = append(todos, todo)
	}
	return todos
}