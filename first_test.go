package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	testcases := []struct {
		res         string
		expectedRes string
	}{
		{"res.txt", "expectedRes.txt"},
		// {"res1.txt", "expectedRes1.txt"},
		// {"res2.txt", "expectedRes2.txt"},
		// {"res3.txt", "expectedRes3.txt"},
		// {"res4.txt", "expectedRes4.txt"},
		// {"res5.txt", "expectedRes5.txt"},
		// {"res6.txt", "expectedRes6.txt"},
	}
	for _, tc := range testcases {
		res_file, err := os.ReadFile(tc.res)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		resStr := string(res_file)
		expected_res, err := os.ReadFile(tc.expectedRes)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		expectedStr := string(expected_res)
		if resStr != expectedStr {
			t.Errorf("For result file %s:\nExpected\n%s\nbut got\n%s", tc.res, expectedStr, resStr)
		}
	}
}
