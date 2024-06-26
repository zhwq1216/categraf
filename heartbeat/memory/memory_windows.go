// This file is licensed under the MIT License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright © 2015 Kentaro Kuribayashi <kentarok@gmail.com>
// Copyright 2014-present Datadog, Inc.

package memory

import (
	"strconv"
	"syscall"
	"unsafe"
)

type MEMORYSTATUSEX struct {
	dwLength               uint32 // size of this structure
	dwMemoryLoad           uint32 // number 0-100 estimating %age of memory in use
	ulTotalPhys            uint64 // amount of physical memory
	ulAvailPhys            uint64 // amount of physical memory that can be used w/o flush to disk
	ulTotalPageFile        uint64 // current commit limit for system or process
	ulAvailPageFile        uint64 // amount of memory current process can commit
	ulTotalVirtual         uint64 // size of user-mode portion of VA space
	ulAvailVirtual         uint64 // amount of unreserved/uncommitted memory in ulTotalVirtual
	ulAvailExtendedVirtual uint64 // reserved (always zero)
}

func getMemoryInfo() (memoryInfo map[string]string, err error) {
	memoryInfo = make(map[string]string)

	var mod = syscall.NewLazyDLL("kernel32.dll")
	var getMem = mod.NewProc("GlobalMemoryStatusEx")

	var mem_struct MEMORYSTATUSEX

	mem_struct.dwLength = uint32(unsafe.Sizeof(mem_struct))

	status, _, err := getMem.Call(uintptr(unsafe.Pointer(&mem_struct)))
	if status != 0 {
		memoryInfo["total"] = strconv.FormatUint(mem_struct.ulTotalPhys, 10)
		err = nil
	}
	return
}
