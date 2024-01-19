package file

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

var fileTimer map[string]*time.Timer = make(map[string]*time.Timer, 0)

// var timerFunc *time.Timer := time.AfterFunc(1 * time.Seconds, func (value string, chan channel) { chan <- value })

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
					// fileChan <- event.Name
					if fileTimer[event.Name] == nil {
						fileTimer[event.Name] = time.AfterFunc(
							1*time.Second,
							func() {
								fileChan <- event.Name
								delete(fileTimer, event.Name)
								log.Printf(
									"removed %s from fileTimer map: %v",
									event.Name,
									fileTimer,
								)
							},
						)
					} else { // reset the timer
						fileTimer[event.Name].Reset(1 * time.Second)
					}
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
