package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SponsoredLink a list of matching sponsored links (if configured for the target project).
type SponsoredLink struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A detailed description of the result.
    description *string
    // A link into some external system (typically HTTP). Can be null.
    externalLink *string
    // If present, this is the record id associated with the link. It is present if and only if the link resembles a record id in the collection.
    recordId *string
    // The confidence for this link (between 0 and 1).
    relevance *float32
    // Link text describing the result.
    title *string
}
// NewSponsoredLink instantiates a new SponsoredLink and sets the default values.
func NewSponsoredLink()(*SponsoredLink) {
    m := &SponsoredLink{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSponsoredLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSponsoredLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSponsoredLink(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SponsoredLink) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A detailed description of the result.
// returns a *string when successful
func (m *SponsoredLink) GetDescription()(*string) {
    return m.description
}
// GetExternalLink gets the externalLink property value. A link into some external system (typically HTTP). Can be null.
// returns a *string when successful
func (m *SponsoredLink) GetExternalLink()(*string) {
    return m.externalLink
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SponsoredLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["externalLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalLink(val)
        }
        return nil
    }
    res["recordId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordId(val)
        }
        return nil
    }
    res["relevance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelevance(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetRecordId gets the recordId property value. If present, this is the record id associated with the link. It is present if and only if the link resembles a record id in the collection.
// returns a *string when successful
func (m *SponsoredLink) GetRecordId()(*string) {
    return m.recordId
}
// GetRelevance gets the relevance property value. The confidence for this link (between 0 and 1).
// returns a *float32 when successful
func (m *SponsoredLink) GetRelevance()(*float32) {
    return m.relevance
}
// GetTitle gets the title property value. Link text describing the result.
// returns a *string when successful
func (m *SponsoredLink) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *SponsoredLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalLink", m.GetExternalLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recordId", m.GetRecordId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("relevance", m.GetRelevance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *SponsoredLink) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A detailed description of the result.
func (m *SponsoredLink) SetDescription(value *string)() {
    m.description = value
}
// SetExternalLink sets the externalLink property value. A link into some external system (typically HTTP). Can be null.
func (m *SponsoredLink) SetExternalLink(value *string)() {
    m.externalLink = value
}
// SetRecordId sets the recordId property value. If present, this is the record id associated with the link. It is present if and only if the link resembles a record id in the collection.
func (m *SponsoredLink) SetRecordId(value *string)() {
    m.recordId = value
}
// SetRelevance sets the relevance property value. The confidence for this link (between 0 and 1).
func (m *SponsoredLink) SetRelevance(value *float32)() {
    m.relevance = value
}
// SetTitle sets the title property value. Link text describing the result.
func (m *SponsoredLink) SetTitle(value *string)() {
    m.title = value
}
type SponsoredLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetExternalLink()(*string)
    GetRecordId()(*string)
    GetRelevance()(*float32)
    GetTitle()(*string)
    SetDescription(value *string)()
    SetExternalLink(value *string)()
    SetRecordId(value *string)()
    SetRelevance(value *float32)()
    SetTitle(value *string)()
}
