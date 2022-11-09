package core

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatalf("测试失败, 1 + 2 预期为 3")
	}
}

func TestGreeting(t *testing.T) {
	if Greeting() != "hello word" {
		t.Fatalf("测试失败, Greeting 预期返回 'hello word'")
	}
}
