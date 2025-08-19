package tool

import (
	"github.com/goccy/go-json"
	"github.com/llxxbb/platform-common/def"
	"go.uber.org/zap"
)

func Object2String(property any) (string, *def.CustomError) {
	rtn := ""
	if property != nil {
		marshal, e := json.Marshal(property)
		if e != nil {
			zap.Error(e)
			return "", def.NewCustomError(def.ET_SYS, def.SYS_C, def.SYS_M+e.Error(), nil)
		}
		rtn = string(marshal)
	}
	return rtn, nil
}

func String2Object[T any](strObj string, nilForEmpty bool) (*T, *def.CustomError) {
	if strObj == "" || strObj == "{}" {
		if nilForEmpty {
			return nil, nil
		} else {
			var rtn T
			return &rtn, nil
		}
	}
	var rtn T
	err := json.Unmarshal([]byte(strObj), &rtn)
	if err != nil {
		return nil, def.NewCustomError(def.ET_SYS, def.SYS_C, def.SYS_M+err.Error(), nil)
	}
	return &rtn, nil
}
