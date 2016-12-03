test:
	go test -v

cover:
	rm -rf *.coverprofile
	go test -coverprofile=mux.coverprofile
	gover
	go tool cover -html=mux.coverprofile