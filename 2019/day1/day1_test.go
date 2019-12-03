package main

import (
	"testing"
)

func TestCalcFuel(t *testing.T) {
	got := calcfuel(12)
	if got != 2 {
		t.Errorf("calcfuel(12) = %d; want 2", got)
	}

	got = calcfuel(14)
	if got != 2 {
		t.Errorf("calcfuel(14) = %d; want 2", got)
	}

	got = calcfuel(1969)
	if got != 654 {
		t.Errorf("calcfuel(1969) = %d; want 654", got)
	}

	got = calcfuel(100756)
	if got != 33583 {
		t.Errorf("calcfuel(100756) = %d; want 33583", got)
	}
}

func TestCalcTotalFuel(t *testing.T) {
	got := fuel([]int{1969})
	if got != 966 {
		t.Errorf("calctotalfuel(1969) = %d; want 966", got)
	}

	got = fuel([]int{100756})
	if got != 50346 {
		t.Errorf("calctotalfuel(100756) = %d; want 50346", got)
	}
}
