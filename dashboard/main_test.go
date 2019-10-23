package main

import (
	"reflect"
	"testing"
)

func TestConvertURLToReadablePath(t *testing.T) {
	var tests = []struct {
		url      string
		expected string
	}{
		{"/video/One%20Two/video.mp4", "/library/One Two"},
		{"/video/video.mp4", "/library"},
		{"/video/subdir1/subdir2/video.mp4", "/library/subdir1/subdir2"},
		{"/video/sub%20dir1/video%201.mp4", "/library/sub dir1"},
	}

	t.Run("Return a readable path without .mp4 at the end of the path", func(t *testing.T) {
		for _, test := range tests {
			if output := ConvertURLToReadablePath(test.url); output != test.expected {
				t.Errorf("input: %v | expected %v | recieved: %v", test.url, test.expected, output)
			}
		}
	})
}

func TestReadDir(t *testing.T) {
	var tests = []struct {
		path     string
		expected []string
	}{
		{"library_test", []string{"sub dir1", "subdir1", "test.mp4"}},
		{"library_test/subdir1", []string{"one two three.mp4", "test1.mp4", "test2.mp4"}},
		{"library_test/sub dir1", []string{"test 1.mp4"}},
		{"library_test/sub%20dir1", []string{"test 1.mp4"}},
	}

	t.Run("Return correct files/directories from a directory", func(t *testing.T) {
		for _, test := range tests {
			if output := ReadDir(test.path); !reflect.DeepEqual(output, test.expected) {
				t.Errorf("input: %q | expected %q | recieved: %q", test.path, test.expected, output)
			}
		}
	})
}
