package postman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"postman/request/allcollections"
	"postman/request/allenvironments"
	"postman/request/allmocks"
	"postman/request/allmonitors"
	"postman/request/allworkspaces"
	"postman/request/apikeyowner"
	"postman/request/createafork"
	"postman/request/createapi"
	"postman/request/createapiversion"
	"postman/request/createcollection"
	"postman/request/createcollectionfromschema"
	"postman/request/createenvironment"
	"postman/request/createmock"
	"postman/request/createmonitor"
	"postman/request/createrelations"
	"postman/request/createschema"
	"postman/request/createuser"
	"postman/request/createwebhook"
	"postman/request/createworkspace"
	"postman/request/deleteanapi"
	"postman/request/deleteanapiversion"
	"postman/request/deletecollection"
	"postman/request/deleteenvironment"
	"postman/request/deletemock"
	"postman/request/deletemonitor"
	"postman/request/deleteworkspace"
	"postman/request/fetchalluserresource"
	"postman/request/fetchuserresource"
	"postman/request/getallapis"
	"postman/request/getallapiversions"
	"postman/request/getanapiversion"
	"postman/request/getcontracttestrelations"
	"postman/request/getdocumentationrelations"
	"postman/request/getenvironmentrelations"
	"postman/request/getintegrationtestrelations"
	"postman/request/getlinkedrelations"
	"postman/request/getmockserverrelations"
	"postman/request/getmonitorrelations"
	"postman/request/getresourcetypes"
	"postman/request/getschema"
	"postman/request/gettestrelations"
	"postman/request/gettestsuiterelations"
	"postman/request/importexporteddata"
	"postman/request/importexternalapispecification"
	"postman/request/mergeafork"
	"postman/request/publishmock"
	"postman/request/runamonitor"
	"postman/request/schemasecurityvalidation"
	"postman/request/serviceproviderconfig"
	"postman/request/singleapi"
	"postman/request/singlecollection"
	"postman/request/singleenvironment"
	"postman/request/singlemock"
	"postman/request/singlemonitor"
	"postman/request/singleworkspace"
	"postman/request/syncrelationswithschema"
	"postman/request/unpublishmock"
	"postman/request/updateanapi"
	"postman/request/updateanapiversion"
	"postman/request/updatecollection"
	"postman/request/updateenvironment"
	"postman/request/updatemock"
	"postman/request/updatemonitor"
	"postman/request/updateschema"
	"postman/request/updateuserinformation"
	"postman/request/updateuserstate"
	"postman/request/updateworkspace"
)

type Client struct {
	baseUrl string
}

