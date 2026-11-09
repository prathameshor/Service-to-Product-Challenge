package day_three

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		Name     string
		Nums     []int
		Expected int
	}{
		{
			Name:     "Test One",
			Nums:     []int{3, 2, 3},
			Expected: 3,
		},
		{
			Name:     "Test Two",
			Nums:     []int{2, 2, 1, 1, 1, 2, 2},
			Expected: 2,
		},
		{
			Name:     "Test Three",
			Nums:     []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 5, 5, 5, 5},
			Expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := MajorityElement_BruteForce(test.Nums)
			if result != test.Expected {
				t.Errorf("MajorityElement(%v) = %d, expected %d", test.Nums, result, test.Expected)
			}
		})

		t.Run(test.Name, func(t *testing.T) {
			result := MajorityElement_Better(test.Nums)
			if result != test.Expected {
				t.Errorf("MajorityElement(%v) = %d, expected %d", test.Nums, result, test.Expected)
			}
		})

		t.Run(test.Name, func(t *testing.T) {
			result := MajorityElement_Optimal(test.Nums)
			if result != test.Expected {
				t.Errorf("MajorityElement(%v) = %d, expected %d", test.Nums, result, test.Expected)
			}
		})
	}
}
