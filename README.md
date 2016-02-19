# About gosk
======
  gosk is a static site generator written in Go.


## Features
* Markdown support
* Custom theme support 

## Getting started
install golang first, if you don't have,see [http://golang.org/doc/install](http://golang.org/doc/install)
(Notice:go version must >=1.2)
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
(Notic: Don't move the gosk(gosk.exe) file,keep it under the bin folder,and don't move the root folder too,Or it won't work)

Now,Open your web browser and visit: http://localhost/  - Enjoy it.

## project category
```bash
bin
 | - publish	             #published folder,include compiled .html files
 	  | - ...	             #assets 
 	  | - index.html
 	  | - rss.xml
 	  | - ...                #more html files
 | - root					 #root folder
      | - assets             #assets folder,include javascript and css files
      | - pages              #custom pages folder
      | - posts              #post folder
      	   | - article1.md   #metadata text file
      	   | - article2.md   #metadata text file
      | - templates          #template folder
           | - default       #default theme template
           | - default-zh    #default-zh theme template
           | - ...		     #more themes template
      | - config.yml         #site global configure
      | - nav.yml            #navbar configure
      | - pages.yml          #custom site configure
 | - gosk                    #gosk  file
```

## more themes
[gosk-theme](https://github.com/scottkiss/gosk-theme)

## who use gosk
* [cocosk](http://www.cocosk.com/)


If you are using gosk too,please tell me by email.


# 中文指南
------

## 特点
* 支持Markdown 
* 支持自定义主题
* 自带默认一个主题，支持代码高亮
* 编译文章速度快
* 更多特点,谁用谁知道^_^

## 开始使用
```bash
$ go get github.com/scottkiss/gosk
```

## 编译 gosk
```bash
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ go build gosk.go
```
编译通过后，进行下一步

## 创建 静态页面(博客)
```bash
$ cd bin
#运行编译生成的gosk(如果在windows平台下，是gosk.exe) 文件
$ ./gosk build(在windows平台下运行gosk build)
```
如果没什么出错提示,那么恭喜你,在当前目录下会创建一个叫public目录,然后你可以用任何
支持静态服务的服务器部署public目录。为了方便测试，gosk也内置了一个简陋的静态web服务器。

## 使用内置静态服务器
```bash
$ cd $GOPATH/github.com/scottkiss/gosk/bin
$ ./gosk run :80
```
如果不指定后面的参数（执行 ./gosk run）默认是运行在 http://localhost:8080/ 
(注意:不要移动gosk（gosk.exe）的位置，也不要移动root目录的位置，不然就无法成功编译或运行了。)

现在,打开您的浏览器并访问: http://localhost/ 

## 项目目录结构
```bash
bin
 | - publish	             #执行编译后生成的目录，静态站点根目录
 	  | - ...	             #资源，如javascript和css 
 	  | - index.html         #生成的首页文件
 	  | - rss.xml
 	  | - ...                #更多html文件
 | - root					 #根目录，存放待编译的模板文件等
      | - assets             #资源目录,包括javascript 和 css 文件
      | - pages              #自定义页面目录
      | - posts              #发布的文章目录
      	   | - article1.md   #markdown编写的元文本文件
      	   | - article2.md   #markdown编写的元文本文件
      | - templates          #模板目录
           | - default       #默认主题模版
           | - default-zh    #默认中文主题模版
           | - ...		     #更多模版主题
      | - config.yml         #站点全局配置文件
      | - nav.yml            #站点导航栏配置文件
      | - pages.yml          #自定义页面配置文件
 | - gosk                    #gosk  执行文件
```

## 更多主题
gosk内置一个default主题，更多主题请前往[gosk-theme](https://github.com/scottkiss/gosk-theme)

## 使用gosk的站点
* [酷酷时空](http://www.cocosk.com/)

如果你也使用gosk，如果不介意，通过邮件告诉我哦。

## License
View the [LICENSE](https://github.com/scottkiss/gosk/blob/master/LICENSE) file