func (client Client) GetAllApis(opts ...getallapis.RequestOption) (*getallapis.Response, error) {
	builder := &getallapis.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.Workspace != nil {
		q.Add("workspace", fmt.Sprint(*builder.Workspace))
	}
	if builder.Since != nil {
		q.Add("since", fmt.Sprint(*builder.Since))
	}
	if builder.Until != nil {
		q.Add("until", fmt.Sprint(*builder.Until))
	}
	if builder.CreatedBy != nil {
		q.Add("createdBy", fmt.Sprint(*builder.CreatedBy))
	}
	if builder.UpdatedBy != nil {
		q.Add("updatedBy", fmt.Sprint(*builder.UpdatedBy))
	}
	if builder.IsPublic != nil {
		q.Add("isPublic", fmt.Sprint(*builder.IsPublic))
	}
	if builder.Name != nil {
		q.Add("name", fmt.Sprint(*builder.Name))
	}
	if builder.Summary != nil {
		q.Add("summary", fmt.Sprint(*builder.Summary))
	}
	if builder.Description != nil {
		q.Add("description", fmt.Sprint(*builder.Description))
	}
	if builder.Sort != nil {
		q.Add("sort", fmt.Sprint(*builder.Sort))
	}
	if builder.Direction != nil {
		q.Add("direction", fmt.Sprint(*builder.Direction))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getallapis.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateApi(opts ...createapi.RequestOption) (*createapi.Response, error) {
	builder := &createapi.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"api": builder.Api})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createapi.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleApi(apiId string) (*singleapi.Response, error) {
	builder := &singleapi.Request{ApiId: apiId}
	url := fmt.Sprintf("%s/apis/%s", client.baseUrl, builder.ApiId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singleapi.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateAnApi(apiId string, opts ...updateanapi.RequestOption) (*updateanapi.Response, error) {
	builder := &updateanapi.Request{ApiId: apiId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s", client.baseUrl, builder.ApiId)
	postBody, _ := json.Marshal(map[string]interface{}{"api": builder.Api})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateanapi.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteAnApi(apiId string) (*deleteanapi.Response, error) {
	builder := &deleteanapi.Request{ApiId: apiId}
	url := fmt.Sprintf("%s/apis/%s", client.baseUrl, builder.ApiId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deleteanapi.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetAllApiVersions(apiId string) (*getallapiversions.Response, error) {
	builder := &getallapiversions.Request{ApiId: apiId}
	url := fmt.Sprintf("%s/apis/%s/versions", client.baseUrl, builder.ApiId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getallapiversions.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateApiVersion(apiId string, opts ...createapiversion.RequestOption) (*createapiversion.Response, error) {
	builder := &createapiversion.Request{ApiId: apiId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions", client.baseUrl, builder.ApiId)
	postBody, _ := json.Marshal(map[string]interface{}{"version": builder.Version})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createapiversion.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetAnApiVersion(apiId string, apiVersionId string) (*getanapiversion.Response, error) {
	builder := &getanapiversion.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getanapiversion.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateAnApiVersion(apiId string, apiVersionId string, opts ...updateanapiversion.RequestOption) (*updateanapiversion.Response, error) {
	builder := &updateanapiversion.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions/%s", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	postBody, _ := json.Marshal(map[string]interface{}{"version": builder.Version})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateanapiversion.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteAnApiVersion(apiId string, apiVersionId string) (*deleteanapiversion.Response, error) {
	builder := &deleteanapiversion.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deleteanapiversion.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetContractTestRelations(apiId string, apiVersionId string) (*getcontracttestrelations.Response, error) {
	builder := &getcontracttestrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/contracttest", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getcontracttestrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetDocumentationRelations(apiId string, apiVersionId string) (*getdocumentationrelations.Response, error) {
	builder := &getdocumentationrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/documentation", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getdocumentationrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetEnvironmentRelations(apiId string, apiVersionId string) (*getenvironmentrelations.Response, error) {
	builder := &getenvironmentrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/environment", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getenvironmentrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetIntegrationTestRelations(apiId string, apiVersionId string) (*getintegrationtestrelations.Response, error) {
	builder := &getintegrationtestrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/integrationtest", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getintegrationtestrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetMockServerRelations(apiId string, apiVersionId string) (*getmockserverrelations.Response, error) {
	builder := &getmockserverrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/mock", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getmockserverrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetMonitorRelations(apiId string, apiVersionId string) (*getmonitorrelations.Response, error) {
	builder := &getmonitorrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/monitor", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getmonitorrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetLinkedRelations(apiId string, apiVersionId string) (*getlinkedrelations.Response, error) {
	builder := &getlinkedrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/relations", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getlinkedrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateRelations(apiId string, apiVersionId string, opts ...createrelations.RequestOption) (*createrelations.Response, error) {
	builder := &createrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/relations", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	postBody, _ := json.Marshal(map[string]interface{}{"documentation": builder.Documentation, "environment": builder.Environment, "mock": builder.Mock, "monitor": builder.Monitor, "test": builder.Test, "contracttest": builder.Contracttest, "testsuite": builder.Testsuite})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateSchema(apiId string, apiVersionId string, opts ...createschema.RequestOption) (*createschema.Response, error) {
	builder := &createschema.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/schemas", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	postBody, _ := json.Marshal(map[string]interface{}{"schema": builder.Schema})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createschema.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetSchema(apiId string, apiVersionId string, schemaId string) (*getschema.Response, error) {
	builder := &getschema.Request{ApiId: apiId, ApiVersionId: apiVersionId, SchemaId: schemaId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/schemas/%s", client.baseUrl, builder.ApiId, builder.ApiVersionId, builder.SchemaId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getschema.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateSchema(apiId string, apiVersionId string, schemaId string, opts ...updateschema.RequestOption) (*updateschema.Response, error) {
	builder := &updateschema.Request{ApiId: apiId, ApiVersionId: apiVersionId, SchemaId: schemaId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/schemas/%s", client.baseUrl, builder.ApiId, builder.ApiVersionId, builder.SchemaId)
	postBody, _ := json.Marshal(map[string]interface{}{"schema": builder.Schema})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateschema.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateCollectionFromSchema(args createcollectionfromschema.Required, opts ...createcollectionfromschema.RequestOption) (*createcollectionfromschema.Response, error) {
	builder := &createcollectionfromschema.Request{ApiId: args.ApiId, ApiVersionId: args.ApiVersionId, SchemaId: args.SchemaId, Name: args.Name, Relations: args.Relations}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/schemas/%s/collections", client.baseUrl, builder.ApiId, builder.ApiVersionId, builder.SchemaId)
	postBody, _ := json.Marshal(map[string]interface{}{"name": builder.Name, "relations": builder.Relations})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createcollectionfromschema.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetTestRelations(apiId string, apiVersionId string) (*gettestrelations.Response, error) {
	builder := &gettestrelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/test", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result gettestrelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetTestSuiteRelations(apiId string, apiVersionId string) (*gettestsuiterelations.Response, error) {
	builder := &gettestsuiterelations.Request{ApiId: apiId, ApiVersionId: apiVersionId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/testsuite", client.baseUrl, builder.ApiId, builder.ApiVersionId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result gettestsuiterelations.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SyncRelationsWithSchema(args syncrelationswithschema.Required) (*syncrelationswithschema.Response, error) {
	builder := &syncrelationswithschema.Request{ApiId: args.ApiId, ApiVersionId: args.ApiVersionId, RelationType: args.RelationType, EntityId: args.EntityId}
	url := fmt.Sprintf("%s/apis/%s/versions/%s/%s/%s/syncWithSchema", client.baseUrl, builder.ApiId, builder.ApiVersionId, builder.RelationType, builder.EntityId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result syncrelationswithschema.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) AllCollections(opts ...allcollections.RequestOption) (*allcollections.Response, error) {
	builder := &allcollections.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/collections", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result allcollections.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateCollection(opts ...createcollection.RequestOption) (*createcollection.Response, error) {
	builder := &createcollection.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/collections", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"collection": builder.Collection})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createcollection.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateAFork(workspace string, collectionUid string, opts ...createafork.RequestOption) (*createafork.Response, error) {
	builder := &createafork.Request{Workspace: workspace, CollectionUid: collectionUid}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/collections/fork/%s", client.baseUrl, builder.CollectionUid)
	postBody, _ := json.Marshal(map[string]interface{}{"label": builder.Label})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("workspace", fmt.Sprint(builder.Workspace))
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createafork.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) MergeAFork(opts ...mergeafork.RequestOption) (*mergeafork.Response, error) {
	builder := &mergeafork.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/collections/merge", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"destination": builder.Destination, "source": builder.Source, "strategy": builder.Strategy})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result mergeafork.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleCollection(collectionUid string) (*singlecollection.Response, error) {
	builder := &singlecollection.Request{CollectionUid: collectionUid}
	url := fmt.Sprintf("%s/collections/%s", client.baseUrl, builder.CollectionUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singlecollection.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateCollection(collectionUid string, opts ...updatecollection.RequestOption) (*updatecollection.Response, error) {
	builder := &updatecollection.Request{CollectionUid: collectionUid}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/collections/%s", client.baseUrl, builder.CollectionUid)
	postBody, _ := json.Marshal(map[string]interface{}{"collection": builder.Collection})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updatecollection.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteCollection(collectionUid string) (*deletecollection.Response, error) {
	builder := &deletecollection.Request{CollectionUid: collectionUid}
	url := fmt.Sprintf("%s/collections/%s", client.baseUrl, builder.CollectionUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deletecollection.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) AllEnvironments(opts ...allenvironments.RequestOption) (*allenvironments.Response, error) {
	builder := &allenvironments.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/environments", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result allenvironments.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateEnvironment(opts ...createenvironment.RequestOption) (*createenvironment.Response, error) {
	builder := &createenvironment.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/environments", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"environment": builder.Environment})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createenvironment.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleEnvironment(environmentUid string) (*singleenvironment.Response, error) {
	builder := &singleenvironment.Request{EnvironmentUid: environmentUid}
	url := fmt.Sprintf("%s/environments/%s", client.baseUrl, builder.EnvironmentUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singleenvironment.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateEnvironment(environmentUid string, opts ...updateenvironment.RequestOption) (*updateenvironment.Response, error) {
	builder := &updateenvironment.Request{EnvironmentUid: environmentUid}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/environments/%s", client.baseUrl, builder.EnvironmentUid)
	postBody, _ := json.Marshal(map[string]interface{}{"environment": builder.Environment})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateenvironment.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteEnvironment(environmentUid string) (*deleteenvironment.Response, error) {
	builder := &deleteenvironment.Request{EnvironmentUid: environmentUid}
	url := fmt.Sprintf("%s/environments/%s", client.baseUrl, builder.EnvironmentUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deleteenvironment.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) ImportExportedData() (*importexporteddata.Response, error) {
	url := fmt.Sprintf("%s/import/exported", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result importexporteddata.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) ImportExternalApiSpecification(body interface{}, opts ...importexternalapispecification.RequestOption) (*importexternalapispecification.Response, error) {
	builder := &importexternalapispecification.Request{Body: body}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/import/openapi", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"body": builder.Body})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result importexternalapispecification.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) ApiKeyOwner() (*apikeyowner.Response, error) {
	url := fmt.Sprintf("%s/me", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result apikeyowner.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) AllMocks() (*allmocks.Response, error) {
	url := fmt.Sprintf("%s/mocks", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result allmocks.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateMock(opts ...createmock.RequestOption) (*createmock.Response, error) {
	builder := &createmock.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/mocks", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"mock": builder.Mock})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createmock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleMock(mockUid string) (*singlemock.Response, error) {
	builder := &singlemock.Request{MockUid: mockUid}
	url := fmt.Sprintf("%s/mocks/%s", client.baseUrl, builder.MockUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singlemock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateMock(mockUid string, opts ...updatemock.RequestOption) (*updatemock.Response, error) {
	builder := &updatemock.Request{MockUid: mockUid}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/mocks/%s", client.baseUrl, builder.MockUid)
	postBody, _ := json.Marshal(map[string]interface{}{"mock": builder.Mock})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updatemock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteMock(mockUid string) (*deletemock.Response, error) {
	builder := &deletemock.Request{MockUid: mockUid}
	url := fmt.Sprintf("%s/mocks/%s", client.baseUrl, builder.MockUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deletemock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) PublishMock(mockUid string) (*publishmock.Response, error) {
	builder := &publishmock.Request{MockUid: mockUid}
	url := fmt.Sprintf("%s/mocks/%s/publish", client.baseUrl, builder.MockUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result publishmock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UnpublishMock(mockUid string) (*unpublishmock.Response, error) {
	builder := &unpublishmock.Request{MockUid: mockUid}
	url := fmt.Sprintf("%s/mocks/%s/unpublish", client.baseUrl, builder.MockUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result unpublishmock.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) AllMonitors() (*allmonitors.Response, error) {
	url := fmt.Sprintf("%s/monitors", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result allmonitors.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateMonitor(opts ...createmonitor.RequestOption) (*createmonitor.Response, error) {
	builder := &createmonitor.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/monitors", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"monitor": builder.Monitor})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createmonitor.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleMonitor(monitorUid string) (*singlemonitor.Response, error) {
	builder := &singlemonitor.Request{MonitorUid: monitorUid}
	url := fmt.Sprintf("%s/monitors/%s", client.baseUrl, builder.MonitorUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singlemonitor.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateMonitor(monitorUid string, opts ...updatemonitor.RequestOption) (*updatemonitor.Response, error) {
	builder := &updatemonitor.Request{MonitorUid: monitorUid}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/monitors/%s", client.baseUrl, builder.MonitorUid)
	postBody, _ := json.Marshal(map[string]interface{}{"monitor": builder.Monitor})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updatemonitor.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteMonitor(monitorUid string) (*deletemonitor.Response, error) {
	builder := &deletemonitor.Request{MonitorUid: monitorUid}
	url := fmt.Sprintf("%s/monitors/%s", client.baseUrl, builder.MonitorUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deletemonitor.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) RunAMonitor(monitorUid string) (*runamonitor.Response, error) {
	builder := &runamonitor.Request{MonitorUid: monitorUid}
	url := fmt.Sprintf("%s/monitors/%s/run", client.baseUrl, builder.MonitorUid)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result runamonitor.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) GetResourceTypes() (*getresourcetypes.Response, error) {
	url := fmt.Sprintf("%s/scim/v2/ResourceTypes", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result getresourcetypes.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) ServiceProviderConfig() (*serviceproviderconfig.Response, error) {
	url := fmt.Sprintf("%s/scim/v2/ServiceProviderConfig", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result serviceproviderconfig.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) FetchAllUserResource(opts ...fetchalluserresource.RequestOption) (*fetchalluserresource.Response, error) {
	builder := &fetchalluserresource.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/scim/v2/Users", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.StartIndex != nil {
		q.Add("startIndex", fmt.Sprint(*builder.StartIndex))
	}
	if builder.Count != nil {
		q.Add("count", fmt.Sprint(*builder.Count))
	}
	if builder.Filter != nil {
		q.Add("filter", fmt.Sprint(*builder.Filter))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result fetchalluserresource.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateUser(opts ...createuser.RequestOption) (*createuser.Response, error) {
	builder := &createuser.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/scim/v2/Users", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"schemas": builder.Schemas, "userName": builder.UserName, "active": builder.Active, "externalId": builder.ExternalId, "groups": builder.Groups, "locale": builder.Locale, "name": builder.Name})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createuser.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) FetchUserResource(userId string) (*fetchuserresource.Response, error) {
	builder := &fetchuserresource.Request{UserId: userId}
	url := fmt.Sprintf("%s/scim/v2/Users/%s", client.baseUrl, builder.UserId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result fetchuserresource.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateUserInformation(userId string, opts ...updateuserinformation.RequestOption) (*updateuserinformation.Response, error) {
	builder := &updateuserinformation.Request{UserId: userId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/scim/v2/Users/%s", client.baseUrl, builder.UserId)
	postBody, _ := json.Marshal(map[string]interface{}{"schemas": builder.Schemas, "name": builder.Name})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateuserinformation.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateUserState(userId string, opts ...updateuserstate.RequestOption) (*updateuserstate.Response, error) {
	builder := &updateuserstate.Request{UserId: userId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/scim/v2/Users/%s", client.baseUrl, builder.UserId)
	postBody, _ := json.Marshal(map[string]interface{}{"schemas": builder.Schemas, "Operations": builder.Operations})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PATCH", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateuserstate.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SchemaSecurityValidation(opts ...schemasecurityvalidation.RequestOption) (*schemasecurityvalidation.Response, error) {
	builder := &schemasecurityvalidation.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/security/api-validation", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"schema": builder.Schema})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result schemasecurityvalidation.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateWebhook(opts ...createwebhook.RequestOption) (*createwebhook.Response, error) {
	builder := &createwebhook.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/webhooks", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"webhook": builder.Webhook})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.WorkspaceId != nil {
		q.Add("workspaceId", fmt.Sprint(*builder.WorkspaceId))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createwebhook.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) AllWorkspaces(opts ...allworkspaces.RequestOption) (*allworkspaces.Response, error) {
	builder := &allworkspaces.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/workspaces", client.baseUrl)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if builder.Type != nil {
		q.Add("type", fmt.Sprint(*builder.Type))
	}
	req.URL.RawQuery = q.Encode()
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result allworkspaces.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) CreateWorkspace(opts ...createworkspace.RequestOption) (*createworkspace.Response, error) {
	builder := &createworkspace.Request{}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/workspaces", client.baseUrl)
	postBody, _ := json.Marshal(map[string]interface{}{"workspace": builder.Workspace})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result createworkspace.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) SingleWorkspace(workspaceId string) (*singleworkspace.Response, error) {
	builder := &singleworkspace.Request{WorkspaceId: workspaceId}
	url := fmt.Sprintf("%s/workspaces/%s", client.baseUrl, builder.WorkspaceId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("GET", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result singleworkspace.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) UpdateWorkspace(workspaceId string, opts ...updateworkspace.RequestOption) (*updateworkspace.Response, error) {
	builder := &updateworkspace.Request{WorkspaceId: workspaceId}
	for _, o := range opts {
		o(builder)
	}
	url := fmt.Sprintf("%s/workspaces/%s", client.baseUrl, builder.WorkspaceId)
	postBody, _ := json.Marshal(map[string]interface{}{"workspace": builder.Workspace})
	body1 := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PUT", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result updateworkspace.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (client Client) DeleteWorkspace(workspaceId string) (*deleteworkspace.Response, error) {
	builder := &deleteworkspace.Request{WorkspaceId: workspaceId}
	url := fmt.Sprintf("%s/workspaces/%s", client.baseUrl, builder.WorkspaceId)
	body1 := bytes.NewBuffer(nil)
	req, err := http.NewRequest("DELETE", url, body1)
	if err != nil {
		return nil, err
	}
	httpclient := http.DefaultClient
	res, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result deleteworkspace.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewClientFromEnv() Client {
	baseUrl, exists := os.LookupEnv("POSTMAN_BASE_URL")
	if !exists {
		fmt.Fprintln(os.Stderr, "Environment variable POSTMAN_BASE_URL is not set")
		os.Exit(1)
	}
	return Client{baseUrl: baseUrl}
}

func NewClient(url string) Client {
	return Client{baseUrl: url}
}
