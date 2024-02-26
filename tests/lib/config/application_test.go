package config_test

import (
	"bitscale/buildnet/lib/config"
	"testing"
)

func TestApplication(t *testing.T) {
	app := config.Application{
		Name: "TestApp",
		Modules: []config.Module{
			{
				Name:         "TestModule",
				IsIBCEnabled: true,
				Dependencies: []string{"Dep1", "Dep2"},
				Entities: []config.Entity{
					{
						Name: "Entity1",
						Fields: []config.Field{
							{Name: "Field1", Type: "Type1"},
							{Name: "Field2", Type: "Type2"},
						},
					},
				},
				Messages: []config.Message{
					{
						Name: "Message1",
						Fields: []config.Field{
							{Name: "Field1", Type: "Type1"},
							{Name: "Field2", Type: "Type2"},
						},
					},
				},
				Queries: []config.Query{
					{
						Name: "Query1",
						Fields: []config.Field{
							{Name: "Field1", Type: "Type1"},
							{Name: "Field2", Type: "Type2"},
						},
					},
				},
			},
		},
	}

	// Verify that the application name is correct.
	if app.Name != "TestApp" {
		t.Errorf("expected application name to be 'TestApp', got '%s'", app.Name)
	}

	// Verify that the module name is correct.
	if app.Modules[0].Name != "TestModule" {
		t.Errorf("expected module name to be 'TestModule', got '%s'", app.Modules[0].Name)
	}

	// Verify that the module has the correct number of dependencies.
	if len(app.Modules[0].Dependencies) != 2 {
		t.Errorf("expected module to have 2 dependencies, got %d", len(app.Modules[0].Dependencies))
	}

	// Verify that the module has the correct number of entities.
	if len(app.Modules[0].Entities) != 1 {
		t.Errorf("expected module to have 1 entity, got %d", len(app.Modules[0].Entities))
	}

	// Verify that the entity name is correct.
	if app.Modules[0].Entities[0].Name != "Entity1" {
		t.Errorf("expected entity name to be 'Entity1', got '%s'", app.Modules[0].Entities[0].Name)
	}

	// Verify that the entity has the correct number of fields.
	if len(app.Modules[0].Entities[0].Fields) != 2 {
		t.Errorf("expected entity to have 2 fields, got %d", len(app.Modules[0].Entities[0].Fields))
	}

	// Verify that the first field name is correct.
	if app.Modules[0].Entities[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", app.Modules[0].Entities[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if app.Modules[0].Entities[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", app.Modules[0].Entities[0].Fields[0].Type)
	}

	// Verify that the module has the correct number of messages.
	if len(app.Modules[0].Messages) != 1 {
		t.Errorf("expected module to have 1 message, got %d", len(app.Modules[0].Messages))
	}

	// Verify that the message name is correct.
	if app.Modules[0].Messages[0].Name != "Message1" {
		t.Errorf("expected message name to be 'Message1', got '%s'", app.Modules[0].Messages[0].Name)
	}

	// Verify that the message has the correct number of fields.
	if len(app.Modules[0].Messages[0].Fields) != 2 {
		t.Errorf("expected message to have 2 fields, got %d", len(app.Modules[0].Messages[0].Fields))
	}

	// Verify that the first field name is correct.
	if app.Modules[0].Messages[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", app.Modules[0].Messages[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if app.Modules[0].Messages[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", app.Modules[0].Messages[0].Fields[0].Type)
	}

	// Verify that the module has the correct number of queries.
	if len(app.Modules[0].Queries) != 1 {
		t.Errorf("expected module to have 1 query, got %d", len(app.Modules[0].Queries))
	}

	// Verify that the query name is correct.
	if app.Modules[0].Queries[0].Name != "Query1" {
		t.Errorf("expected query name to be 'Query1', got '%s'", app.Modules[0].Queries[0].Name)
	}

	// Verify that the query has the correct number of fields.
	if len(app.Modules[0].Queries[0].Fields) != 2 {
		t.Errorf("expected query to have 2 fields, got %d", len(app.Modules[0].Queries[0].Fields))
	}

	// Verify that the first field name is correct.
	if app.Modules[0].Queries[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", app.Modules[0].Queries[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if app.Modules[0].Queries[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", app.Modules[0].Queries[0].Fields[0].Type)
	}
}

