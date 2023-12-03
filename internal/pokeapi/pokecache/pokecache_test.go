package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}

	for _, cse := range cases {
		cache.Add(cse.inputKey, cse.inputVal)
		actual, ok := cache.Get(cse.inputKey)
		if !ok {
			t.Errorf("%s not found", cse.inputKey)
		}
		if string(actual) != string(cse.inputVal) {
			t.Errorf("%s doesnt match %s", cse.inputVal, cse.inputKey)
		}
	}
}
