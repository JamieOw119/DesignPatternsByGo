package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		if got := NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
		}
	}
}
