all: go

go: srv.go
	6g srv.go && 6l -o srv srv.6
