package main

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := NewConfig()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}

	wantEnv := "dev"

	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}

}
