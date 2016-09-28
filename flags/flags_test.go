package flags

import (
	"flag"
	"fmt"
	"gopkg.in/stretchr/testify.v1/assert"
	"os"
	"testing"
)

var defaultArgs []string

func init() {
	copy(defaultArgs, os.Args)
}

func TestFlags_All(t *testing.T) {
	args(func() {
		assert.True(t, flag.Parsed())
		assert.Empty(t, HttpPort)
		assert.Empty(t, HttpsPort)
		assert.Empty(t, CustomPath)
		assert.Empty(t, Log)
		assert.False(t, IsInDevMode)
	})
}

func TestFlags_All_WithValue(t *testing.T) {
	args(func() {
		assert.True(t, flag.Parsed())
		assert.Equal(t, "8080", HttpPort)
		assert.Equal(t, "8081", HttpsPort)
		assert.Equal(t, "/absolute/path/to/custom/folder", CustomPath)
		assert.Equal(t, "path/to/log.txt", Log)
		assert.True(t, IsInDevMode)

	}, "-log=path/to/log.txt", "-dev", "-custom-path=/absolute/path/to/custom/folder", "-http-port=8080", "-https-port=8081")
}

var run = 0

func args(test func(), args ...string) {
	os.Args = []string{fmt.Sprintf("journey-test-%d", run)}
	run++
	os.Args = append(os.Args, args...)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = nil
	parseFlags()
	test()
	fmt.Printf("flags test run %d times\n", run)
}
