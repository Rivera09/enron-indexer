@echo off
cd ./mamuro
echo creating build files...
call npm run build
echo serving files
call go run ./server.go
pause