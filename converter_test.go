package converter

import "testing"

/*
type converterTest struct {
	i        float64
	expected Unit
}
*/

func TestMtoF(t *testing.T) {
	MtoFTests := []struct {
		a float64
		b Unit
	}{
		{12, 39.37007874015748031496062992126},
		{50, 164.04199475065616797900262467192},
		{7, 22.965879265091863517060367454068},
	}

	for _, ct := range MtoFTests {
		now := MtoF(ct.a)
		if now != ct.b {
			t.Errorf("expected %v, now %v", ct.b, now)
		}
	}
}

func TestFtoM(t *testing.T) {
	FtoMTests := []struct {
		a float64
		b Unit
	}{
		{45, 13.716},
		{16, 4.8768},
		{23, 7.0104},
	}
	for _, ct := range FtoMTests {
		now := FtoM(ct.a)
		if now != ct.b {
			t.Errorf("expected %v, now %v", ct.b, now)
		}
	}
}
