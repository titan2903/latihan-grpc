package server

import (
	"context"
	"errors"
	"latihan-grpc-go/model"
	"latihan-grpc-go/service"
	taskpb "latihan-grpc-go/task/taskpb"
)

type Server struct {
	
}

func mapToTaskpb(task model.Task) *taskpb.Task {
	return &taskpb.Task{
		Id:     task.Id,
		Title: task.Title,
		Description:  task.Description,
		Completed: task.Completed,
	}
}

// Get Task untuk mendapatkan data task berdasarkan id
func (*Server) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	// mengambil inputan id dari request
	var taskId string = req.GetId()

	// memanggil GetTask untuk mendapatkan data task berdasarkan id
	_, result := service.GetTask(taskId)

	// jika tidak ditemukan, beri pesan error
	if result.Id != taskId {
		return &taskpb.GetTaskResponse{Status: false, Data: nil}, errors.New("Data not found!")
	}

	// mengolah hasil dari GetTask
	// untuk data di objek response
	var taskData *taskpb.Task = &taskpb.Task{
		Id:     result.Id,
		Title:  result.Title,
		Description: result.Description,
		Completed: result.Completed,
	}

	// memberikan response berupa data task
	return &taskpb.GetTaskResponse{
		Status: true,
		Data:   taskData,
	}, nil
}