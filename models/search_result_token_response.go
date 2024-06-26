package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SearchResultTokenResponse struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date at which this token expires and will be deleted. It has to be touched or used before this timestamp in order to keep it alive. The format is 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Note that a search token is bound to the user session as well: it will be closed if the user session times out or if the users logs out.
    eol *string
    // The number of results found by the search.
    numberResults *int64
    // The status of the response.
    status StatusObjectable
    // The search result token.
    token *string
}
// NewSearchResultTokenResponse instantiates a new SearchResultTokenResponse and sets the default values.
func NewSearchResultTokenResponse()(*SearchResultTokenResponse) {
    m := &SearchResultTokenResponse{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResultTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSearchResultTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResultTokenResponse(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *SearchResultTokenResponse) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SearchResultTokenResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEol gets the eol property value. The date at which this token expires and will be deleted. It has to be touched or used before this timestamp in order to keep it alive. The format is 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Note that a search token is bound to the user session as well: it will be closed if the user session times out or if the users logs out.
// returns a *string when successful
func (m *SearchResultTokenResponse) GetEol()(*string) {
    return m.eol
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SearchResultTokenResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEol(val)
        }
        return nil
    }
    res["numberResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberResults(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatusObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(StatusObjectable))
        }
        return nil
    }
    res["token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToken(val)
        }
        return nil
    }
    return res
}
// GetNumberResults gets the numberResults property value. The number of results found by the search.
// returns a *int64 when successful
func (m *SearchResultTokenResponse) GetNumberResults()(*int64) {
    return m.numberResults
}
// GetStatus gets the status property value. The status of the response.
// returns a StatusObjectable when successful
func (m *SearchResultTokenResponse) GetStatus()(StatusObjectable) {
    return m.status
}
// GetToken gets the token property value. The search result token.
// returns a *string when successful
func (m *SearchResultTokenResponse) GetToken()(*string) {
    return m.token
}
// Serialize serializes information the current object
func (m *SearchResultTokenResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("eol", m.GetEol())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("numberResults", m.GetNumberResults())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("token", m.GetToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResultTokenResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEol sets the eol property value. The date at which this token expires and will be deleted. It has to be touched or used before this timestamp in order to keep it alive. The format is 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Note that a search token is bound to the user session as well: it will be closed if the user session times out or if the users logs out.
func (m *SearchResultTokenResponse) SetEol(value *string)() {
    m.eol = value
}
// SetNumberResults sets the numberResults property value. The number of results found by the search.
func (m *SearchResultTokenResponse) SetNumberResults(value *int64)() {
    m.numberResults = value
}
// SetStatus sets the status property value. The status of the response.
func (m *SearchResultTokenResponse) SetStatus(value StatusObjectable)() {
    m.status = value
}
// SetToken sets the token property value. The search result token.
func (m *SearchResultTokenResponse) SetToken(value *string)() {
    m.token = value
}
type SearchResultTokenResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEol()(*string)
    GetNumberResults()(*int64)
    GetStatus()(StatusObjectable)
    GetToken()(*string)
    SetEol(value *string)()
    SetNumberResults(value *int64)()
    SetStatus(value StatusObjectable)()
    SetToken(value *string)()
}
