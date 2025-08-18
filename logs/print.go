package logs

import (
	"encoding/json"
	"fmt"
	"gitlab.com/package7225033/go-modules/check"
	"gitlab.com/package7225033/go-modules/get-env"
	"runtime"
	"time"
)

var goEnv, goDebugLevel string
var goDebug bool

func init() {
	goEnv = get_env.String("GO_ENV", "development")
	goDebug = get_env.GetBool("DEBUG", false)
	goDebugLevel = get_env.String("DEBUG_LEVEL", "all")
}

func PrettyPrint(data interface{}) {
	if !check.ContainString(goEnv, []string{"production", "prod"}) && (goDebug && check.ContainString(goDebugLevel, []string{"all", "info"})) && data != nil {
		p, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			ErrorHandler(err)
			return
		}
		fmt.Printf("[%s]\n%s \n", get_env.String("ENGINE_NAME", "UNDEFINED"), p)
	}
}

// ErrorHandler for print error message to console with engine name if error != nil
func ErrorHandler(err error) {

	if !check.ContainString(goEnv, []string{"production", "prod"}) && (goDebug && check.ContainString(goDebugLevel, []string{"all", "error"})) && err != nil {
		var (
			_, file, line, _ = runtime.Caller(1)
			loc, _           = time.LoadLocation("Asia/Jakarta")
			timeStr          = time.Now().In(loc).Format("2006/01/02 15:04:05")
		)

		fmt.Printf("[%s] %s\n%s %s:%d\n", get_env.String("ENGINE_NAME", "UNDEFINED"), err.Error(), timeStr, file, line)
	}
}
