package util

import (
    "fmt"
    "qqbot-reconstruction/internal/pkg/variable"
    "strings"
)

func In(target string, arr []string) bool {
    for _, element := range arr {
        if target == element {
            return true
        }
    }
    return false
}

func IsStringArraysEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func RemoveRepeatedElement(s []string) []string {
    result := make([]string, 0)
    m := make(map[string]bool) //map的值不重要
    for _, v := range s {
        if _, ok := m[v]; !ok {
            result = append(result, v)
            m[v] = true
        }
    }
    return result
}

func RemoveElement(targets []string, arr []string) []string {
    for _, tgt := range targets {
        for i, v := range arr {
            if v == tgt {
                arr = append(arr[:i], arr[i+1:]...)
            }
        }
    }
    return arr
}

func ParseHelp(scope []string) string {
    if len(scope) == 0 {
        return fmt.Sprintf("[%s,%s]", variable.GROUPMESSGAEZH, variable.PRIVATEMESSAGEZH)
    }
    for key, value := range scope {
        scope[key] = ParseMessageType(value)
    }
    return fmt.Sprintf("[%s]", strings.Join(scope, ","))
}
