
go mod init github.com/williamszk/a_lib_in_go 

go mod init github.com/williamszk/a_lib_in_go

go test

go run main.go

go build

git add .
git commit -m "just adding tag to git"
git commit -m "fix problem with package main, to bananas"
git push

git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# we need to tag
git tag "v1.0.0"
git tag "v1.0.1"

git push --tags

git add .
git commit -m "add another function for patch 2"
git tag "v1.0.2"
git push