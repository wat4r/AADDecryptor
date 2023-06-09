set VERSION=1.0.0

@REM Linux
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -trimpath -o bin/AADDecryptor_%VERSION%_%GOOS%_%GOARCH% main.go

@REM Mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-s -w" -trimpath -o bin/AADDecryptor_%VERSION%_%GOOS%_%GOARCH% main.go

@REM Windows
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -trimpath -o bin/AADDecryptor_%VERSION%_%GOOS%_%GOARCH%.exe main.go
