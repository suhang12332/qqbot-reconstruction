package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"math"
	"sync"
	"time"
)

func WatchFile(path string, handler func(e fsnotify.Event)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(fmt.Errorf("无法监听配置文件： %s", path))
	}
	defer watcher.Close()

	go watchLoop(watcher, handler)

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan interface{})
}

func watchLoop(w *fsnotify.Watcher, handler func(e fsnotify.Event)) {
	var (
		waitFor = 100 * time.Millisecond
		mu      sync.Mutex
		timers  = make(map[string]*time.Timer)
		action  = func(e fsnotify.Event, handler func(e fsnotify.Event)) {
			handler(e)

			mu.Lock()
			delete(timers, e.Name)
			mu.Unlock()
		}
	)

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				mu.Lock()
				t, ok := timers[event.Name]
				mu.Unlock()

				if !ok {
					t = time.AfterFunc(math.MaxInt64, func() { action(event, handler) })
					t.Stop()

					mu.Lock()
					timers[event.Name] = t
					mu.Unlock()
				}

				t.Reset(waitFor)
			}
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}
