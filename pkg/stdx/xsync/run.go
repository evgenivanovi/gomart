package xsync

import (
	"fmt"
	"runtime/debug"
	"sync"

	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
)

/* __________________________________________________ */

func RUN(task func()) {
	doRUN(task)
}

func doRUN(task func()) {
	if task != nil {
		var action = func(task func()) {
			defer recoverer()
			task()
		}
		action(task)
	}
}

/* __________________________________________________ */

func GO(task func()) {
	go doGO(task)
}

func doGO(task func()) {
	if task != nil {
		var action = func(task func()) {
			defer recoverer()
			task()
		}
		action(task)
	}
}

/* __________________________________________________ */

func GOWG(task func(), wgs ...*sync.WaitGroup) {
	for _, wg := range wgs {
		wg.Add(1)
	}
	go doGOWG(task, wgs...)
}

func doGOWG(task func(), wgs ...*sync.WaitGroup) {

	done := func() {
		for _, wg := range wgs {
			wg.Done()
		}
	}

	defer done()

	if task != nil {
		var action = func(task func()) {
			defer recoverer()
			task()
		}
		action(task)
	}

}

/* __________________________________________________ */

func recoverer() {
	if result := recover(); result != nil {
		err := fmt.Errorf("xsync: %v", result)
		slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
	}
}

/* __________________________________________________ */
