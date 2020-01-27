package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var fromParam, toParam string
var offsetParam, limitParam int64

func init() {
	flag.StringVar(&fromParam, "from", "", "file to read from")
	flag.StringVar(&toParam, "to", "", "file to write to")
	flag.Int64Var(&limitParam, "limit", 0, "limit of new file")
	flag.Int64Var(&offsetParam, "offset", 0, "offset in input file")
}

func main() {

	flag.Parse()

	err := Copy(fromParam, toParam, limitParam, offsetParam)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

}

// Copy .
func Copy(from string, to string, limit int64, offset int64) (err error) {
	const bufSize int64 = 1024 * 1024

	if from == "" {
		return fmt.Errorf("empty value -from")
	}

	if to == "" {
		return fmt.Errorf("empty value -to")
	}

	if from == to {
		return fmt.Errorf("-from and -to are equal")
	}

	file, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("failed to open file info: %v", err)
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}
	fileSize := fi.Size()

	if offset > fileSize {
		return fmt.Errorf("too big offset. Offset = %v fileSize = %v", offset, fileSize)
	}

	if limit == 0 {
		limit = fileSize
	}

	if limit+offset < fileSize {
		fileSize = limit + offset
	}

	tmpl := "{{counters . }} {{percent . }} \n"
	bar := pb.New64(limit)
	bar.SetTemplateString(tmpl)
	bar.Set(pb.Bytes, true)
	bar.Set(pb.SIBytesPrefix, true)
	bar.SetWriter(os.Stdout)
	//bar.Set(pb.Static, true)
	bar.Start()
	//bar.Write()

	resFile, _ := os.Create(to)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	barReader := bar.NewProxyWriter(resFile)

	for offset < fileSize {
		_, err := file.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("failed to seek: %v", err)
		}
		bufSize := int64(math.Min(float64(bufSize), float64(fileSize-offset)))
		buf := make([]byte, bufSize)
		read, err := file.Read(buf)
		offset += int64(read)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read: %v", err)
		}

		_, err = barReader.Write(buf)
		if err != nil {
			return fmt.Errorf("failed to write: %v", err)
		}
		//bar.Write()
	}
	resFile.Close()
	bar.Finish()
	return
}
