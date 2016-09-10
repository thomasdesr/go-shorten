package storage

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var replacerTestTable = []struct {
	chars    string
	input    string
	expected string
}{
	{
		chars:    "ad",
		input:    "asdf",
		expected: "sf",
	},
	{
		chars:    "",
		input:    "asdf",
		expected: "asdf",
	},
	{
		chars:    "asdf",
		input:    "asdf",
		expected: "",
	},
	{
		chars:    "",
		input:    "",
		expected: "",
	},
}

func mapFilterChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
func TestMapFilterChars(t *testing.T) {
	for _, tbl := range replacerTestTable {
		t.Logf("Table: %#v", tbl)

		actual := mapFilterChars(tbl.input, tbl.chars)
		t.Logf("Actual: %q", actual)

		assert.Equal(t, actual, tbl.expected)
	}
}

func dropCharsReplacer(chars string) *strings.Replacer {
	dropChars := []rune(chars)
	args := make([]string, len(dropChars)*2)
	for i := 0; i < len(dropChars); i++ {
		args[i*2] = string(dropChars[i])
	}
	return strings.NewReplacer(args...)
}

func TestDropCharsReplacer(t *testing.T) {
	for _, tbl := range replacerTestTable {
		t.Logf("Table: %#v", tbl)

		actual := dropCharsReplacer(tbl.chars).Replace(tbl.input)
		t.Logf("Actual: %q", actual)

		assert.Equal(t, actual, tbl.expected)
	}
}

func BenchmarkMapFilter_ShortString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapFilterChars("asdfasdfasdf", "ad")
	}
}

func BenchmarkReplacer_ShortString(b *testing.B) {
	r := dropCharsReplacer("ad")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Replace("asdfasdfasdf")
	}
}

func BenchmarkMapFilter_LongString(b *testing.B) {
	in := strings.Repeat("asdf", 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mapFilterChars(in, "ad")
	}
}

func BenchmarkReplacer_LongString(b *testing.B) {
	in := strings.Repeat("asdf", 1000)
	r := dropCharsReplacer("ad")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Replace(in)
	}
}

func BenchmarkMapFilter_LongReplace(b *testing.B) {
	in := strings.Repeat("asdf", 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mapFilterChars(in, "qwertyuiopasdfghjklzxcvbnm,./;'[]>?<,:\"{}[]|\\+_=-!@#$%^&*()1234567890")
	}
}

func BenchmarkReplacer_LongReplace(b *testing.B) {
	in := strings.Repeat("asdf", 1000)
	r := dropCharsReplacer("qwertyuiopasdfghjklzxcvbnm,./;'[]>?<,:\"{}[]|\\+_=-!@#$%^&*()1234567890")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Replace(in)
	}
}
