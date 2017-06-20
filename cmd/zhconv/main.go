package main

import (
	"fmt"
	"github.com/jiusanzhou/tentacle/log"
	"github.com/jiusanzhou/zhconv/pkg/zhconv"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(len(opts.inputs))

	fmt.Printf("即將轉換%d個文件\n", len(opts.inputs))
	for _, i := range opts.inputs {

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			bs, err := ioutil.ReadFile(i)
			if err != nil {
				log.Error("Read file error, %s, %s", i, err.Error())
				return
			}

			var sb []byte
			switch opts.cmd {
			case "2s":
				sb = conv2s(string(bs))
			case "2t":
				sb = conv2t(string(bs))
			default:
				fmt.Println("Unknow command:", opts.cmd)
				return
			}

			_path, name := path.Split(i)
			_path = path.Join(_path, opts.outputDir)

			if s, err := os.Stat(_path); err != nil && os.IsNotExist(err) {
				os.MkdirAll(_path, 0700)
			} else {
				if !s.IsDir() {

					os.MkdirAll(_path, 0700)
				}
			}

			f, err := os.Create(path.Join(_path, name))
			if err != nil {
				fmt.Println("Create file:", name, "in:", _path, "error:", err.Error())
				return
			}

			f.Write(sb)
			fmt.Print(name, "轉換完成\n")
		}(&wg)
	}

	fmt.Println("等待轉換完成")
	wg.Wait()
	fmt.Println("全部轉換完成")
}

func conv2s(s string) (d []byte) {
	for _, i := range s {
		str := string(i)
		if zhconv.IsChinese(str) {
			d = append(d, zhconv.ConvertToSimplifiedChinese(str)...)
		} else {
			d = append(d, str...)
		}
	}

	return
}

func conv2t(s string) (d []byte) {
	for _, i := range s {
		str := string(i)
		if zhconv.IsChinese(str) {
			d = append(d, zhconv.ConvertToTraditionalChinese(str)...)
		} else {
			d = append(d, str...)
		}
	}

	return
}
