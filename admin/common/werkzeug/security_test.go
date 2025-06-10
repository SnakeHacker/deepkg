package werkzeug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordHash(t *testing.T) {
	actual := CheckPasswordHash("password", "pbkdf2:sha256:150000$myj7pNIO$895b262253b5b2944ba342340d0e9c08100b40a5218cd9e944ed72b43cabf397")
	assert.True(t, actual)
}
