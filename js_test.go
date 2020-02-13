package js_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tetratom/js"
)

func TestValue(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		raw := json.RawMessage(`
	{
		"str": "test",
		"num": 3.5,
		"bool": true,
		"arr": ["tset", 3.5, true],
		"obj": {
			"test": 1
		}
	}
	`)

		obj, err := js.ParseObject(raw)
		require.NoError(t, err)

		require.Equal(t, js.Object{
			"str":  js.String("test"),
			"num":  js.Number(3.5),
			"bool": js.Bool(true),
			"arr": js.Array{
				js.String("tset"),
				js.Number(3.5),
				js.Bool(true),
			},
			"obj": js.Object{
				"test": js.Number(1),
			},
		}, obj)

		require.Equal(t, "test", obj.GetString("str"))
		require.Equal(t, 3.5, obj.GetFloat64("num"))
		require.Equal(t, int64(3), obj.GetInt64("num"))
		require.Equal(t, true, obj.GetBool("bool"))
		require.Equal(t, "tset", obj.GetArray("arr").GetString(0))
		require.Equal(t, 3.5, obj.GetArray("arr").GetFloat64(1))
		require.Equal(t, true, obj.GetArray("arr").GetBool(2))
		require.Equal(t, int64(1), obj.GetObject("obj").GetInt64("test"))
	})

	t.Run("marshal", func(t *testing.T) {
		data, err := json.Marshal(js.Object{"str": js.String("foo"), "null": js.Null{}})
		require.NoError(t, err)
		require.Equal(t, `{"null":null,"str":"foo"}`, string(data))
	})
}
