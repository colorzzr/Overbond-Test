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


func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
// function to compare the year of coorperate to the nearest two government bond
// and return the yield difference
func computeClosePointAndYield(c1 BondInfo, g1 BondInfo, g2 BondInfo) float64 {
	// conner case if the year is below the smallest one
	if g1.BondName == "G-1"{
		return c1.Yield - g2.Yield
	}


	d1 := math.Abs(c1.Year - g1.Year)
	d2 := math.Abs(c1.Year - g2.Year)

	// if we are closed to the year of g1
	if d1 < d2 {
		return c1.Yield - g1.Yield
	// else g2 is our target
	}else{
		return c1.Yield - g2.Yield
	}

	return -1
}

func findBestBenchmarkPoint(corpInfo BondInfo, govInfo []BondInfo, ){
	// conner case 0
	var i int
	for i = 0; i < len(govInfo); i++{
		if corpInfo.Year <= govInfo[i].Year{
			fmt.Println(corpInfo, govInfo[i-1], govInfo[i])
			ans := computeClosePointAndYield(corpInfo, govInfo[i-1], govInfo[i])
			// use round to remove floating error like 1.00000000001
			fmt.Println(Round(ans, 0.01))
			break
		}
	}
}


func main() {
	fmt.Println("test hello")

	var corpInfo []BondInfo
	var govInfo []BondInfo
	govInfo = append(govInfo, BondInfo{"G-1", true, -1 ,-1})

	csvFile, _ := os.Open("sample_input.csv")
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


	fmt.Println("------Print Government Bond------")
	for i := 0; i < len(govInfo);i++{
		fmt.Println(govInfo[i])
	}

	fmt.Println("------Print corperate Bond------")
	for i := 0; i < len(corpInfo);i++{
		fmt.Println(corpInfo[i])
	}

	fmt.Println("------Challange 1 test------")
	findBestBenchmarkPoint(corpInfo[2], govInfo)
	corpInfo[0].Year = 0.3
	findBestBenchmarkPoint(corpInfo[0], govInfo)
}