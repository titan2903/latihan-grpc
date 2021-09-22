migration-up:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/latihan_grpc?query" -verbose up

migration-down:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/latihan_grpc?query" -verbose down

migration-force:
	migrate -path db/migration -database "mysql://titan:titan@tcp(localhost:3306)/latihan_grpc?query" force $(force)