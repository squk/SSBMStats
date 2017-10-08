package melee

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type DolphinManager struct {
	SelfPort     int
	OpponentPort int
	DolphinPath  string
	MemoryPath   string
	PipesPath    string
	RUNNING      bool
}

func NewDolphinManager() *DolphinManager {
	d := DolphinManager{}
	d.SetPath("/Users/christian/Desktop/FM/DolphinManager.app/Contents/Resources/User")
	d.RUNNING = true

	if d.MemoryPath != "" {
		_ = os.Mkdir(d.MemoryPath, os.ModePerm)
		log.Println("Created MemoryWatcher path")
	}
	if d.PipesPath != "" {
		_ = os.Mkdir(d.PipesPath, os.ModePerm)
		log.Println("Created Pipes path")
	}

	return &d
}

func (d *DolphinManager) StopLoop() {
	d.RUNNING = false
}

func (d *DolphinManager) Init() bool {
	if d.DolphinPath == "" || d.MemoryPath == "" || d.PipesPath == "" {
		return false
	} else {
		err := CopyFile("Locations.txt", filepath.Join(d.MemoryPath, "Locations.txt"))
		if err != nil {
			fmt.Println(err)
		}
	}

	return true
}

func (d *DolphinManager) SetPath(path string) bool {
	ex, _ := FilepathExists(path)

	if ex {
		d.DolphinPath = path
		d.MemoryPath = filepath.Join(path, "MemoryWatcher/")
		d.PipesPath = filepath.Join(path, "Pipes/")

		_ = os.Mkdir(d.MemoryPath, os.ModePerm)
		//fmt.Println("Created MemoryWatcher path")
		_ = os.Mkdir(d.PipesPath, os.ModePerm)
		//fmt.Println("Created Pipes path")
	}
	return ex
}

func FilepathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
