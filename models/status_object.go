package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusObject the status of the response.
type StatusObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The status of the backend response if there is one. This can be 'null'. The value 'meaningless search terms' indicates that the query consists only of stop words and the-like and results in HTTP code 200. The value 'all search terms unknown' is similiar, but the server does not know any of the provided terms. The special value 'incomplete search results due to pending synchronization' can only occur in early startup phases of the target server (if synchronizes searchable folder properties asynchronously); it results in HTTP code 200. The value 'child collection invalid' means that a sub-engine is currently unreachable and its results are not part of the result and also results in code 200. The value 'unexpandable wildcard' indicates that a wildcard prefix is too short (it results in HTTP code 400). The value 'expanded query too long' indicates that a wildcard expression resulted in too many terms (which is also code 400). The values here are constants and are subject to api stability requirements. The value 'syntax error' indicates a syntax error in a query. The value 'unsupported expression' indicates that the query as such is valid, but cannot be used in the current context.
    backendStatus *StatusObject_backendStatus
    // A short description of the error. This can be 'null'
    errorMessage *string
    // The http status of the response. The status depends on 'backendStatus': 'unknown error' is mapped to 'BAD REQUEST' and all others are mapped to 'OK' 
    httpStatus *int32
    // Will be 'true' if the request was fully successful.
    successful *bool
}
// NewStatusObject instantiates a new StatusObject and sets the default values.
func NewStatusObject()(*StatusObject) {
    m := &StatusObject{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStatusObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatusObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusObject(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StatusObject) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBackendStatus gets the backendStatus property value. The status of the backend response if there is one. This can be 'null'. The value 'meaningless search terms' indicates that the query consists only of stop words and the-like and results in HTTP code 200. The value 'all search terms unknown' is similiar, but the server does not know any of the provided terms. The special value 'incomplete search results due to pending synchronization' can only occur in early startup phases of the target server (if synchronizes searchable folder properties asynchronously); it results in HTTP code 200. The value 'child collection invalid' means that a sub-engine is currently unreachable and its results are not part of the result and also results in code 200. The value 'unexpandable wildcard' indicates that a wildcard prefix is too short (it results in HTTP code 400). The value 'expanded query too long' indicates that a wildcard expression resulted in too many terms (which is also code 400). The values here are constants and are subject to api stability requirements. The value 'syntax error' indicates a syntax error in a query. The value 'unsupported expression' indicates that the query as such is valid, but cannot be used in the current context.
// returns a *StatusObject_backendStatus when successful
func (m *StatusObject) GetBackendStatus()(*StatusObject_backendStatus) {
    return m.backendStatus
}
// GetErrorMessage gets the errorMessage property value. A short description of the error. This can be 'null'
// returns a *string when successful
func (m *StatusObject) GetErrorMessage()(*string) {
    return m.errorMessage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StatusObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backendStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatusObject_backendStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackendStatus(val.(*StatusObject_backendStatus))
        }
        return nil
    }
    res["errorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["httpStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpStatus(val)
        }
        return nil
    }
    res["successful"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessful(val)
        }
        return nil
    }
    return res
}
// GetHttpStatus gets the httpStatus property value. The http status of the response. The status depends on 'backendStatus': 'unknown error' is mapped to 'BAD REQUEST' and all others are mapped to 'OK' 
// returns a *int32 when successful
func (m *StatusObject) GetHttpStatus()(*int32) {
    return m.httpStatus
}
// GetSuccessful gets the successful property value. Will be 'true' if the request was fully successful.
// returns a *bool when successful
func (m *StatusObject) GetSuccessful()(*bool) {
    return m.successful
}
// Serialize serializes information the current object
func (m *StatusObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBackendStatus() != nil {
        cast := (*m.GetBackendStatus()).String()
        err := writer.WriteStringValue("backendStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("httpStatus", m.GetHttpStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("successful", m.GetSuccessful())
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
func (m *StatusObject) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBackendStatus sets the backendStatus property value. The status of the backend response if there is one. This can be 'null'. The value 'meaningless search terms' indicates that the query consists only of stop words and the-like and results in HTTP code 200. The value 'all search terms unknown' is similiar, but the server does not know any of the provided terms. The special value 'incomplete search results due to pending synchronization' can only occur in early startup phases of the target server (if synchronizes searchable folder properties asynchronously); it results in HTTP code 200. The value 'child collection invalid' means that a sub-engine is currently unreachable and its results are not part of the result and also results in code 200. The value 'unexpandable wildcard' indicates that a wildcard prefix is too short (it results in HTTP code 400). The value 'expanded query too long' indicates that a wildcard expression resulted in too many terms (which is also code 400). The values here are constants and are subject to api stability requirements. The value 'syntax error' indicates a syntax error in a query. The value 'unsupported expression' indicates that the query as such is valid, but cannot be used in the current context.
func (m *StatusObject) SetBackendStatus(value *StatusObject_backendStatus)() {
    m.backendStatus = value
}
// SetErrorMessage sets the errorMessage property value. A short description of the error. This can be 'null'
func (m *StatusObject) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// SetHttpStatus sets the httpStatus property value. The http status of the response. The status depends on 'backendStatus': 'unknown error' is mapped to 'BAD REQUEST' and all others are mapped to 'OK' 
func (m *StatusObject) SetHttpStatus(value *int32)() {
    m.httpStatus = value
}
// SetSuccessful sets the successful property value. Will be 'true' if the request was fully successful.
func (m *StatusObject) SetSuccessful(value *bool)() {
    m.successful = value
}
type StatusObjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackendStatus()(*StatusObject_backendStatus)
    GetErrorMessage()(*string)
    GetHttpStatus()(*int32)
    GetSuccessful()(*bool)
    SetBackendStatus(value *StatusObject_backendStatus)()
    SetErrorMessage(value *string)()
    SetHttpStatus(value *int32)()
    SetSuccessful(value *bool)()
}
