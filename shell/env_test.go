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

func TestUpdate(t *testing.T) {
	env1 := make(Env)
	env1["key"] = "value"
	env1["updated"] = "value"

	env2 := make(Env)
	env2["other"] = "other"
	env2["updated"] = "updated"

	env3 := env1.Update(env2)

	assert.Contains(t, env1, "other")
	assert.Equal(t, "value", env1["key"])
	assert.Equal(t, "other", env1["other"])
	assert.Equal(t, "updated", env1["updated"])

	assert.Equal(t, env1, env3)
}

func TestUpdateFromAnyMap(t *testing.T) {
	env := make(Env)
	env["key"] = "value"
	env["updated"] = "value"

	m := make(map[string]string)
	m["other"] = "other"
	m["updated"] = "updated"

	env.Update(m)

	assert.Contains(t, env, "other")
	assert.Equal(t, "value", env["key"])
	assert.Equal(t, "other", env["other"])
	assert.Equal(t, "updated", env["updated"])
}

func TestEnvEnviron(t *testing.T) {
	env := GetEnv()
	environ := env.Environ()
	assert.Equal(t, len(env), len(environ))
	for _, value := range environ {
		assert.Contains(t, value, "=")
	}
}
