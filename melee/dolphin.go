package melee

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type DolphinManager struct {
	SelfPort    int
	DolphinPath string
	MemoryPath  string
	PipesPath   string
	RUNNING     bool
}

func NewDolphinManager() *DolphinManager {
	d := DolphinManager{SelfPort: 1}
	d.SetPath("/Users/christian/Desktop/FM/Dolphin.app/Contents/Resources/User")
	log.Println(d.DolphinPath)
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

func (d *DolphinManager) IncreasePort() {
	d.SelfPort++

	if d.SelfPort > 4 {
		d.SelfPort = 1
	}
}

func (d *DolphinManager) DecreasePort() {
	d.SelfPort--

	if d.SelfPort < 1 {
		d.SelfPort = 4
	}
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
			return false
		}
	}

	return true
}

func (d *DolphinManager) SetPath(path string) bool {
	ex, err := FilepathExists(path)

	if err != nil {
		log.Fatalln(err)
	}

	if ex {
		d.DolphinPath = path
		d.MemoryPath = filepath.Join(path, "MemoryWatcher/")
		d.PipesPath = filepath.Join(path, "Pipes/")

		_ = os.Mkdir(d.MemoryPath, os.ModePerm)
		_ = os.Mkdir(d.PipesPath, os.ModePerm)
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
