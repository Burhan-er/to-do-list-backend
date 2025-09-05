package db

import (
	"encoding/json"
	"os"
)
type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}
type Db struct {
	tasksFilePath string
}

func NewDatabase(filePath string) *Db {
	return &Db{tasksFilePath: filePath}
}

func (db *Db) saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(db.tasksFilePath, data, 0644)
}

func (db *Db) loadTasks() ([]Task, error) {
	if _, err := os.Stat(db.tasksFilePath); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(db.tasksFilePath)

	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}

func (db *Db) Add(text string) error {
    tasks, err := db.loadTasks()
    if err != nil {
        return err
    }
    
    newTask := Task{Text: text, Done: false}
    tasks = append(tasks, newTask)

    return db.saveTasks(tasks)
}

func (db *Db) Delete(index int) error {
    tasks, err := db.loadTasks()
    if err != nil {
        return err
    }

    if index < 0 || index >= len(tasks) {
        return os.ErrInvalid // Veya özel bir hata döndürün
    }

    // Görevi listeden kaldırma
    tasks = append(tasks[:index], tasks[index+1:]...)

    return db.saveTasks(tasks)
}
func (db *Db) Complete(index int) error {
    tasks, err := db.loadTasks()
    if err != nil {
        return err
    }

    if index < 0 || index >= len(tasks) {
        return os.ErrInvalid
    }

    tasks[index].Done = true

    return db.saveTasks(tasks)
}
func (db *Db) GetTasks() ([]Task, error) {
    return db.loadTasks()
}