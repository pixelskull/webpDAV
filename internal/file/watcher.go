package file

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func SetupFileWatcher() {
	dir := "./data"
	newestFileChan := make(chan string)
	go startFileWatcher(dir, newestFileChan)

	for newestFile := range newestFileChan {
		log.Println("Newer file was found: ", newestFile)

		ending, err := checkImageType(newestFile)
		if err != nil {
			log.Println("wrong file type found")
		} else {
			switch ending {
			case "jpg":
				encodeJpg(newestFile)
			default:
				encodePng(newestFile)
			}
		}
	}
}

func startFileWatcher(path string, fileChan chan<- string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					log.Println("New file added: ", event.Name)
				}
				if event.Has(fsnotify.Write) {
					log.Println("File was modified: ", event.Name)
					fileChan <- event.Name
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("ERROR: ", err)
			}
		}
	}()

	watcherErr := watcher.Add(path)
	if watcherErr != nil {
		log.Fatal(watcherErr)
	}
	// prevent this function from ending
	<-make(chan struct{})
}
