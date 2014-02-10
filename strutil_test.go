package gosk

import (
    "testing"
)

func TestTrimHTML(t *testing.T) {
    str := "<html> test </htm"
    if trimHTML(str) == " test " {
        t.Log("trimHTML test passing")
    } else {
        t.Error("trimHTML test error")
    }
}
