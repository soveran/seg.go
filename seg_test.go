package seg_test

import (
	"github.com/soveran/seg.go"
	"testing"
)

func TestConsume(t *testing.T) {
	var result bool

	s := seg.New("/posts/42")

	if s.Prev() != "" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "/posts/42" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Consume("foo"); result != false {
		t.Errorf("Consumed the wrong segment")
	}

	if result = s.Consume("posts"); result != true {
		t.Errorf("Didn't consume the right segment")
	}

	if s.Prev() != "/posts" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "/42" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Consume("42"); result != true {
		t.Errorf("Didn't consume the right segment")
	}

	if s.Prev() != "/posts/42" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Consume("foo"); result != false {
		t.Errorf("Consumed the wrong segment")
	}

	if s.Prev() != "/posts/42" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}
}

func TestCapture(t *testing.T) {
	var result bool

	inbox := make(map[string]string)

	s := seg.New("/posts/42")

	if s.Prev() != "" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "/posts/42" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Capture("p1", inbox); result != true {
		t.Errorf("Didn't capture the segment")
	}

	if inbox["p1"] != "posts" {
		t.Errorf("Didn't store the segment: %s", inbox["p1"])
	}

	if s.Prev() != "/posts" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "/42" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Capture("p2", inbox); result != true {
		t.Errorf("Didn't capture the segment")
	}

	if inbox["p2"] != "42" {
		t.Errorf("Didn't store the segment: %s", inbox["p2"])
	}

	if s.Prev() != "/posts/42" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}

	if result = s.Capture("p3", inbox); result != false {
		t.Errorf("Captured the wrong segment")
	}

	if s.Prev() != "/posts/42" {
		t.Errorf("Wrong previous path: %s", s.Prev())
	}

	if s.Curr() != "" {
		t.Errorf("Wrong current path: %s", s.Curr())
	}
}
