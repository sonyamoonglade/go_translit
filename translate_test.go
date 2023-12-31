package go_translit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	values := []string{"privet мир", "bluetooth", "vg123", "абвгдеёжзийклмнопрстуфхцчшщъыьэюя", "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "githjub123@mail.com"}
	encoded := Encode(values)
	require.Equal(t, "=privet mir", encoded[0])
	require.Equal(t, "=bluetooth", encoded[1])
	require.Equal(t, "=vg123", encoded[2])
	require.Equal(t, "abvgde<jzi*klmnoprstufh#(?%@+.,[{", encoded[3])
	// Not all uppercase symbols are encoded!
	require.Equal(t, "ABVGDE>JZI~KLMNOPRSTUFH!)&$ЪЫЬЭ]}", encoded[4])
	require.True(t, len(values[3])/2 == len(encoded[3]))
}

func TestDecode(t *testing.T) {
	values := []string{"privet мир", "bluetooth", "vg123", "абвгдеёжзийклмнопрстуфхцчшщъыьэюя", "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "githjub123@mail.com"}
	encoded := Encode(values)
	decoded := Decode(encoded)
	for i := range values {
		require.Equal(t, values[i], decoded[i])
	}
}

func TestEncodeWordDecodeWord(t *testing.T) {
	v := "privet мир123 hello mir. @ma2"
	encoded := EncodeWord(v)
	require.Equal(t, DecodeWord(encoded), v)
}
