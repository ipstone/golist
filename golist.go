// This utility recursively goes through folder to output files info in csv format.

package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
  "strings"
)

func visit(path string, f os.FileInfo, err error) error {
	//fmt.Printf("Visited: %s\n", path)
	//info, serror:=os.Stat(f)
	name, size, mtime := f.Name(), float64(f.Size())/1024.0, f.ModTime()
	f_path := filepath.Dir(path)
	splitted := strings.SplitAfter(name, ".")
	ext :=splitted[len(splitted)-1]
	// mtimeStr := mtime.Format("2012-01-01")
	// fmt.Printf("%s, %s, %.2f, %s, %s\n", name, ext, size, mtimeStr, f_path)

	fmt.Printf("%s, %s, %.2f, %s, %s\n", name, ext, size, mtime, f_path)
	//fmt.Printf("File stats: %s \n",info)
	//fmt.Printf("The type of stats info is: %T\n", info)
  return nil
}

func main() {
	helpMsg := "Usage: golist folderPath. If no folderPath given, using ./ instead."
	flag.Usage=func() {fmt.Println(helpMsg)}	
	flag.Parse()
	root := flag.Arg(0)
	// fmt.Println(root)

	if len(root) == 0 { root="."}
	fmt.Printf("Name,Extension,SizeKb,ModifiedTime,Path\n")
	err := filepath.Walk(root, visit)
	if err==nil {
		// successful run, do nothing!
	} else {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
