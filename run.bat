
start CMD /k "cd frontend && rmdir /s dist && npm run build && cd .. && cd backend && go env -w GOPROXY=direct && go mod tidy && go run ."