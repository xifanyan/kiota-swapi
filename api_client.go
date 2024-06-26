package kiotaswapi

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i2258a36aee3c6e6a0fedc2b744b86c3fc1c0890b1dce696a831d68f818b9db00 "github.com/xifanyan/kiota-swapi/jobs"
    i63007c7783dbecd1b34fff4d17f3c08d1a55a13edf41b60e9921736d425881c7 "github.com/xifanyan/kiota-swapi/projects"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic16b95d508cfc853112fe2824484a6ada4960dc37085ba8c3908bd6a4fcf1175 "github.com/xifanyan/kiota-swapi/logout"
    id8c459c4abeb71247047f7e0b40fbd3ed230f9dc29878bce234921e9c40facc0 "github.com/xifanyan/kiota-swapi/login"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://vm-rhauswirth2.otxlab.net:8443/searchWebApi")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Jobs the jobs property
// returns a *JobsRequestBuilder when successful
func (m *ApiClient) Jobs()(*i2258a36aee3c6e6a0fedc2b744b86c3fc1c0890b1dce696a831d68f818b9db00.JobsRequestBuilder) {
    return i2258a36aee3c6e6a0fedc2b744b86c3fc1c0890b1dce696a831d68f818b9db00.NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Login the login property
// returns a *LoginRequestBuilder when successful
func (m *ApiClient) Login()(*id8c459c4abeb71247047f7e0b40fbd3ed230f9dc29878bce234921e9c40facc0.LoginRequestBuilder) {
    return id8c459c4abeb71247047f7e0b40fbd3ed230f9dc29878bce234921e9c40facc0.NewLoginRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logout the logout property
// returns a *LogoutRequestBuilder when successful
func (m *ApiClient) Logout()(*ic16b95d508cfc853112fe2824484a6ada4960dc37085ba8c3908bd6a4fcf1175.LogoutRequestBuilder) {
    return ic16b95d508cfc853112fe2824484a6ada4960dc37085ba8c3908bd6a4fcf1175.NewLogoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ProjectsRequestBuilder when successful
func (m *ApiClient) Projects()(*i63007c7783dbecd1b34fff4d17f3c08d1a55a13edf41b60e9921736d425881c7.ProjectsRequestBuilder) {
    return i63007c7783dbecd1b34fff4d17f3c08d1a55a13edf41b60e9921736d425881c7.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
