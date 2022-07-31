package file

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestWriteString(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/" + strconv.Itoa(rand.Intn(999-100)) + ".txt"
	defer os.Remove(file)

	content := "aaa"
	WriteString(file, content)
	if ReadString(file) != content {
		t.Error()
	}
}

func TestAppendString(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/" + strconv.Itoa(rand.Intn(999-100)) + ".txt"
	defer os.Remove(file)

	WriteString(file, "aaa")
	AppendString(file, "bbb")
	readString := ReadString(file)
	if readString != "aaabbb" {
		t.Error(readString)
	}
}

func TestAppendLine(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/" + strconv.Itoa(rand.Intn(999-100)) + ".txt"
	defer os.Remove(file)

	WriteString(file, "aaa")
	AppendLine(file, "bbb")
	readString := ReadString(file)
	if readString != "aaa\nbbb" {
		t.Error(readString)
	}
}

func TestAppendAllLine(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/" + strconv.Itoa(rand.Intn(999-100)) + ".txt"
	defer os.Remove(file)

	WriteString(file, "aaa")
	str := []string{"bbb", "ccc"}
	AppendAllLine(file, str)
	readString := ReadString(file)
	if readString != "aaa\nbbb\nccc" {
		t.Error(readString)
	}
}
