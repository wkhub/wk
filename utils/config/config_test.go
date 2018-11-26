package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wkhub/wk/utils/config"
)

func TestCascadeGet(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", "value1")
	cfg2.Set("key", "value2")

	cfg1.Set("key1", "value1")

	cfg2.Set("key2", "value2")

	// Value taken from cfg1 over cfg2
	assert.True(t, cfg.IsSet("key"))
	assert.Equal(t, cfg.Get("key"), "value1")

	// Only present in cfg1
	assert.True(t, cfg.IsSet("key1"))
	assert.Equal(t, cfg.Get("key1"), "value1")

	// Only present in cfg2
	assert.True(t, cfg.IsSet("key2"))
	assert.Equal(t, cfg.Get("key2"), "value2")

	// Missing from every config
	assert.False(t, cfg.IsSet("missing"))
	assert.Equal(t, cfg.Get("missing"), nil)
}

func TestCascadeSet(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg.Set("key", "value")

	assert.True(t, cfg.IsSet("key"))
	assert.Equal(t, cfg.Get("key"), "value")

	assert.True(t, cfg1.IsSet("key"))
	assert.Equal(t, cfg1.Get("key"), "value")

	assert.False(t, cfg2.IsSet("key"))
	assert.Equal(t, cfg2.Get("key"), nil)
}

func TestCascadeSub(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()
	cfg3 := config.New()

	cfg := config.Cascade(cfg1, cfg2, cfg3)

	cfg1.Set("root.key", "value1")
	cfg2.Set("root.key", "value2")

	cfg1.Set("root.key1", "value1")

	cfg2.Set("root.key2", "value2")

	cfg3.Set("outside.missing", "value")

	sub := cfg.Sub("root")

	// Value taken from cfg1 over cfg2
	assert.True(t, sub.IsSet("key"))
	assert.Equal(t, sub.Get("key"), "value1")

	// Only present in sub1
	assert.True(t, sub.IsSet("key1"))
	assert.Equal(t, sub.Get("key1"), "value1")

	// Only present in sub2
	assert.True(t, sub.IsSet("key2"))
	assert.Equal(t, sub.Get("key2"), "value2")

	// Missing from every config
	for _, key := range []string{"outside.missing", "outside", "missing"} {
		assert.False(t, sub.IsSet(key))
		assert.Equal(t, sub.Get(key), nil)
	}
}

func TestCascadeGetString(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", "value1")
	cfg2.Set("key", "value2")

	cfg1.Set("key1", "value1")

	cfg2.Set("key2", "value2")

	// Value taken from cfg1 over cfg2
	assert.Equal(t, cfg.GetString("key"), "value1")

	// Only present in cfg1
	assert.Equal(t, cfg.GetString("key1"), "value1")

	// Only present in cfg2
	assert.Equal(t, cfg.GetString("key2"), "value2")

	// Missing from every config
	assert.Equal(t, cfg.GetString("missing"), "")
}

func TestCascadeGetBool(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", true)
	cfg2.Set("key", false)

	cfg1.Set("key1", true)

	cfg2.Set("key2", false)

	// Value taken from cfg1 over cfg2
	assert.True(t, cfg.GetBool("key"))

	// Only present in cfg1
	assert.True(t, cfg.GetBool("key1"))

	// Only present in cfg2
	assert.False(t, cfg.GetBool("key2"))

	// Missing from every config
	assert.False(t, cfg.GetBool("missing"))
}

func TestCascadeGetInt(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", 1)
	cfg2.Set("key", 2)

	cfg1.Set("key1", 1)

	cfg2.Set("key2", 2)

	// Value taken from cfg1 over cfg2
	assert.Equal(t, cfg.GetInt("key"), 1)

	// Only present in cfg1
	assert.Equal(t, cfg.GetInt("key1"), 1)

	// Only present in cfg2
	assert.Equal(t, cfg.GetInt("key2"), 2)

	// Missing from every config
	assert.Equal(t, cfg.GetInt("missing"), 0)
}

func TestCascadeGetInt32(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", 1)
	cfg2.Set("key", 2)

	cfg1.Set("key1", 1)

	cfg2.Set("key2", 2)

	// Value taken from cfg1 over cfg2
	assert.Equal(t, cfg.GetInt32("key"), int32(1))

	// Only present in cfg1
	assert.Equal(t, cfg.GetInt32("key1"), int32(1))

	// Only present in cfg2
	assert.Equal(t, cfg.GetInt32("key2"), int32(2))

	// Missing from every config
	assert.Equal(t, cfg.GetInt32("missing"), int32(0))
}

func TestCascadeGetInt64(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", 1)
	cfg2.Set("key", 2)

	cfg1.Set("key1", 1)

	cfg2.Set("key2", 2)

	// Value taken from cfg1 over cfg2
	assert.Equal(t, cfg.GetInt64("key"), int64(1))

	// Only present in cfg1
	assert.Equal(t, cfg.GetInt64("key1"), int64(1))

	// Only present in cfg2
	assert.Equal(t, cfg.GetInt64("key2"), int64(2))

	// Missing from every config
	assert.Equal(t, cfg.GetInt64("missing"), int64(0))
}

func TestCascadeUnmarshal(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", "value1")
	cfg2.Set("key", "value2")

	cfg1.Set("key1", "value1")

	cfg2.Set("key2", "value2")

	type Out struct {
		Key  string
		Key1 string
		Key2 string
	}

	out := Out{}

	err := cfg.Unmarshal(&out)

	assert.Nil(t, err)
	assert.Equal(t, "value1", out.Key)
	assert.Equal(t, "value1", out.Key1)
	assert.Equal(t, "value2", out.Key2)
}

func TestCascadeAllSettings(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", "value1")
	cfg2.Set("key", "value2")

	cfg1.Set("key1", "value1")

	cfg2.Set("key2", "value2")

	all := cfg.AllSettings()

	assert.Equal(t, 3, len(all))
	assert.Equal(t, "value1", all["key"])
	assert.Equal(t, "value1", all["key1"])
	assert.Equal(t, "value2", all["key2"])
}

func TestCascadeAllKeys(t *testing.T) {
	cfg1 := config.New()
	cfg2 := config.New()

	cfg := config.Cascade(cfg1, cfg2)

	cfg1.Set("key", "value1")
	cfg2.Set("key", "value2")

	cfg1.Set("key1", "value1")

	cfg2.Set("key2", "value2")

	keys := cfg.AllKeys()

	expected := []string{"key", "key1", "key2"}

	assert.ElementsMatch(t, expected, keys)
}
