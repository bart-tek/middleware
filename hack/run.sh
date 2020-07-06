#! bash
# build all main.go
echo "Building redis_sub..."
go build -o ./bin/redis_sub   ./cmd/redis_sub/
echo "Ok"
echo "Building csv_sub..."
go build -o ./bin/csv_sub     ./cmd/csv_sub/
echo "Ok"
echo "Building temp_captor..."
go build -o ./bin/temp_captor ./cmd/captors/temp_captor/
echo "Ok"
echo "Building wind_captor..."
go build -o ./bin/wind_captor ./cmd/captors/wind_captor/
echo "Ok"
echo "Building pres_captor..."
go build -o ./bin/pres_captor ./cmd/captors/pressure_captor/
echo "Ok"
echo "Building json_serv..."
go build -o ./bin/json_serv   ./cmd/json_serv/
echo "Ok"