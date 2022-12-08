package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "day07",
		Short: "Day 7",
		RunE:  execute,
	}
}

func execute(cmd *cobra.Command, args []string) error {
	filesystem := parseFilesystem("day07/input")

	filesystem.CalculateSize()

	filesystemTotal := 70000000
	spaceRequired := 30000000
	totalUsedSpace := filesystem.Size
	unusedSpace := filesystemTotal - totalUsedSpace
	spaceToFree := spaceRequired - unusedSpace

	dirs := filesystem.GetDirectoriesWithMaxSizeOf(100000)

	acc := 0

	for _, dir := range dirs {
		acc += dir.Size
	}

	// part 1
	fmt.Println(acc)

	dirToDelete := 0

	for _, dir := range filesystem.GetFlattenedDirectories() {
		if dirToDelete == 0 {
			dirToDelete = dir.Size
		}

		if dir.Size >= spaceToFree && dir.Size < dirToDelete {
			dirToDelete = dir.Size
		}
	}

	// part 2
	fmt.Println(dirToDelete)

	return nil
}

type File struct {
	Path   string
	Parent *Directory
	Size   int
}

type Directory struct {
	Path        string
	Parent      *Directory
	Files       []*File
	Directories []*Directory
	Size        int
}

func (d *Directory) GetFlattenedDirectories() []*Directory {
	dirs := []*Directory{}

	for _, dir := range d.Directories {
		dirs = append(dirs, dir)

		subDirs := dir.GetFlattenedDirectories()
		dirs = append(dirs, subDirs...)
	}

	return dirs
}

func (d *Directory) GetDirectoriesWithMaxSizeOf(size int) []*Directory {
	dirs := []*Directory{}

	for _, dir := range d.Directories {
		if dir.Size <= size {
			dirs = append(dirs, dir)
		}

		subDirs := dir.GetDirectoriesWithMaxSizeOf(size)

		dirs = append(dirs, subDirs...)
	}

	return dirs
}

func (d *Directory) CalculateSize() {
	for _, file := range d.Files {
		d.Size += file.Size
	}

	for _, dir := range d.Directories {
		dir.CalculateSize()
		d.Size += dir.Size
	}
}

func (d *Directory) NewDirectory(path string) *Directory {
	if string(path[0]) == " " {
		panic("a)")
	}
	for _, dir := range d.Directories {
		if dir.Path == path {
			return dir
		}
	}

	newDir := &Directory{
		Path:   path,
		Parent: d,
	}

	d.Directories = append(d.Directories, newDir)

	return newDir
}

func (d *Directory) NewFile(path string, size int) *File {
	newFile := &File{
		Path:   path,
		Parent: d,
		Size:   size,
	}

	d.Files = append(d.Files, newFile)

	return newFile
}

func changeDirectory(line string) (path string, change bool, back bool) {
	if string(line[0:4]) != "$ cd" {
		return "", false, false
	}

	p := string(line[5:])

	return p, true, p == ".."
}

func isCommand(line string) bool {
	return string(line[0]) == "$"
}

func isDir(line string) (bool, string) {
	return string(line[0:3]) == "dir", string(line[4:])
}

func isFile(line string) (bool, string, int) {
	isDir, _ := isDir(line)

	if isCommand(line) || isDir {
		return false, "", 0
	}

	file := strings.Split(line, " ")

	size, err := strconv.Atoi(file[0])

	if err != nil {
		// fmt.Println(file[0])
		panic("file size is not int")
	}

	return true, file[1], size
}

func parseFilesystem(path string) *Directory {
	file, err := os.Open(path)

	if err != nil {
		// return err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	root := &Directory{
		Path: "/",
	}

	pwd := root

	for scanner.Scan() {
		line := scanner.Text()

		if isDir, dir := isDir(line); isDir {
			pwd.NewDirectory(dir)
			continue
		}

		if isFile, file, size := isFile(line); isFile {
			pwd.NewFile(file, size)
			continue
		}

		path, shouldChange, goBack := changeDirectory(line)

		if path == "/" {
			pwd = root
			continue
		}

		if goBack {
			pwd = pwd.Parent
			continue
		}

		if shouldChange {
			pwd = pwd.NewDirectory(path)
		}
	}

	return root
}
