package main

import (
	"reflect"
	"testing"
)

func TestNewCarResearch(t *testing.T) {
	tests := []struct {
		name string
		want *CarResearch
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCarResearch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCarResearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessPrice(t *testing.T) {
	type args struct {
		rawPrice string
		percent  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessPrice(tt.args.rawPrice, tt.args.percent); got != tt.want {
				t.Errorf("ProcessPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessTax(t *testing.T) {
	type args struct {
		rawPrice string
		percent  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessTax(tt.args.rawPrice, tt.args.percent); got != tt.want {
				t.Errorf("ProcessTax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessVolume(t *testing.T) {
	type args struct {
		hp  string
		evt float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessVolume(tt.args.hp, tt.args.evt); got != tt.want {
				t.Errorf("ProcessVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}
