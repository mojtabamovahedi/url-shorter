package cache

import (
	"encoding/json"
)

func (c *ObjectCacher[T]) unmarshal(data []byte, out any) error {
	if c.serializationType == SerializationTypeJson {
		return json.Unmarshal(data, out)
	}
	// don't need to implement anything for now

	return nil
}

func (c *ObjectCacher[T]) marshal(in any) ([]byte, error) {
	if c.serializationType == SerializationTypeJson {
		return json.Marshal(in)
	}
	// don't need to implement anything for now

	return nil, nil
}
