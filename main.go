//go:build linux

package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const LINUX_URL = "https://discord.com/api/download/stable?platform=linux&format=tar.gz"

func main() {

	startTime := time.Now()

	installPath := flag.String("i", "/usr/share/discord", "path to install discord")

	flag.Parse()

	resp, err := http.Get(LINUX_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("[Download] discord.tar.gz (size : %d)\n", resp.ContentLength)

	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}

	defer gr.Close()

	r := tar.NewReader(gr)

	var totalWritten int64

	for {

		header, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		dest := filepath.Join(*installPath, strings.TrimPrefix(header.Name, "Discord/"))
		fPerm := header.FileInfo().Mode().Perm()

		fmt.Printf("[Extract] %s (size : %d)\n", dest, header.Size)

		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(dest, fPerm); err != nil {
				panic(err)
			}
			continue
		}

		f, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fPerm)
		if err != nil {
			panic(err)
		}

		n, err := io.Copy(f, r)
		if err != nil {
			panic(err)
		}

		totalWritten += n

	}

	fmt.Println("Running post install script")

	cmd := exec.Command(filepath.Join(*installPath, "postinst.sh"))

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Printf("Done in %2.f seconds\n", time.Since(startTime).Seconds())

}
