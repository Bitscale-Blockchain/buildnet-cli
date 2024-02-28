package ignite

import (
	"bitscale/buildnet/lib/app"
	"bitscale/buildnet/lib/build"
	"bitscale/buildnet/lib/utils"
	"fmt"
)

const SCAFFOLD_CHAIN_CMD = "ignite scaffold chain"
const SCAFFOLD_MODULE_CMD = "ignite scaffold module"
const SCAFFOLD_TYPE_CMD = "ignite scaffold type"
const SCAFFOLD_MESSAGE_CMD = "ignite scaffold message"
const SCAFFOLD_QUERY_CMD = "ignite scaffold query"

func ScaffoldIgniteProjectTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	cmdString := fmt.Sprintf("%s %s --no-module", SCAFFOLD_CHAIN_CMD, app.Name)
	utils.ShellExecute(cmdString, context.Environment.TargetDir)
	return nil
}

func ScaffoldIgniteModulesTask(context *build.BuildContext) error {
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	for _, module := range app.Modules {
		cmdString := fmt.Sprintf("%s %s", SCAFFOLD_MODULE_CMD, module.Name)

		if module.IsIBCEnabled {
			cmdString = fmt.Sprintf("%s --ibc", cmdString)
		}
		// Process the module dependencies
		if len(module.Dependencies) > 0 {

			cmdString = fmt.Sprintf("%s --dep", cmdString)

			for index, dependency := range module.Dependencies {
				if index == 0 {
					cmdString = fmt.Sprintf("%s %s", cmdString, dependency)
				} else {
					cmdString = fmt.Sprintf("%s,%s", cmdString, dependency)
				}
			}
		}
		targetDir := context.Environment.ProjectDir
		utils.ShellExecute(cmdString, targetDir)
		utils.GitCommitChanges(fmt.Sprintf("Built module %s", module.Name), targetDir)
	}
	return nil
}

func ScaffoldIgniteTypesTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}

	targetDir := context.Environment.ProjectDir

	for _, module := range app.Modules {

		for _, entity := range module.Entities {
			listOfField := utils.GetListOfFields(entity.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_TYPE_CMD, entity.Name, listOfField, module.Name)

			fmt.Printf("Scaffolding type %s in module %s, with command %s", entity.Name, module.Name, buildCmd)
			utils.ShellExecute(buildCmd, targetDir)

			commitMsg := fmt.Sprintf(
				"Scaffolded type %s in module %s", entity.Name, module.Name)

			utils.GitCommitChanges(commitMsg, targetDir)

		}
	}
	return nil
}

func ScaffoldIgniteMessagesTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	targetDir := context.Environment.ProjectDir

	for _, module := range app.Modules {

		for _, message := range module.Messages {
			listOfField := utils.GetListOfFields(message.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_MESSAGE_CMD, message.Name, listOfField, module.Name)

			fmt.Printf("Scaffolding message %s in module %s, with command %s", message.Name, module.Name, buildCmd)
			utils.ShellExecute(buildCmd, targetDir)

			commitMsg := fmt.Sprintf(
				"Scaffolded message %s in module %s", message.Name, module.Name)

			utils.GitCommitChanges(commitMsg, targetDir)

		}
	}
	return nil
}

func ScaffoldIgniteQueriesTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	targetDir := context.Environment.ProjectDir

	for _, module := range app.Modules {

		for _, query := range module.Queries {
			listOfField := utils.GetListOfFields(query.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_QUERY_CMD, query.Name, listOfField, module.Name)

			fmt.Printf("Scaffolding query %s in module %s, with command %s", query.Name, module.Name, buildCmd)
			utils.ShellExecute(buildCmd, targetDir)

			commitMsg := fmt.Sprintf(
				"Scaffolded query %s in module %s", query.Name, module.Name)

			utils.GitCommitChanges(commitMsg, targetDir)

		}
	}
	return nil
}
