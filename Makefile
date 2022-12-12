bin_dir=build
tool=cryptool
psql_connection=postgresql://postgres:postgres@127.0.0.1:45433/creator_suite_test


test-prepare:
	@psql ${psql_connection} -c 'CREATE DATABASE creator_suite_test;'

test:
	@go test -v ./...

psql:
	@psql ${psql_connection}

clean:
	@rm -rf build

build:clean
	@go build -o ${bin_dir}/ ./cmd/...

install-tool:
	@go build -o ${tool} ./cmd/${tool}/main.go