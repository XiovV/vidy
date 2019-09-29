package main

import (
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

	for _, test := range tests {
		if output := convertURLToReadablePath(test.url); output != test.expected {
			t.Errorf("input: %v | expected %v | recieved: %v", test.url, test.expected, output)
		}
	}
}