func TestModule(t *testing.T) {
	mod := config.Module{
		Name:         "TestModule",
		IsIBCEnabled: true,
		Dependencies: []string{"Dep1", "Dep2"},
		Entities: []config.Entity{
			{
				Name: "Entity1",
				Fields: []config.Field{
					{Name: "Field1", Type: "Type1"},
					{Name: "Field2", Type: "Type2"},
				},
			},
		},
		Messages: []config.Message{
			{
				Name: "Message1",
				Fields: []config.Field{
					{Name: "Field1", Type: "Type1"},
					{Name: "Field2", Type: "Type2"},
				},
			},
		},
		Queries: []config.Query{
			{
				Name: "Query1",
				Fields: []config.Field{
					{Name: "Field1", Type: "Type1"},
					{Name: "Field2", Type: "Type2"},
				},
			},
		},
	}

	// Verify that the module name is correct.
	if mod.Name != "TestModule" {
		t.Errorf("expected module name to be 'TestModule', got '%s'", mod.Name)
	}

	// Verify that the module has the correct number of dependencies.
	if len(mod.Dependencies) != 2 {
		t.Errorf("expected module to have 2 dependencies, got %d", len(mod.Dependencies))
	}

	// Verify that the module has the correct number of entities.
	if len(mod.Entities) != 1 {
		t.Errorf("expected module to have 1 entity, got %d", len(mod.Entities))
	}

	// Verify that the entity name is correct.
	if mod.Entities[0].Name != "Entity1" {
		t.Errorf("expected entity name to be 'Entity1', got '%s'", mod.Entities[0].Name)
	}

	// Verify that the entity has the correct number of fields.
	if len(mod.Entities[0].Fields) != 2 {
		t.Errorf("expected entity to have 2 fields, got %d", len(mod.Entities[0].Fields))
	}

	// Verify that the first field name is correct.
	if mod.Entities[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", mod.Entities[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if mod.Entities[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", mod.Entities[0].Fields[0].Type)
	}

	// Verify that the module has the correct number of messages.
	if len(mod.Messages) != 1 {
		t.Errorf("expected module to have 1 message, got %d", len(mod.Messages))
	}

	// Verify that the message name is correct.
	if mod.Messages[0].Name != "Message1" {
		t.Errorf("expected message name to be 'Message1', got '%s'", mod.Messages[0].Name)
	}

	// Verify that the message has the correct number of fields.
	if len(mod.Messages[0].Fields) != 2 {
		t.Errorf("expected message to have 2 fields, got %d", len(mod.Messages[0].Fields))
	}

	// Verify that the first field name is correct.
	if mod.Messages[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", mod.Messages[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if mod.Messages[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", mod.Messages[0].Fields[0].Type)
	}

	// Verify that the module has the correct number of queries.
	if len(mod.Queries) != 1 {
		t.Errorf("expected module to have 1 query, got %d", len(mod.Queries))
	}

	// Verify that the query name is correct.
	if mod.Queries[0].Name != "Query1" {
		t.Errorf("expected query name to be 'Query1', got '%s'", mod.Queries[0].Name)
	}

	// Verify that the query has the correct number of fields.
	if len(mod.Queries[0].Fields) != 2 {
		t.Errorf("expected query to have 2 fields, got %d", len(mod.Queries[0].Fields))
	}

	// Verify that the first field name is correct.
	if mod.Queries[0].Fields[0].Name != "Field1" {
		t.Errorf("expected first field name to be 'Field1', got '%s'", mod.Queries[0].Fields[0].Name)
	}

	// Verify that the first field type is correct.
	if mod.Queries[0].Fields[0].Type != "Type1" {
		t.Errorf("expected first field type to be 'Type1', got '%s'", mod.Queries[0].Fields[0].Type)
	}
}
