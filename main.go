package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"

	"github.com/bmatcuk/doublestar"
	"github.com/jstutters/dicom"
)

func readSeriesNumberAndDescription(path string) (string, string, error) {
	p, err := dicom.NewParserFromFile(path, nil)
	if err != nil {
		return "", "", err
	}
	opts := dicom.ParseOptions{DropPixelData: true}
	dataset, err := p.Parse(opts)
	if dataset == nil || err != nil {
		return "", "", err
	}
	seriesNumberElement, err := dataset.FindElementByName("SeriesNumber")
	if err != nil {
		return "", "", err
	}
	seriesNumber, err := seriesNumberElement.GetString()
	if err != nil {
		return "", "", err
	}
	seriesDescriptionElement, err := dataset.FindElementByName("SeriesDescription")
	if err != nil {
		return "", "", err
	}
	seriesDescription, err := seriesDescriptionElement.GetString()
	if err != nil {
		return "", "", err
	}
	return seriesNumber, seriesDescription, nil
}

func decideSearchPath() (string, error) {
	flag.Parse()
	searchPath := flag.Arg(0)
	if searchPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		searchPath = cwd
	}
	searchPath += "/**"
	searchPath = path.Clean(searchPath)
	return searchPath, nil
}

func scanFiles(files []string) map[int]string {
	series := make(map[int]string)
	for _, f := range files {
		seriesNumberStr, seriesDescription, err := readSeriesNumberAndDescription(f)
		if err != nil {
			continue
		}
		seriesNumber, err := strconv.Atoi(seriesNumberStr)
		if err != nil {
			continue
		}
		series[seriesNumber] = seriesDescription
	}
	return series
}

func printSeries(series map[int]string) {
	var keys []int
	for k := range series {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		desc := series[k]
		fmt.Println(k, "\t", desc)
	}
}

func main() {
	// Search is working directory or first argument
	searchPath, err := decideSearchPath()
	if err != nil {
		panic(err)
	}

	// Do a recursive glob for all files in search path
	files, err := doublestar.Glob(searchPath)
	if err != nil {
		panic(err)
	}

	// Read DICOM information from all files found
	series := scanFiles(files)

	// Print the series numbers and descriptions
	printSeries(series)
}
