package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ArtemHvozdov/task-cli/internal/entity"
)

type JSONStorage struct {
	pathFile string
}

func NewJSONStorage(pathFile string) (*JSONStorage, error) {
	_, err := os.Stat(pathFile)

	if os.IsNotExist(err) {
		file, createErr := os.Create(pathFile)
		if createErr != nil {
			return nil, createErr
		}

		defer file.Close()

		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &JSONStorage{
		pathFile: pathFile,
	}, nil
}

func (s *JSONStorage) Add(description string) (uint16, error) {
	tasks, err := s.load()
	if err != nil {
		return 0, err
	}

	var taskID uint16 = 1
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}

	task := entity.Task{
			ID:          taskID,
			Description: description,
			Status:      entity.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

	tasks = append(tasks, task)

	err = s.save(tasks)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
	
}

func (s *JSONStorage) Update(id uint16, desscription string) error {
	tasks, err := s.load()
	if err != nil {
		return err
	}

	var task *entity.Task
	for i := range tasks {
		if tasks[i].ID == id {
			task = &tasks[i]
			break
		}
	}

	if task == nil {
		return fmt.Errorf("task with id %d not found", id)
	}

	task.Description = desscription
	task.UpdatedAt = time.Now()
	err = s.save(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) Delete(id uint16) error {
	tasks, err := s.load()
	if err != nil {
		return err
	}

	index := -1
	for i := range tasks {
		if tasks[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with id %d not found", id)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	err = s.save(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) UpdateStatus(id uint16, newStatus entity.Status) error {
	tasks, err := s.load()
	if err != nil {
			return err
		}

	var task *entity.Task

	for i := range tasks {
		if tasks[i].ID == id {
				task = &tasks[i]
				break
			}
		}

	if task == nil {
		return fmt.Errorf("task with id %d not found", id)
	}

	task.Status = newStatus
	task.UpdatedAt = time.Now()

	err = s.save(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) GetAll() ([]entity.Task, error) {
	tasks, err := s.load()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *JSONStorage) GetByStatus(status entity.Status) ([]entity.Task, error) {
	tasks, err := s.load()
	if err != nil {
		return nil, err
	}

	var filteredTasks []entity.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}

func (s *JSONStorage) load() ([]entity.Task, error) {
	data, err := os.ReadFile(s.pathFile)
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *JSONStorage) save(tasks []entity.Task)  error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.pathFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}