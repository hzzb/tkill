package main

import (
	"os"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

var usage = `
NAME
    tkill -- send signal to thread instead of process.

SYNOPSIS
    tkill [-s signal_name] pid tid
    tkill -signal_name 	   pid tid
    tkill -signal_number   pid tid
`

func main() {
	var (
		nmnum, pidArg, tidArg string
		length, pid, tid      int
		sig                   syscall.Signal
		err                   error
	)

	length = len(os.Args)
	if length < 3 {
		goto usage
	}
	length--
	tidArg = os.Args[length]
	length--
	pidArg = os.Args[length]
	if i, e := strconv.ParseInt(pidArg, 10, 64); e != nil {
		goto usage
	} else {
		pid = int(i)
	}
	if i, e := strconv.ParseInt(tidArg, 10, 64); e != nil {
		goto usage
	} else {
		tid = int(i)
	}

	switch length {
	default:
		goto usage
	case 1:
		nmnum = "SIGTERM"
	case 2:
		if os.Args[1][0] != '-' {
			goto usage
		}
		nmnum = os.Args[1][1:]
	case 3:
		if os.Args[1] != "-s" {
			goto usage
		}
		nmnum = os.Args[2]
	}

	if i, e := strconv.ParseInt(nmnum, 10, 64); e != nil {
		sig = unix.SignalNum(nmnum)
	} else {
		sig = syscall.Signal(i)
	}
	if sig != 0 {
		goto signal
	}
usage:
	os.Stderr.WriteString(usage)
	return

signal:
	err = unix.Tgkill(pid, tid, sig)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
}
