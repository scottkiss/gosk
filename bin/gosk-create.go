package main

import (
	"bufio"
	"github.com/scottkiss/gosk"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	RENDER_DIR   = "../root"
	PUBLICSH_DIR = "publish"
)

func main() {
	//publish
	if !isExists(RENDER_DIR + "/" + PUBLICSH_DIR) {
		err := os.Mkdir(RENDER_DIR+"/"+PUBLICSH_DIR, 0777)
		if err != nil {
			log.Panic("create publish dir error -- " + err.Error())
		}
	}
	//render all files
	var rf = new(gosk.RenderFactory)
	rf.Render(RENDER_DIR)

	//copy res
	err := copyDir(RENDER_DIR+"/assets", RENDER_DIR+"/"+PUBLICSH_DIR+"/assets")
	if err != nil {
		log.Println(err)
	}
	log.Println("blog process ok！")
	rd := bufio.NewReader(os.Stdin)
	rd.ReadLine()

}

func isExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func copyFile(src, dst string) (w int64, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()
	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer dstf.Close()
	return io.Copy(dstf, f)
}

func copyDir(source, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// _, err = os.Open(dest)
	// if !os.IsNotExist(err) {
	// 	return &CustomError{"Destination already exists"}
	// }

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			_, err = copyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}
