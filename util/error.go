package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func CheckErr(err error) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, "Error: ", err)
	logrus.Errorf("Error: %+v", err)
	os.Exit(1)
}
