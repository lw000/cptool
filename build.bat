set prjPath=%cd%
echo %prjPath%
cd ../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd %prjPath%
go build -o cptool.exe -v -ldflags="-s -w"

REM go build -gcflags "-m -m"命令，来显示编译器将变量转义到堆的具体操作