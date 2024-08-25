package utils

import (
	"fmt"
	"strings"
)

func ErrorMassage(field string, tag string, param string) string{
    if (tag != "" && param == "" ) {
        return fmt.Sprintf("This %s input must be %s", strings.ToLower(field), tag)
    }
    if (tag == "min" || tag == "max" && param != "" ) {
        return fmt.Sprintf("This %s input %s %s characters", strings.ToLower(field), tag, param)
    }
    return ""
}