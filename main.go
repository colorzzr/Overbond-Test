package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/bradfitz/slice"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type BondInfo struct{
	BondName string
	// if true it is government else corpoerate
	BondType bool
	Year float64
	Yield float64
}

//--------------------------------------------Challenge 1--------------------------------------
// function to compare the year of coorperate to the nearest two government bond
// and return the yield difference and cloest government bond
func computeClosePointAndYield(c1 BondInfo, g1 BondInfo, g2 BondInfo) (float64, BondInfo) {
	// conner case if the year is below the smallest one
	if g1.BondName == "G-1"{
		return c1.Yield - g2.Yield, g2
	}


	d1 := math.Abs(c1.Year - g1.Year)
	d2 := math.Abs(c1.Year - g2.Year)

	// if we are closed to the year of g1
	if d1 < d2 {
		return c1.Yield - g1.Yield, g1
	// else g2 is our target
	}else{
		return c1.Yield - g2.Yield, g2
	}

	return -1, BondInfo{}
}

// return the interval that target corperate bond between
func findClosestTwoGoverBond(corpInfo BondInfo, govInfo []BondInfo)(BondInfo, BondInfo){
	// conner case 0
	for i := 0; i < len(govInfo); i++{
		if corpInfo.Year <= govInfo[i].Year{
			//fmt.Println(corpInfo, govInfo[i-1], govInfo[i])
			//ans := computeClosePointAndYield(corpInfo, govInfo[i-1], govInfo[i])
			//// use round to remove floating error like 1.00000000001
			//fmt.Println(Round(ans, 0.01))


			return govInfo[i-1], govInfo[i]
		}
	}

	return BondInfo{}, BondInfo{}
}

// search the best government bond for target corperated info
func findBestBenchmarkPoint(corpInfo []BondInfo, govInfo []BondInfo){

	fmt.Println("bond,benchmark,spread_to_benchmark")
	for i := 0; i < len(corpInfo); i++{
		// find interval
		g1, g2 := findClosestTwoGoverBond(corpInfo[i], govInfo)
		// compute the yield to closest one
		yield, bond := computeClosePointAndYield(corpInfo[i], g1, g2)

		fmt.Print(corpInfo[i].BondName, ",", bond.BondName, ",", math.Round(yield*100)/100 ,"%\n")
	}
}

// --------------------------------------------Challenge 2-------------------------------------------------

// we use g1.yield + slope * delta of Year
func linearApprox(c1 BondInfo, g1 BondInfo, g2 BondInfo) float64 {
	slope := (g2.Yield - g1.Yield) / (g2.Year - g1.Year)
	// get the change of year to compute yield by slope
	dYear := c1.Year - g1.Year

	return c1.Yield - (g1.Yield + dYear * slope)
}

// find the estimate value of yield using linear approximation on two
// closest government bond
func findYieldInCurve(corpInfo []BondInfo, govInfo []BondInfo){

	fmt.Println("bond,spread_to_curve")
	for i := 0; i < len(corpInfo); i++{
		// find interval
		g1, g2 := findClosestTwoGoverBond(corpInfo[i], govInfo)
		// project on to the line
		yield := linearApprox(corpInfo[i], g1, g2)

		fmt.Print(corpInfo[i].BondName, "," , math.Round(yield*100)/100, "%\n")
	}

}

// load the target file and return sorted array of corperate bond and government bond
func loadCsvFile(fileName string)([]BondInfo, []BondInfo){
	corpInfo :=  make([]BondInfo, 0)
	govInfo := make([]BondInfo, 0)
	govInfo = append(govInfo, BondInfo{"G-1", true, 0 ,0})

	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// skip the heading
	line, error := reader.Read()
	for {
		line, error = reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var temp BondInfo
		// get the name
		temp.BondName = line[0];
		// filter the year to float
		s := strings.Split(line[2], " ");
		//fmt.Println(s);
		tempFloat, err := strconv.ParseFloat(s[0], 64)
		if err != nil{
			log.Fatal(err)
		}
		temp.Year = tempFloat;

		//filter out the yield
		s = strings.Split(line[3], "%");
		tempFloat, err = strconv.ParseFloat(s[0], 64)
		if err != nil{
			log.Fatal(err)
		}
		temp.Yield = tempFloat;

		// get the type
		if line[1] == "corporate" {
			temp.BondType = false
			corpInfo = append(corpInfo, temp)
		}else {
			temp.BondType = true
			govInfo = append(govInfo, temp)
		}

		// append to new
		//info = append(info, temp)
		//fmt.Println(line[0], line[1], line[2], line[3])
	}

	// remove it
	slice.Sort(corpInfo[:], func(i, j int) bool {
		return corpInfo[i].Year < corpInfo[j].Year
	})
	// sort grovernmet array
	slice.Sort(govInfo[:], func(i, j int) bool {
		return govInfo[i].Year < govInfo[j].Year
	})

	return corpInfo, govInfo
}


func main() {
	//fmt.Println("test hello")



	corpInfo, govInfo := loadCsvFile("Q2_simple_1.csv")


	fmt.Println("------Print Government Bond------")
	for i := 0; i < len(govInfo);i++{
		fmt.Println(govInfo[i])
	}

	fmt.Println("------Print corperate Bond------")
	for i := 0; i < len(corpInfo);i++{
		fmt.Println(corpInfo[i])
	}

	fmt.Println("------Challange 1 test------")
	findYieldInCurve(corpInfo, govInfo)
}