
start CMD /k "cd frontend && npm run build "
start CMD /k "cd .. && cd backend && go env -w GOPROXY=direct && go mod tidy && go run ."