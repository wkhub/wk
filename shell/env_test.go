package shell

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvHas(t *testing.T) {
	os.Setenv("WK_TEST", "value")
	assert.True(t, GetEnv().Has("WK_TEST"))
	assert.False(t, GetEnv().Has("WK_TEST_NOT_FOUND"))
}

func TestEnvGet(t *testing.T) {
	os.Setenv("WK_TEST", "value")
	env := GetEnv()
	actual := env.Get("WK_TEST", "not found")
	assert.Equal(t, "value", actual)
}

func TestEnvGetWithDefault(t *testing.T) {
	os.Unsetenv("WK_TEST")
	env := GetEnv()
	actual := env.Get("WK_TEST", "not found")
	assert.Equal(t, "not found", actual)
}

func TestEnvGetInt(t *testing.T) {
	os.Setenv("WK_TEST", "15")
	env := GetEnv()
	actual := env.GetInt("WK_TEST", 42)
	assert.Equal(t, 15, actual)
}

func TestEnvGetIntDefault(t *testing.T) {
	os.Unsetenv("WK_TEST")
	env := GetEnv()
	actual := env.GetInt("WK_TEST", 42)
	assert.Equal(t, 42, actual)
}

func TestEnvGetIntInvalid(t *testing.T) {
	os.Setenv("WK_TEST", "unparseable")
	env := GetEnv()
	actual := env.GetInt("WK_TEST", 42)
	assert.Equal(t, 42, actual)
}

func TestEnvClone(t *testing.T) {
	env := GetEnv()
	clone := env.Clone()
	assert.Equal(t, env, clone)
}

func TestEnvEnviron(t *testing.T) {
	env := GetEnv()
	environ := env.Environ()
	assert.Equal(t, len(env), len(environ))
	for _, value := range environ {
		assert.Contains(t, value, "=")
	}
}
