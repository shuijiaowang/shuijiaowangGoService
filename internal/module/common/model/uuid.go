package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUID 自定义 UUID 类型，嵌入 google/uuid.UUID
type UUID uuid.UUID

// Value 实现 driver.Valuer 接口（Go -> 数据库）
func (u UUID) Value() (driver.Value, error) {
	return uuid.UUID(u).MarshalBinary()
}

// Scan 实现 sql.Scanner 接口（数据库 -> Go）
func (u *UUID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case []byte:
		if len(v) == 0 {
			return nil
		}
		var temp uuid.UUID
		err := temp.UnmarshalBinary(v)
		if err != nil {
			return fmt.Errorf("invalid UUID bytes: %w", err)
		}
		*u = UUID(temp)
		return nil
	case string:
		if v == "" {
			return nil
		}
		parsedUUID, err := uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("invalid UUID string: %w", err)
		}
		*u = UUID(parsedUUID)
		return nil
	default:
		return fmt.Errorf("cannot scan %T into UUID", src)
	}
}

// String 返回 UUID 的字符串表示形式
func (u UUID) String() string {
	return uuid.UUID(u).String()
}

// MarshalJSON 实现 json.Marshaler 接口
func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (u *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	*u = UUID(parsed)
	return nil
}

// NewUUID 创建一个新的 UUID
func NewUUID() UUID {
	return UUID(uuid.New())
}

// ParseUUID 从字符串解析 UUID
func ParseUUID(s string) (UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID(parsed), nil
}

// MustParseUUID 从字符串解析 UUID，解析失败时 panic
func MustParseUUID(s string) UUID {
	return UUID(uuid.MustParse(s))
}

// IsZero 检查 UUID 是否为零值
func (u UUID) IsZero() bool {
	return u == UUID{}
}
