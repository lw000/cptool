set prjPath=%cd%
echo %prjPath%
cd ../../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
cd %prjPath%

set buildTime=%date:~0,4%%date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%%time:~6,2%
set BuildVersion=1
set BuildName=1
set CommitID=1

go build -a -o cptool -v -ldflags "-s -w -X 'main.BuildVersion=%BuildVersion%' -X 'main.BuildTime=%buildTime%' -X 'main.BuildName=%BuildName%' -X 'main.CommitID=%CommitID%'"