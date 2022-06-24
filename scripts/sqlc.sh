if [ ! -d "scripts" ]; then
  cd ..
fi
go get github.com/kyleconroy/sqlc/cmd/sqlc
cd internal/database
go run github.com/kyleconroy/sqlc/cmd/sqlc generate
go mod tidy
