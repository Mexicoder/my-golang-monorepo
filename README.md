# My Golang Monorepo

This repository contains several Go modules including a module for a linked list and another module that uses the linked list. Here are the steps for setting it up:

## Setup

1. Go to your workspace:

```
$ cd $HOME/GoProjects/
```

2. Create the new monorepo directory and initialize it as a Git repository:

```
$ mkdir my-golang-monorepo
$ cd my-golang-monorepo
$ git init
```

3. Create a directory for your linked-list module and initialize it as a Go module:

```
$ mkdir linked-list
$ cd linked-list
$ go mod init github.com/Mexicoder/my-golang-monorepo/linked-list
```

4. Write your linked list code in this directory. You can use a simple file like `linkedlist.go`.

5. Go back to the monorepo root and create the directory for your other module where you want to use the linked list:

```
$ cd ..
$ mkdir my-other-module
$ cd my-other-module
```

6. Initialize this as a Go module as well:

```
$ go mod init github.com/Mexicoder/my-golang-monorepo/my-other-module
```

7. In your Go code for `my-other-module`, you can now import the linked list module with its full path:

```
import "github.com/Mexicoder/my-golang-monorepo/linked-list"
```

8. Once you've written your code for `my-other-module`, go back to the root directory of the project:

```
$ cd ..
```

9. At this point, you should be able to run `go build ./...` from the root of the monorepo and see no errors. If everything compiles without errors, you can add all your changes to git and commit:

```
$ git add .
$ git commit -m "Initial commit"
```

10. Now, push your changes to GitHub:

```
$ git remote add origin https://github.com/Mexicoder/my-golang-monorepo.git
$ git push -u origin master
```

## Notes

Remember to replace `my-golang-monorepo` with the name of your repository.

Please note that this setup assumes that `linked-list` and `my-other-module` are both Go modules and are treated as separate packages. It's crucial that they are structured and coded like Go packages (i.e., they should contain package declarations at the start of every Go file, etc.).

The idea behind a monorepo is that you have multiple loosely coupled packages living together. They should be able to be built and tested independently of each other. If you have shared code that `my-other-module` depends on, that code should live in its own Go package (in this case, `linked-list`). The import path for the package should be the full GitHub URL path to allow Go to find and build it.

Remember to turn on `Enable Go modules integration` in GoLand is it still doesn't work

