package ai

import (
	"reflect"
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
)

func TestNewAIService(t *testing.T) {
	type args struct {
		provider config.Provider
	}
	tests := []struct {
		name string
		args args
		want AIservice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAIService(tt.args.provider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAIService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aiServiceImpl_Summarize(t *testing.T) {
	type fields struct {
		provider config.Provider
	}
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &aiServiceImpl{
				provider: tt.fields.provider,
			}
			got, err := g.Summarize(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Summarize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Summarize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
