package adt

type DefaultMap struct {
	Data         map[interface{}]interface{}
	defaultValue interface{}
}

func NewDefaultMap(defaultValue interface{}) *DefaultMap {
	var newMap DefaultMap

	newMap.Data = map[interface{}]interface{}{}
	newMap.defaultValue = defaultValue

	return &newMap
}

func (m *DefaultMap) Get(key interface{}) interface{} {
	if value, exists := m.Data[key]; exists {
		return value
	}

	return m.defaultValue
}

func (m *DefaultMap) Set(key interface{}, value interface{}) {
	m.Data[key] = value
}

func (m *DefaultMap) SetDefaultValue(defaultValue interface{}) {
	m.defaultValue = defaultValue
}
