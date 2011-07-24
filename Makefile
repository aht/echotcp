all: srv tcpcc

srv: srv.go
	6g srv.go && 6l -o srv srv.6

tcpcc: clnt.go
	6g clnt.go && 6l -o tcpcc clnt.6
