# About gosk
======
  gosk is a static site generator written in Go.


## Features
* Markdown support
* Custom theme support 

## Getting started
```bash
$ go get github.com/scottkiss/gosk
```

## Compile gosk
```bash
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ go build gosk.go
```
if build passing,do the next.

## Create Site(Blog)
```bash
$ cd bin
#run the gosk(if in windows os,will be gosk.exe) file
$ ./gosk build(if in windows os : gosk build)
```
if there is no errors,Congratulations,a folder named public will be created in current folder,then you can host the publish folder use any web server which support serving static content.In gosk,it also provide a simple static web server  for testing in local.

## Use the built-in server gosk-server
```bash
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ ./gosk run :80
```
It will run on http://localhost:8080/ if don't specify the port
(Notic: Don't move the gosk(gosk.exe) file,keep it under bin folder,and don't move the root folder too,Or it won't work)

Now,Open your web browser and visit: http://localhost/  - Enjoy it.




## License
View the [LICENSE](https://github.com/scottkiss/gosk/blob/master/LICENSE) file