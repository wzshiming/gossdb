// +build windows

package winseq

import (
	"syscall"
	"unsafe"
)

// This package only works with Windows.
// Virtual terminal sequences are control character sequences that can control cursor movement,
// color/font mode, and other operations when written to the output stream.
// Sequences may also be received on the input stream in response to an output stream query information sequence or
// as an encoding of user input when the appropriate mode is set.

func init() {
	enableVirtualTerminalProcessing()
}

var (
	kernel32                      = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode            = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode            = kernel32.NewProc("SetConsoleMode")
	hStdout                       = uintptr(syscall.Stdout)
	flagVirtualTerminalProcessing = uintptr(0x0004)
)

func enableVirtualTerminalProcessing() (err error) {
	var mode uintptr
	var ok uintptr
	defer recover()
	ok, _, err = procGetConsoleMode.Call(hStdout, uintptr(unsafe.Pointer(&mode)))
	if ok != 0 {
		err = nil
	}
	if err != nil {
		return err
	}

	if 0 != mode&flagVirtualTerminalProcessing {
		return nil
	}

	ok, _, err = procSetConsoleMode.Call(hStdout, mode|flagVirtualTerminalProcessing)
	if ok != 0 {
		err = nil
	}
	if err != nil {
		return err
	}

	return nil
}
