package main

import (
	"fmt"
	"math"
	"testing"
)


func Test_Load_CSV_1(t *testing.T){
	// try to load csv file below three are success
	_, _, err := loadCsvFile("./csv_test/sample_input.csv")
	if err != nil {
		t.Error("sample_input.csv should be able to load")
	}
	_, _, err = loadCsvFile("./csv_test/Q1_simple_1.csv")
	if err != nil {
		t.Error("Q1_simple_1.csv should be able to load")
	}
	_, _, err = loadCsvFile("./csv_test/Q2_simple_1.csv")
	if err != nil {
		t.Error("Q2_simple_1.csv should be able to load")
	}

	// load the file that not exist
	_, _, err = loadCsvFile("file_not_eixt")
	if err.Error() != "Fail to load file" {
		t.Error("file_not_eixt should not be loaded")
	}

	// try to load wrong format missing column
	_, _, err = loadCsvFile("./csv_test/wrong_format_1.csv")
	if err.Error() != "Document should have 4 columns" {
		fmt.Println(err)
		t.Error("wrong_format_1.csv has problem with column of data shouldnot be loaded")
	}

	// try to load the file that data is not proper fomate
	_, _, err = loadCsvFile("./csv_test/wrong_format_2.csv")
	if err.Error() != "invalid document format" {
		t.Error("wrong_format_1.csv has problem with data, should not be load")
	}


	t.Log("Finish Test_Load_CSV_1")
}

type ans struct{
	bond string
	yield float64
}

func Test_C1_Normal(t *testing.T){
	/*
	 *	the case that cover all normal situation
	 *	1. year is below the smallest government bond
	 *	2. yield is below the government bond
	 *	3. between the two bond
	 *	4. year is larger than the largest on
	 *	5. the bond is tie for both side then pick best yield
	 */
	 // G-1 and GMAX would be create by loadcsv file for removing conner case
	govBond := []BondInfo{	{"G-1", 0,0},
							{"G1", 1,1},
							{"G2", 5,5},
							{"GMAX", math.MaxFloat64 ,math.MaxFloat64}}
	corBond := []BondInfo{
							{"C1", 0.5, 0.5},
							{"C2", 3, 2},
							{"C3", 4.5, 1},
							{"C4", 7, 10}}

	expected := []ans{
						{"G1", -0.5},
						{"G1", 1},
						{"G2", -4},
						{"G2", 5}}

	var answer []ans
	for i := 0; i < len(corBond); i++{
		g1, g2 := findClosestTwoGoverBond(corBond[i], govBond)
		yield, bond := computeClosePointAndYield(corBond[i], g1, g2)
		answer = append(answer, ans{bond.BondName, yield})
	}

	// check length
	if len(answer) != len(expected){
		t.Error("answer is not enough")
		return
	}

	for i:=0; i <len(answer); i++{
		if (expected[i].yield != math.Round(answer[i].yield*100)/100) || (expected[i].bond != answer[i].bond){
			t.Error("at ", i," expected: ", expected[i], "but get: ", answer[i]);
		}
	}

	t.Log("finish Challenge 1 Normal")
}

func Test_C1_Hard(t *testing.T){
	govBond := []BondInfo{	{"G-1", 0,0},
		{"G1", 1,1},
		{"G2", 5,5},
		{"GMAX", math.MaxFloat64 ,math.MaxFloat64}}


	// if we have corperate bond near the max year
	corBond := []BondInfo{{"C1", math.MaxFloat64 - 1 ,7}}
	expected := []ans{{"G2", 2}}

	var answer []ans
	for i := 0; i < len(corBond); i++{
		g1, g2 := findClosestTwoGoverBond(corBond[i], govBond)
		yield, bond := computeClosePointAndYield(corBond[i], g1, g2)
		answer = append(answer, ans{bond.BondName, yield})
	}

	// check length
	if len(answer) != len(expected){
		t.Error("answer is not enough")
		return
	}

	for i:=0; i <len(answer); i++{
		if (expected[i].yield != math.Round(answer[i].yield*100)/100) || (expected[i].bond != answer[i].bond){
			t.Error("at ", i," expected: ", expected[i], "but get: ", answer[i]);
		}
	}


	// if we have cor bond have extact same year as one of government
	corBond = []BondInfo{{"C1", 5 ,-111}}
	expected = []ans{{"G2", -116}}
	answer = []ans{}

	for i := 0; i < len(corBond); i++{
		g1, g2 := findClosestTwoGoverBond(corBond[i], govBond)
		yield, bond := computeClosePointAndYield(corBond[i], g1, g2)
		answer = append(answer, ans{bond.BondName, yield})
	}

	// check length
	if len(answer) != len(expected){
		t.Error("answer is not enough")
		return
	}

	for i:=0; i <len(answer); i++{
		if (expected[i].yield != answer[i].yield) || (expected[i].bond != answer[i].bond){
			t.Error("at ", i," expected: ", expected[i], "but get: ", answer[i]);
		}
	}

	t.Log("Finish the Challenge 1 Hard")
}

func Test_C2_Normal(t *testing.T){
	// using same data of C1
	govBond := []BondInfo{	{"G-1", 0,0},
							{"G1", 1,1},
							{"G2", 5,5},
							{"GMAX", math.MaxFloat64 ,math.MaxFloat64}}
	corBond := []BondInfo{
							{"C1", 0.5, 0.5},
							{"C2", 3, 9},
							{"C3", 4.5, 1},
							{"C4", 7, 10}}

	expected := []ans{
						{"", 0},
						{"", 6},
						{"", -3.5},
						{"", 3}}


	var answer []ans
	for i := 0; i < len(corBond); i++{
		g1, g2 := findClosestTwoGoverBond(corBond[i], govBond)
		yield := linearApprox(corBond[i], g1, g2)
		answer = append(answer, ans{"", yield})
	}

	// check length
	if len(answer) != len(expected){
		t.Error("answer is not enough")
		return
	}

	for i:=0; i <len(answer); i++{
		if expected[i].yield != math.Round(answer[i].yield*100)/100{
			t.Error("at ", i," expected: ", expected[i], "but get: ", answer[i]);
		}
	}

	t.Log("finish Challenge 2 Normal")
}

func Test_C2_Hard(t *testing.T){
	// for challenge 2 here is not many extreme case because we linear Appoximation
	// at interval [0, max_float] so here is test for not unit slope
	govBond := []BondInfo{	{"G-1", 0,0},
							{"G1", 3,2.5},
							{"G11", 4,5},
							{"G2", 7,3.5},
							{"GMAX", math.MaxFloat64 ,10}}
	corBond := []BondInfo{
							{"C1", 0.5, 0.5},
							{"C2", 3, 5},
							{"C3", 4.5, 7},
							{"C4", 50, -1}}
	expected := []ans{
						{"", 0.08},
						{"", 2.5},
						{"", 2.25},
						{"", (-1) - (3.5 + (10 - 3.5)/(math.MaxFloat64 - 7) * (50 - 7))}}


	var answer []ans
	for i := 0; i < len(corBond); i++{
		g1, g2 := findClosestTwoGoverBond(corBond[i], govBond)
		yield := linearApprox(corBond[i], g1, g2)
		answer = append(answer, ans{"", yield})
	}

	// check length
	if len(answer) != len(expected){
		t.Error("answer is not enough")
		return
	}

	for i:=0; i <len(answer); i++{
		if expected[i].yield != (math.Round(answer[i].yield*100)/100){
			t.Error("at ", i," expected: ", expected[i], "but get: ", answer[i]);
		}
	}

	t.Log("finish Challenge 2 Hard")
}
