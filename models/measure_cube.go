package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MeasureCube struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dimensions property
    dimensions []MeasureDimensionResultable
    // The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a long integer. It will be 'null' if the value is a double. A value is long if the underlying field type is long and the measureType results in long values.
    values []int64
    // The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a floating point number. It will be 'null' if the value is a long. A value is double if the underlying field type is double or the measureType results in double values.
    valuesDouble []float64
}
// NewMeasureCube instantiates a new MeasureCube and sets the default values.
func NewMeasureCube()(*MeasureCube) {
    m := &MeasureCube{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeasureCubeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMeasureCubeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeasureCube(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MeasureCube) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDimensions gets the dimensions property value. The dimensions property
// returns a []MeasureDimensionResultable when successful
func (m *MeasureCube) GetDimensions()([]MeasureDimensionResultable) {
    return m.dimensions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MeasureCube) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dimensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeasureDimensionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeasureDimensionResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeasureDimensionResultable)
                }
            }
            m.SetDimensions(res)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int64))
                }
            }
            m.SetValues(res)
        }
        return nil
    }
    res["valuesDouble"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("float64")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]float64, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*float64))
                }
            }
            m.SetValuesDouble(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a long integer. It will be 'null' if the value is a double. A value is long if the underlying field type is long and the measureType results in long values.
// returns a []int64 when successful
func (m *MeasureCube) GetValues()([]int64) {
    return m.values
}
// GetValuesDouble gets the valuesDouble property value. The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a floating point number. It will be 'null' if the value is a long. A value is double if the underlying field type is double or the measureType results in double values.
// returns a []float64 when successful
func (m *MeasureCube) GetValuesDouble()([]float64) {
    return m.valuesDouble
}
// Serialize serializes information the current object
func (m *MeasureCube) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDimensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDimensions()))
        for i, v := range m.GetDimensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dimensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        err := writer.WriteCollectionOfInt64Values("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    if m.GetValuesDouble() != nil {
        err := writer.WriteCollectionOfFloat64Values("valuesDouble", m.GetValuesDouble())
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
func (m *MeasureCube) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDimensions sets the dimensions property value. The dimensions property
func (m *MeasureCube) SetDimensions(value []MeasureDimensionResultable)() {
    m.dimensions = value
}
// SetValues sets the values property value. The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a long integer. It will be 'null' if the value is a double. A value is long if the underlying field type is long and the measureType results in long values.
func (m *MeasureCube) SetValues(value []int64)() {
    m.values = value
}
// SetValuesDouble sets the valuesDouble property value. The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a floating point number. It will be 'null' if the value is a long. A value is double if the underlying field type is double or the measureType results in double values.
func (m *MeasureCube) SetValuesDouble(value []float64)() {
    m.valuesDouble = value
}
type MeasureCubeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDimensions()([]MeasureDimensionResultable)
    GetValues()([]int64)
    GetValuesDouble()([]float64)
    SetDimensions(value []MeasureDimensionResultable)()
    SetValues(value []int64)()
    SetValuesDouble(value []float64)()
}
