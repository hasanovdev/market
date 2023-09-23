package helper

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// CustomIDGenerator represents an ID generator with a specific format.
type CustomIDGenerator struct {
	prefix string
	mu     sync.Mutex
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, "@"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, "@"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func ReplaceStringForSql(s string) string {
	if strings.Contains(s, "'") {
		s = strings.ReplaceAll(s, "'", "''")
	}

	return s
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewCustomIDGenerator creates a new CustomIDGenerator with default values.
func NewCustomIDGenerator() *CustomIDGenerator {
	return &CustomIDGenerator{
		prefix: "P-",
	}
}

// GenerateID generates a unique custom ID using a timestamp and a counter.
func (gen *CustomIDGenerator) GenerateID() string {
	gen.mu.Lock()
	defer gen.mu.Unlock()

	timestamp := time.Now().UnixNano()
	id := fmt.Sprintf("%s%d", gen.prefix, timestamp)

	id = id[:2] + id[11:]

	return id
}
