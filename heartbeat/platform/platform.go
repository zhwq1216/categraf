// This file is licensed under the MIT License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright © 2015 Kentaro Kuribayashi <kentarok@gmail.com>
// Copyright 2014-present Datadog, Inc.

//go:build linux || darwin
// +build linux darwin

package platform

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// GetArchInfo() returns basic host architecture information
func GetArchInfo() (archInfo map[string]interface{}, err error) {
	archInfo = make(map[string]interface{})

	out, err := exec.Command("uname", unameOptions...).Output()
	if err != nil {
		return nil, err
	}
	line := fmt.Sprintf("%s", out)
	values := regexp.MustCompile(" +").Split(line, 7)
	updateArchInfo(archInfo, values)

	out, err = exec.Command("uname", "-v").Output()
	if err != nil {
		return nil, err
	}
	archInfo["kernel_version"] = strings.Trim(string(out), "\n")

	return
}
