git add .
git commit -m "Laast changes"
git push
set GOOS=linux
go build main.go
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main bootstrap