package domain

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {

	p := NewPackage()
	err := p.Submit(SubmitCommand{})
	if err != nil {
		fmt.Println("submit error", err)
		t.Fail()
	}
	if p.Status() != StatusSubmitted {
		fmt.Println("submit change unexpected")
		t.Fail()
	}
}
