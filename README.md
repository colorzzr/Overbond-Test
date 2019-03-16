# Overbond-Test
Overbond-Test

Problem requires to analysis the corperate bond and govornment bond

I decide to use the golang to provide solution.

since function need to process large amount of data. So the pressure test is need to get the speed in the future. and the linear search may change to reduce the time complexity from O(n) to O(logn)


## structure

```
type BondInfo struct{
	BondName string
	Year float64
	Yield float64
}
```

## API document

`func loadCsvFile(fileName string)([]BondInfo, []BondInfo, error)`

 - this function would scan the target string and return two array for corperate and government each. the array would be sorted for further usage. it would append two dummy node `{"G-1", 0, 0}` and `{"GMAX", Max_Float, Max_Float}` to solve the conner case

 - Error control: 

 		- fatal when cannot open file
 		- fatal when csv has not enough column for computation
 		- fatal when csn has invalid data

---

`func findBestBenchmarkPoint(corpInfo []BondInfo, govInfo []BondInfo)`

- challenge 1 function. it would find the government bond that has closest year to the target corperate bond

- find the interval that `govInfo[i - 1].year` <= `corpInfo.year` <= `govInfo[i].year`. Then compare which one has closer year to the corpInfo


---

`func findYieldInCurve(corpInfo []BondInfo, govInfo []BondInfo)`

- challenge 2 function. it would find the interval that `govInfo[i - 1].year` <= `corpInfo.year` <= `govInfo[i].year`. Then user linear approximation to find the `yield` at `corpInfo.year`. then return the `yield - corpInfo.year`


# Test Case

Test cover the most case that I can think of.



