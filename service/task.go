package service

import (
	"latihan-grpc-go/model"
	"latihan-grpc-go/repository"
)

// AddTask untuk menambahkan data task
func AddTask(taskData model.Task) model.Task {
	return repository.AddTask(taskData)
}

// GetTask untuk mendapatkan data task berdasarkan id
func GetTask(taskId string) (int, model.Task) {
	return repository.GetTask(taskId)
}

// GetTasks untuk mendapatkan seluruh data task
func GetTasks() []model.Task {
	return repository.GetTasks()
}

// UpdateTask untuk mengedit data task
func UpdateTask(taskData model.Task, id string) model.Task {
	return repository.UpdateTask(taskData, id)
}

// DeleteTask untuk menghapus data task
func DeleteTask(id string) bool {
	return repository.DeleteTask(id)
}
