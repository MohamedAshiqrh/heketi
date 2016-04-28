package cmds

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/heketi/heketi/client/cli/go/cmds"
)

var (
	diff                    int
	HEKETI_CLI_TEST_VERSION           = "testing"
	sout                    io.Writer = os.Stdout
	serr                    io.Writer = os.Stderr
)
var test_version = []struct {
	input  []string
	output string
}{
	{[]string{"--version", "-v"}, ""},
	{[]string{"--veri"}, "unknown flag: --veri"},
}

var test_cluster = []struct {
	input  []string
	output string
}{
	{[]string{"", "list"}, ""},
}

var test_device = []struct {
	input  []string
	output string
}{
	{[]string{""}, ""},
	{[]string{"--veri"}, "unknown flag: --veri"},
}

var test_load = []struct {
	input  []string
	output string
}{
	{[]string{""}, ""},
	{[]string{"--veri"}, "unknown flag: --veri"},
}

var test_node = []struct {
	input  []string
	output string
}{
	{[]string{""}, ""},
	{[]string{"--veri"}, "unknown flag: --veri"},
}

var test_volume = []struct {
	input  []string
	output string
}{
	{[]string{""}, ""},
	{[]string{"--veri"}, "unknown flag: --veri"},
}

func TestVersion(t *testing.T) {
	c := cmds.NewHeketiCli(HEKETI_CLI_TEST_VERSION, sout, serr)

	buf := new(bytes.Buffer)
	for _, test_arg := range test_version {
		c.SetArgs(test_arg.input)
		c.SetOutput(buf)
		output := c.Execute()
		if output != nil {
			diff = strings.Compare(output.Error(), test_arg.output)
			if diff != 0 {
				t.Error("Expected ", test_arg.output, ",Got ", output.Error())
			}
		}
	}
}

func TestCluster(t *testing.T) {
	c := cmds.NewHeketiCli(HEKETI_CLI_TEST_VERSION, sout, serr)
	c.ResetFlags()
	buf := new(bytes.Buffer)
	for _, test_clu := range test_cluster {
		c.SetArgs(test_clu.input)
		c.SetOutput(buf)
		output := c.Execute()
		if output != nil {
			diff = strings.Compare(output.Error(), test_clu.output)
			if diff != 0 {
				t.Error("Expected ", test_clu.output, ",Got ", output.Error())
			}
		}
	}
}

func TestDevice(t *testing.T) {
	c := cmds.NewHeketiCli(HEKETI_CLI_TEST_VERSION, sout, serr)
	buf := new(bytes.Buffer)
	for _, test_arg := range test_device {
		c.SetArgs(test_arg.input)
		c.SetOutput(buf)
		output := c.Execute()
		if output != nil {
			diff = strings.Compare(output.Error(), test_arg.output)
			if diff != 0 {
				t.Error("Expected ", test_arg.output, ",Got ", output.Error())
			}
		}
	}
}

func TestLoad(t *testing.T) {
	c := cmds.NewHeketiCli(HEKETI_CLI_TEST_VERSION, sout, serr)
	buf := new(bytes.Buffer)
	for _, test_arg := range test_load {
		c.SetArgs(test_arg.input)
		c.SetOutput(buf)
		output := c.Execute()
		if output != nil {
			diff = strings.Compare(output.Error(), test_arg.output)
			if diff != 0 {
				t.Error("Expected ", test_arg.output, ",Got ", output.Error())
			}
		}
	}
}

func TestVolume(t *testing.T) {
	c := cmds.NewHeketiCli(HEKETI_CLI_TEST_VERSION, sout, serr)
	buf := new(bytes.Buffer)
	for _, test_arg := range test_volume {
		c.SetArgs(test_arg.input)
		c.SetOutput(buf)
		output := c.Execute()
		if output != nil {
			diff = strings.Compare(output.Error(), test_arg.output)
			if diff != 0 {
				t.Error("Expected ", test_arg.output, ",Got ", output.Error())
			}
		}
	}
}
