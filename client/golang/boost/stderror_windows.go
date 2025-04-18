//go:build windows

/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package boost

import (
	"context"
	"github.com/opendatav/mesh/client/golang/log"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

var stderrFd *os.File

func RedirectStderrFile(ctx context.Context, stderr string) {
	if _, err := os.Stat(filepath.Dir(stderr)); nil != err && os.IsNotExist(err) {
		if err = os.MkdirAll(filepath.Dir(stderr), 0755); nil != err {
			log.Error(ctx, err.Error())
			return
		}
	}
	file, err := os.OpenFile(stderr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if nil != err {
		log.Error(ctx, "Redirect stderr with error, %s", err.Error())
		return
	}
	stderrFd = file
	// save global avoid gc recover
	kernel32 := windows.NewLazyDLL("kernel32.dll")
	setStdHandle := kernel32.NewProc("SetStdHandle")
	sh := syscall.STD_ERROR_HANDLE
	if _, _, err = setStdHandle.Call(uintptr(sh), file.Fd()); nil != err {
		log.Error(ctx, "Redirect stderr with error, %s", err.Error())
		return
	}
	// close file describer
	runtime.SetFinalizer(stderrFd, func(fd *os.File) {
		log.Catch(fd.Close())
	})
}
