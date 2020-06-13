package cmd

import (
	"reflect"
	"testing"
)


func TestParseLine(t *testing.T) {
	type args struct {
		inputStr     string
		previousLine LogLine
	}
	tests := []struct {
		name         string
		args         args
		wantLogLine  LogLine
		wantErrorMsg string
	}{
		{
			"Parse band and mode only", 
			args{ inputStr: "40m cw", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Band: "40m", Mode: "CW", RSTsent: "599", RSTrcvd: "599"}, "",
		},
		{
			"Parse for time", 
			args{ inputStr: "1314 g3noh", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Time: "1314", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse partial time - 1", 
			args{ inputStr: "4 g3noh", previousLine: LogLine{ Time: "", Mode: "SSB", }}, 
			LogLine{ Time: "4", Call: "G3NOH", Mode: "SSB",}, "", //TODO: should fail
		},
		{
			"Parse partial time - 2", 
			args{ inputStr: "15 g3noh", previousLine: LogLine{ Time: "1200", Mode: "SSB", }}, 
			LogLine{ Time: "1215", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse partial time - 3", 
			args{ inputStr: "4 g3noh", previousLine: LogLine{ Time: "1200", Mode: "SSB", }}, 
			LogLine{ Time: "1204", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse for comment", 
			args{ inputStr: "4 g3noh <PSE QSL Direct>", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Time: "4", Comment: "PSE QSL Direct", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse for QSL", 
			args{ inputStr: "g3noh [Custom QSL message]", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ QSLmsg: "Custom QSL message", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Wrong mode", 
			args{ inputStr: "cww", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Mode: "SSB",}, "Unable to parse cww ",
		},
		{
			"Parse OM name", 
			args{ inputStr: "@Jean", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ OMname: "Jean", Mode: "SSB",}, "",
		},
		{
			"Parse Grid locator", 
			args{ inputStr: "#grid", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ GridLoc: "grid", Mode: "SSB",}, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogLine, gotErrorMsg := ParseLine(tt.args.inputStr, tt.args.previousLine)
			if !reflect.DeepEqual(gotLogLine, tt.wantLogLine) {
				t.Errorf("ParseLine() gotLogLine = %v, want %v", gotLogLine, tt.wantLogLine)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ParseLine() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
