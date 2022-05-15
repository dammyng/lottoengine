run-test:
	go test ./...
start-service: 
	cd cmd && go build -o bin/lottoengine main.go && cd bin && ./lottoengine