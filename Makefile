generate-mock:
	echo "Generating mocks"
	go generate -v ./...

run-tests:
	echo "Running all tests"
	go test ./tests/... -v