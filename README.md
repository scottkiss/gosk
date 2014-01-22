# About gosk
======

  gosk is a static site generator written in Go.
  
## Getting started
```bash
$ go get github.com/scottkiss/gosk
```

## Compile gosk
```bash
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ go build gosk-create.go
```
if build passing,do the next.

## Create Site(Blog)
```bash
$ cd bin
#run the gosk-create(if in windows,will be gosk-create.exe) file
$ ./gosk-create
```
if there is no errors,Congratulations,a folder named public will be created under the root folder,then you can host the publish folder use any web server which support serving static content.In gosk,it also provide a simple static web server named gosk-server for testing in local.

## Use the built-in server gosk-server
```bash
$ go get github.com/scottkiss/gostas
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ go build gosk-server.go
```
then run the gosk-server,It will run on http://localhost:80/
(Notic: Don't move the gosk-server file,keep it under bin folder,and don't move the root folder too,Or it won't work)

Now,Open your web browser and visit: http://127.0.0.1:80.Enjoy it.
(====more guide coming soon====)



## License
View the [LICENSE](https://github.com/scottkiss/gosk/blob/master/LICENSE) file