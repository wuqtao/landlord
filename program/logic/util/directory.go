package util

import (
	"os"
	"runtime"
)

var OsType = runtime.GOOS

func GetConfigFilePath() string{
	if OsType == "windows"{
		return getProjectPath()+"\\config\\config.toml"
	}else{
		return getProjectPath()+"/config/config.toml"
	}
}

func getProjectPath() string{
	path,_ := os.Getwd()
	return path
}
