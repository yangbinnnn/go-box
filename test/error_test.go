package test

import (
	"errors"
	"go-box/common"
	"testing"
)

func TestError(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	common.CheckErrorWithIgnore(err1, err1, err2)
}
