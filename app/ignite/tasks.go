package ignite

import (
	"bitscale/buildnet/lib/app"
	"bitscale/buildnet/lib/build"
	"bitscale/buildnet/lib/utils"
	"fmt"
)

const SCAFFOLD_CMD = "ignite scaffold"
const GENERATE_CMD = "ignite generate"
const SCAFFOLD_CHAIN_CMD = "ignite scaffold chain"
const SCAFFOLD_MODULE_CMD = "ignite scaffold module"
const SCAFFOLD_TYPE_CMD = "ignite scaffold type"
const SCAFFOLD_MESSAGE_CMD = "ignite scaffold message"
const SCAFFOLD_QUERY_CMD = "ignite scaffold query"
const SCAFFOLD_VUE_CLIENT_CMD = "ignite scaffold vue"
const SCAFFOLD_TS_CLIENT_CMD = "ignite generate ts-client"
const SCAFFOLD_DENOM_CMD = "ignite scaffold map Denom description:string ticker:string precision:int url:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module "

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
	// Execute the command
	utils.ShellExecute(
		fmt.Sprintf("%s %s", SCAFFOLD_CHAIN_CMD, app.Name), context.GetTargetDir())

	return nil
}

func ScaffoldIgniteTokenFactoryTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	// Execute the command string
	utils.ShellExecute(
		fmt.Sprintf("%s %s", SCAFFOLD_DENOM_CMD, app.Name), context.GetProjectDir())
	// Commit changes
	utils.GitCommitChanges(
		fmt.Sprintf("Scaffolded token factory for module %s", app.Name), context.GetProjectDir())
	return nil
}

func ScaffoldIgniteFrontendTask(context *build.BuildContext) error {
	// Task1 implementation
	// Get current working directory
	config := context.Configuration
	// Type switch to check the type of eventObj.Data
	app, ok := config.Data.(*app.Blockchain)
	if !ok {
		// eventObj.Data is not a string, return an error
		return fmt.Errorf("unexpected app data type: %T", app)
	}
	// Allowed frontend values
	allowedValues := map[string]bool{
		"vue":       true,
		"react":     true,
		"ts-client": true,
	}
	// Ensure array of frontends contains only valid values
	if !utils.ContainsOnlyAllowedValues(app.Frontends, allowedValues) {
		return fmt.Errorf("unexpected value declared as fronten: %s", app.Frontends)
	}

	for _, frontend := range app.Frontends {
		// Set as scaffold command
		cmdString := SCAFFOLD_CMD
		// ...unless ts-client is specified
		if frontend == "ts-client" {
			cmdString = GENERATE_CMD
		}
		utils.ShellExecute(
			fmt.Sprintf("%s %s", cmdString, frontend), context.GetProjectDir())

		utils.GitCommitChanges(
			fmt.Sprintf("Added frontend %s", frontend), context.GetProjectDir())
	}

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
		utils.ShellExecute(cmdString, context.GetProjectDir())

		utils.GitCommitChanges(
			fmt.Sprintf("Built module %s", module.Name), context.GetProjectDir())
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

	for _, module := range app.Modules {

		for _, entity := range module.Entities {
			listOfField := utils.GetListOfFields(entity.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_TYPE_CMD, entity.Name, listOfField, module.Name)

			utils.ShellExecute(buildCmd, context.GetProjectDir())

			utils.GitCommitChanges(fmt.Sprintf(
				"Scaffolded type %s in module %s", entity.Name, module.Name), context.GetProjectDir())

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
	for _, module := range app.Modules {

		for _, message := range module.Messages {
			listOfField := utils.GetListOfFields(message.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_MESSAGE_CMD, message.Name, listOfField, module.Name)

			utils.ShellExecute(buildCmd, context.GetProjectDir())

			utils.GitCommitChanges(fmt.Sprintf(
				"Scaffolded message %s in module %s", message.Name, module.Name), context.GetProjectDir())

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
	for _, module := range app.Modules {

		for _, query := range module.Queries {
			listOfField := utils.GetListOfFields(query.Fields)

			buildCmd := fmt.Sprintf("%s %s %s --module %s",
				SCAFFOLD_QUERY_CMD, query.Name, listOfField, module.Name)

			utils.ShellExecute(buildCmd, context.GetProjectDir())

			utils.GitCommitChanges(fmt.Sprintf(
				"Scaffolded query %s in module %s", query.Name, module.Name), context.GetProjectDir())
		}
	}
	return nil
}
