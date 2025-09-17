package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// JSONDate 自定义日期类型（只包含年月日）
type JSONDate time.Time

// 实现JSON反序列化接口（处理前端输入）
func (j *JSONDate) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}

	// 去除JSON字符串两端的引号
	if len(s) > 1 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	// 支持多种日期格式解析
	layouts := []string{
		"2006-1-2",    // 单数字月/日
		"2006-01-02",  // 双数字月/日
		time.DateOnly, // Go 1.20+ 标准格式
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, s)
		if err == nil {
			*j = JSONDate(parsedTime)
			return nil
		}
	}

	return fmt.Errorf("invalid date format: %s, supported formats: YYYY-M-D, YYYY-MM-DD", s)
}

// 实现JSON序列化接口（返回前端）
func (j JSONDate) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

// 实现数据库Scan接口（从数据库读取）
func (j *JSONDate) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid type for JSONDate")
	}
	*j = JSONDate(t)
	return nil
}

// 实现数据库Value接口（写入数据库）
func (j JSONDate) Value() (driver.Value, error) {
	t := time.Time(j)
	// 处理零值情况
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

// 转换为标准time.Time
func (j JSONDate) ToTime() time.Time {
	return time.Time(j)
}

// 字符串表示（用于日志/调试）
func (j JSONDate) String() string {
	return time.Time(j).Format("2006-01-02")
}

//需要添加一个函数把time。time类型的转为格式
