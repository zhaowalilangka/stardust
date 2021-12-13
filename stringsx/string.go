package stringsx

import (
    "fmt"
)

func Hello(name string, lang string) (string, error) {
    switch lang {
    case "en":
        return fmt.Sprintf("hello %s", name), nil
    case "zh":
        return fmt.Sprintf("你好 %s", name), nil
    default:
        return "", fmt.Errorf("unknow language ")
    }
}
