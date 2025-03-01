package otk_test

import (
	"go/build"
	"log"
	"runtime"
	"testing"

	"github.com/alpineiq/otk"
)

func TestCaller(t *testing.T) {
	log.Printf("%#+v", build.Default)
	t.Log(otk.Caller(0, true))
	t.Log(otk.Caller(0, false))
	t.Log(runtime.Caller(0))
}
