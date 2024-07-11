//初始化
go mod init module名称（模块名称）
例：go mod init main

//精简构建-1：go build -ldflags "-s -w" -o rsw.exe routine.go
//精简构建-2：go build -ldflags "-s -w -H=windowsgui" -o rswH.exe routine.go