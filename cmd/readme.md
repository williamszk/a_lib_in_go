
# Experiment

https://ieftimov.com/posts/golang-package-multiple-binaries/

How should we build a module with many executables?

Inside the cmd dir which lives in the root of the project,
we create the directories "hi" and "bye" those could be other names.
Then we create main.go in both of them.
With the package as "main"

Then we need to go into those dirs...

```go
cd cmd/hi/
go build
```

Then build it.