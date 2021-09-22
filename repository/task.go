package repository

import (
	"database/sql"
	"latihan-grpc-go/database"
	"latihan-grpc-go/model"
	"log"

	"github.com/google/uuid"
)

// AddTask untuk menambahkan data task
func AddTask(taskData model.Task) model.Task {

	// membuat id menggunakan uuid
	var uuid string = uuid.New().String()

	// memasukkan data task ke dalam database
	_, err := database.DB.Query("INSERT INTO tasks (id, title, author, is_read) VALUES (?, ?, ?, ?)",
		uuid,
		taskData.Title,
		taskData.Description,
		taskData.Completed,
	)

	// jika terdapat error, tampilkan pesan error
	if err != nil {
		log.Fatalf("Insert data failed: %v", err)
		return model.Task{}
	}

	// mengembalikan data task yang dimasukkan
	taskData.Id = uuid

	return taskData

}

// GetTask untuk mendapatkan data task berdasarkan id
func GetTask(taskId string) (int, model.Task) {

	// membuat variabel task
	// untuk menyimpan data task berdasarkan id
	var task model.Task = model.Task{}

	// mendapatkan data task berdasarkan id
	row, err := database.DB.Query("SELECT * FROM tasks WHERE id = ?", taskId)

	// menampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return 0, model.Task{}
	}

	// Close() akan dipanggil
	// jika data dari database sudah didapatkan
	defer row.Close()

	// untuk setiap baris data
	for row.Next() {
		// masukkan berbagai atribut data task seperti judul dll.
		// ke dalam variabel task
		switch err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Completed); err {

		// jika data tidak ditemukan
		// tampilkan pesan error
		case sql.ErrNoRows:
			log.Printf("Data not found: %v", err)
			return 0, model.Task{}

		// jika tidak ada error
		// tampilkan data task
		case nil:
			log.Println(task)

		// tampilkan error untuk kondisi default
		default:
			log.Printf("Data cannot be retrieved: %v", err)
			return 0, model.Task{}
		}
	}

	// mengembalikan angka 1 sebagai tanda data ditemukan
	// dan data task yang ditemukan
	return 1, task

}

// GetTasks untuk mendapatkan seluruh data task
func GetTasks() []model.Task {

	// mendapatkan seluruh data task
	rows, err := database.DB.Query("SELECT * FROM tasks")

	// tampilkan pesan error jika terdapat error
	if err != nil {
		log.Fatalf("Data cannot be retrieved: %v", err)
		return []model.Task{}
	}

	// close dipanggil jika
	// seluruh data berhasil diambil
	defer rows.Close()

	// membuat variabel tasks
	// untuk menampung berbagai data task
	var tasks []model.Task = []model.Task{}

	// untuk setiap data
	for rows.Next() {
		// membuat variabel task
		// untuk menyimpan sebuah data task
		var task model.Task = model.Task{}

		// masukkan berbagai atribut data task seperti judul dll.
		// ke dalam variabel task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Completed)

		// jika terdapat error, tampilkan error
		if err != nil {
			log.Printf("Data cannot be retrieved: %v", err)
			return []model.Task{}
		}

		// masukkan data task ke dalam tasks
		tasks = append(tasks, task)
	}

	// jika jumlah data di dalam tasks
	// sama dengan 0, maka data kosong
	if len(tasks) == 0 {
		log.Println("Tasks data not found")
	}

	// mengembalikan sekumpulan data task
	return tasks

}

// UpdateTask untuk mengedit data task
func UpdateTask(taskData model.Task, id string) model.Task {

	// mengubah data task berdasarkan id
	_, err := database.DB.Query("UPDATE tasks SET title=?, author=?, is_read=? WHERE id=?",
		taskData.Title,
		taskData.Description,
		taskData.Completed,
		taskData.Id,
	)

	// jika terdapat error, tampilkan error
	if err != nil {
		log.Fatalf("Update data failed: %v", err)
		return model.Task{}
	}

	// mengembalikan data task yang telah diubah
	return taskData

}

// DeleteTask untuk menghapus data task
func DeleteTask(id string) bool {

	// menghapus data task berdasarkan id
	_, err := database.DB.Query("DELETE FROM tasks WHERE id=?", id)

	// jika terdapat error, tampilkan pesan error
	// nilai false dikembalikan
	if err != nil {
		log.Fatalf("Delete data failed: %v", err)
		return false
	}

	// nilai true dikembalikan jika data berhasil dihapus
	return true
}
