package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SiteSource 
type SiteSource struct {
    DataSource
    // The site property
    site i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Siteable
}
// NewSiteSource instantiates a new SiteSource and sets the default values.
func NewSiteSource()(*SiteSource) {
    m := &SiteSource{
        DataSource: *NewDataSource(),
    }
    odataTypeValue := "#microsoft.graph.security.siteSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSiteSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSiteSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSiteSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SiteSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["site"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateSiteFromDiscriminatorValue , m.SetSite)
    return res
}
// GetSite gets the site property value. The site property
func (m *SiteSource) GetSite()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Siteable) {
    return m.site
}
// Serialize serializes information the current object
func (m *SiteSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSite sets the site property value. The site property
func (m *SiteSource) SetSite(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Siteable)() {
    m.site = value
}
