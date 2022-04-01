package listfile

import (
	"context"
	"os"
	"path/filepath"
	"time"
)

type TargetFile struct {
	Path string
	Name string
}

type FileList map[string]TargetFile

type FileInfo interface {
	os.FileInfo
	Path() string
}

type fileInfo struct {
	os.FileInfo
	path string
}

func (fi fileInfo) Path() string {
	return fi.path
}

//Ограничить глубину поиска заданым числом, по SIGUSR2 увеличить глубину поиска на +2
func ListDirectory(ctx context.Context, dir string, deep int) ([]FileInfo, error) {
	//log.Println(dir)
	vd := ctx.Value("deep").(int)
	if deep > vd {
		return nil, nil
	}

	select {
	case <-ctx.Done():
		return nil, nil
	default:
		//По SIGUSR1 вывести текущую директорию и текущую глубину поиска
		time.Sleep(time.Millisecond * 100)
		var result []FileInfo
		res, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, entry := range res {
			path := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				child, err := ListDirectory(ctx, path, deep+1) //Дополнительно: вынести в горутину
				if err != nil {
					return nil, err
				}
				result = append(result, child...)
			} else {

				info, err := entry.Info()
				if err != nil {
					return nil, err
				}

				result = append(result, fileInfo{info, path})
			}
		}
		return result, nil
	}
}

func FindFiles(ctx context.Context, ext string) (FileList, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	files, err := ListDirectory(ctx, wd, 0)
	if err != nil {
		return nil, err
	}
	fl := make(FileList, len(files))
	for _, file := range files {

		if filepath.Ext(file.Name()) == ext {

			fl[file.Path()] = TargetFile{
				Name: file.Name(),
				Path: file.Path(),
			}
		}
	}
	return fl, nil
}
