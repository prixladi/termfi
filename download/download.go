package download

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/prixladi/termfi/utils"
)

func DownloadFile(url string, dest string, fileSize int64) {
	fmt.Printf("Downloading file '%s' from %s\n", dest, url)

	var path bytes.Buffer
	path.WriteString(dest)

	start := time.Now()

	out, err := os.Create(path.String())
	if err != nil {
		fmt.Println(path.String())
		panic(err)
	}
	defer out.Close()

	done := make(chan int64)
	go printDownloadPercent(done, path.String(), int64(fileSize))

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	done <- n

	utils.ReplacefLine("Download to local file '%s' completed in %s", dest, time.Since(start))
}

func printDownloadPercent(done chan int64, path string, total int64) {
	bar := pb.New64(total)
	utils.ReplaceLine(bar.String())

	var stop bool = false

	for i := 0; !stop; i++ {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(path)
			if err != nil {
				fmt.Print(err)
			}
			defer file.Close()

			fi, err := file.Stat()
			if err != nil {
				fmt.Print(err)
			}

			size := fi.Size()

			bar.SetCurrent(size)
			utils.ReplaceLine(bar.String())

			if !stop {
				time.Sleep(time.Second)
			}
		}
	}
}
