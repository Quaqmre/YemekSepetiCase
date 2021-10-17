package store

import (
	"io/ioutil"
	"os"
	"testing"
)

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "test.temp")
	if err != nil {
		t.Fatal(t)
	}
	t.Log("Using temp-dir:", dir)
	return dir
}

func Test_Store_File_Is_Exist(t *testing.T) {
	_, _ = NewStore()
	defer os.Remove("./test.tmp")

	if _, err := os.Stat("./test.tmp"); os.IsNotExist(err) {
		t.Fatal("test.temp file is not exist")
	}

}

var cases = []struct {
	key      int32
	value    string
	expected string
	err      string
}{
	{0, "val1", "val1", "Putting failed for value1 key 1"},
	{1, "val2", "val2", "Putting failed for value2 key 2"},
	{2, "val3", "val3", "Putting failed for value3 key 3"},
	{3, "val4", "val4", "Putting failed for value4 key 4"},
}

func Test_Store_Key_Is_Exist(t *testing.T) {
	fs, _ := NewStore()

	for _, v := range cases {
		fs.Put(v.value)
	}

	for _, v := range cases {

		if val := fs.Get(string(v.key)); val != v.expected {
			t.Fatal(v.err)
		}
	}

}

func Test_Store_Can_Flush(t *testing.T) {
	fs, _ := NewStore()

	for _, v := range cases {
		fs.Put(v.value)
	}
	fs.Flush()
	for _, _ = range fs.db {
		t.Fatal("flush cant finish as expected")
	}
}
func Test_Store_Pointer_Is_Zero(t *testing.T) {
	fs, _ := NewStore()

	for _, v := range cases {
		fs.Put(v.value)
	}
	fs.Flush()
	var test int32 = 0
	if *fs.Ops != test {
		t.Fatal("after flush Ops isnt zero Ops:", *fs.Ops)
	}
}
