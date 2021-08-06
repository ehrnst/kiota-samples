package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Entity struct {
    additionalData map[string]interface{};
    id *string;
}
func NewEntity()(*Entity) {
    m := &Entity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Entity) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *Entity) GetId()(*string) {
    return m.id
}
func (m *Entity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *Entity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := writer.WritePrimitiveValue("id", m.GetId())
    if err != nil {
        return err
    }
    err = writer.WriteAdditionalData(m.GetAdditionalData())
    if err != nil {
        return err
    }
    return nil
}
func (m *Entity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Entity) SetId(value *string)() {
    m.id = value
}
