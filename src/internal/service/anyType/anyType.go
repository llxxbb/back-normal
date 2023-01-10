package anyType

import (
	"fmt"
	"reflect"
)

// 空接口示例，下面的 any 等价于 interface{}
func PrintType(v any) string {
	_, ok := v.(string)
	if ok == true {
		return fmt.Sprintf("type: string, value: %v", v)
	} else {
		t := reflect.TypeOf(v)
		return fmt.Sprintf("name: %s, kind: %s, string: %s", t.Name(), t.Kind(), t.String())
	}
}
