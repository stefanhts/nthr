package files

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"log"
	"os"
)

type FileStructure struct {
	Path string  `json:"path"`
	Root *Folder `json:"root"`
}

type Folder struct {
	Name    string    `json:"name"`
	Files   []File    `json:"files"`
	Folders []*Folder `json:"folders"`
}

type File struct {
	Name string      `json:"name"`
	Info fs.FileInfo `json:"info"`
}

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

//func Initialize(path string) {
//	err := os.Mkdir(path, os.ModePerm)
//	if os.IsExist(err) {
//		// TODO: confirm with user whether they want to continue with preexisting directory
//		log.Println("directory exists already")
//	} else if err != nil {
//		log.Fatal("Could not make directory:", err)
//	}
//	fs := setupFileSystem(path)
//	fmt.Printf("%v", fs.Hash())
//	out, _ := json.Marshal(fs)
//	//fmt.Printf("%v", out)
//	f2 := &FileStructure{}
//	json.Unmarshal(out, f2)
//	f2.Root.Name = "testDir2"
//	f2.display()
//	f2.Write()
//}

func GetFileStructure(path string) *FileStructure {
	fileStructure := &FileStructure{
		Root: &Folder{
			Files:   []File{},
			Folders: []*Folder{},
		},
	}
	fileStructure.Path = path
	fileStructure.Root.fill(path)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fileStructure.display()
	return fileStructure
}

func (f *FileStructure) Write() error {
	err := os.Mkdir(f.Root.Name, os.ModePerm)
	if os.IsExist(err) {
		// TODO: confirm with user whether they want to continue with preexisting directory
		log.Println("directory exists already")
	} else if err != nil {
		log.Fatal("Could not make directory:", err)
		return err
	}
	return f.Root.write(f.Root.Name)
}

func (f *Folder) write(prefix string) error {
	if len(prefix) > 0 {
		prefix = prefix + "/"
	}
	for _, folder := range f.Folders {
		err := os.Mkdir(prefix+folder.Name, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			log.Fatal("directory creation error: ", err.Error())
		}
		err = folder.write(prefix + folder.Name)
		if err != nil {
			log.Fatal(errors.Wrap(err, "Recursive writing error"))
		}
	}
	for _, file := range f.Files {
		_, err := os.Create(prefix + file.Name)
		if err != nil {
			log.Fatal("file writing error: ", err.Error())
		}
	}
	return nil
}

func (f *Folder) fill(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Error parsing dir", path)
	}

	for _, file := range files {
		if file.IsDir() {
			thisFolder := &Folder{
				Name:    file.Name(),
				Files:   []File{},
				Folders: []*Folder{},
			}
			thisFolder.fill(path + "/" + file.Name())
			f.Folders = append(f.Folders, thisFolder)
		} else {
			info, _ := file.Info()
			f.Files = append(f.Files, File{
				Name: file.Name(),
				Info: info,
			})
		}
	}
	return nil
}

func (f *FileStructure) Display() {
	fmt.Printf("Filestructure: \n-%v/\n", f.Path)
	f.Root.display(1)
}

func (f *Folder) display(indent int) {
	ind := ""
	for i := 0; i < indent; i++ {
		ind += "    "
	}
	for _, folder := range f.Folders {
		fmt.Printf("%s-%s/\n", ind, folder.Name)
		folder.display(indent + 1)
	}
	for _, file := range f.Files {
		fmt.Printf("%s-%s\n", ind, file.Name)
	}
}

func (f *FileStructure) Stringify() string {
	out := "" //fmt.Sprintf("%v", f.Path)
	out += f.Root.stringify("")
	return out
}

func (f *Folder) stringify(prefix string) string {
    prefix = prefix + "/" + f.Name
	out := "\n" + prefix
	for _, folder := range f.Folders {
		//out += fmt.Sprintf("\n%s%s/", prefix, folder.Name)
		out += folder.stringify(prefix)
	}
	for _, file := range f.Files {
		out += file.stringify(prefix)
	}
	return out
}

func (f *File) stringify(prefix string) string {
    return fmt.Sprintf(
        "\n%s/%s : %d : %d",
        prefix,
        f.Name,
        f.Info.Size(),
        f.Info.ModTime().Unix(),
    )
}

func (f *Folder) Hash() string {
	return Hash(f.stringify(""))
}

func (f *File) Hash() string {
	return Hash(f.stringify(""))
}

func (f *FileStructure) Hash() string {
	return Hash(f.Stringify())
}
