package util

import (
    "qqbot-reconstruction/internal/pkg/log"
    "runtime/debug"
)

func SafeGo(fn func()) {
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				log.Error("err recovered: %+v", e)
				log.Error("%s", debug.Stack())
			}
		}()
		fn()
	}()
}