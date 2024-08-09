package utils

import (
	"path"
	"runtime"
)

func RootDirectory () string {
    _, b, _, _ := runtime.Caller(0)
    d1 := path.Join(path.Dir(b), "../..")    

    return d1
}

func CurrentDirectory () string {
    _, b, _, _ := runtime.Caller(0)
    d1 := path.Join(path.Dir(b))    

    return d1
}