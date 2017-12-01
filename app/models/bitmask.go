package models

// Bitmask is binary representation
type Bitmask uint8

// HasFlag return true if mask has specific flag set
func (b Bitmask) HasFlag(flag Bitmask) bool { return b&flag != 0 }

// AddFlag to the binary representation
func (b *Bitmask) AddFlag(flag Bitmask)     { *b |= flag }

// ClearFlag remove flag
func (b *Bitmask) ClearFlag(flag Bitmask)   { *b &= ^flag }

// ToggleFlag inverse state of the flag
func (b *Bitmask) ToggleFlag(flag Bitmask)  { *b ^= flag }

/*
// Value transform Json to friendly value for database
func (b Bitmask) Value() (driver.Value, error) {
	return b, nil
}

// Scan retrieve value from database
func (b *Bitmask) Scan(value interface{}) error {
	if value == nil {
		*b = 0
		return nil
	}

	if v, ok := value.(int64); ok {
		*b = Bitmask(v)
		return nil
	}

	return errors.New("can't cast Bitmask to uint8")
}
*/
