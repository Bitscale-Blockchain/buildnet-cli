package app

import (
	"bitscale/buildnet/lib/build"
	"bitscale/buildnet/lib/utils"
	"log"
)

func InitializeTargetDirectoryTask(context *build.BuildContext) error {
	environment := context.Environment
	log.Printf("Creating target directory %s", context.Environment.WorkingDir)
	utils.DeleteAndCreateDirectory(environment.TargetDir)
	return nil
}
