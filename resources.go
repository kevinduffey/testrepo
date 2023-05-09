package sdk

import resources "github.com/kevinduffey/testrepo/resources"

type TestAPI struct {
	resources.AuthorizationsResources
}

func NewTestAPI() *TestAPI {
	return &TestAPI{}
}
