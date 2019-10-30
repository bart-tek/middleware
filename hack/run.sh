#! bash
# build all main.go

go build -o ./bin/redis_sub   ./cmd/redis_sub/
go build -o ./bin/csv_sub     ./cmd/csv_sub/
go build -o ./bin/temp_captor ./cmd/captors/temp_captor/
go build -o ./bin/wind_captor ./cmd/captors/wind_captor/
go build -o ./bin/pres_captor ./cmd/captors/pressure_captor/
go build -o ./bin/json_serv   ./cmd/json_serv/

