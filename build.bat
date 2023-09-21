SET GOOS=linux
SET GOARCH=amd64
go build  -o ./build/dory-mall

for /f %%i in ('git describe --tags --abbrev^=0') do ( SET version=%%i)
for /f %%i in ('git rev-parse --short HEAD') do ( SET commitId=%%i)
SET tag=feverzy/repo:dory-mall-%version%-%commitId%
docker build -t %tag% .
docker push %tag%

pause