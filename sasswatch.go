package sasswatch

import (
	"github.com/howeyc/fsnotify"
	"github.com/marksteve/go-sasswatch/gosass"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Printf("[error] %s", err)
	}
}

func isSass(file string) bool {
	return path.Ext(file) == ".scss" || path.Ext(file) == ".sass"
}

func compile(name string, options gosass.Options) {
	dir, file := path.Split(name)
	if !strings.HasPrefix(file, "_") && isSass(file) {
		ctx := &gosass.FileContext{
			Options:   options,
			InputPath: name,
		}
		gosass.CompileFile(ctx)
		if ctx.ErrorStatus != 0 {
			if ctx.ErrorMessage != "" {
				log.Printf("[error] %s", ctx.ErrorMessage)
			} else {
				log.Print("[error] Unknown error")
			}
		} else {
			fn := path.Join(dir, strings.TrimSuffix(file, path.Ext(file))+".css")
			fi, err := os.Stat(name)
			checkErr(err)
			ioutil.WriteFile(fn, []byte(ctx.OutputString), fi.Mode())
			log.Printf("[compile] %s", fn)
		}
	}
}

func SassWatcher(root string, options gosass.Options) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	checkErr(err)
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				switch {
				case ev.IsCreate():
				case ev.IsModify():
					compile(ev.Name, options)
				}
			case err := <-watcher.Error:
				log.Println("[error]", err)
			}
		}
	}()
	checkErr(filepath.Walk(root, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Watch(name)
			if err != nil {
				return err
			}
			log.Printf("[watch] %s", name)
		} else {
			compile(name, options)
		}
		return nil
	}))
	return watcher
}
