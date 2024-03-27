package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func main() {

	// VARIABLES TO CHANGE HERE FOR SENSITIVITY:
	// max_number_meets_half_day := 3
	// max_number_meets_full_day := 5
	// max_number_meets_full_week := 10

	/////
	//// this is the profiler that measures performance
	f, err := os.Create("myprogram.prof")
	if err != nil {
		fmt.Println(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// dont change these pls
	// av := import_availabilities()
	// create_columns(av)
	// remove_halfday_columns(max_number_meets_half_day, len(av))
	// make_full_days(max_number_meets_full_day, len(av))

	// make2dayWeeks_part1(max_number_meets_full_week, len(av))
	// make2dayWeeks_part2(max_number_meets_full_week, len(av))
	// make2dayWeeks_part3(max_number_meets_full_week, len(av))

	// link_4ab_with_3_3(max_number_meets_full_week, len(av))
	best_availabiltbiies := readLines(1, 3)

	best_bound := greedy(best_availabiltbiies)
	fmt.Println(best_bound)

}

func greedy(av [][]string) int {
	massString := ""
	fmt.Println(massString)
	// need to check that iter3_1 to iter3_50
	// massString_c1 = stringFeasibilitySingleDay(av, 0, 1, 2, 3, 4)
	// massString_c2 = stringFeasibilitySingleDay(av, 5, 6, 7, 8, 9)
	// massString := av[0] + av[1] + av[2] + av[3] + av[4]

	possibleLargeString := []string{}

	for i := 0; i < len(av); i++ {
		possibleLargeString = append(possibleLargeString, stringFeasibilitySingleDay(av, i))

	}
	return 0
}

func stringFeasibilitySingleDay(av [][]string, j1 int) string {
	massString := ""
	for i1 := 0; i1 < len(av[j1]); i1++ {
		// for i2 := 0; i2 < len(av[j2])-1; i2++ {
		// 	for i3 := 0; i3 < len(av[j3])-1; i3++ {
		// 		for i4 := 0; i4 < len(av[j4])-1; i4++ {
		// 			for i5 := 0; i5 < len(av[j5])-1; i5++ {
		fmt.Println(i1)
		massString = av[j1][i1] //+ av[j2][i2] + av[j3][i3] + av[j4][i4] + av[j5][i5]
		// 			}
		// 		}
		// 	}
		// }
	}
	return massString
}

func create_columns(av [][]int) {

	availabilities := av
	for i := 0; i < len(availabilities); i++ {
		fmt.Println("creating column:", i+1, "of", len(availabilities))
		// possibleColumns := [][]int{}
		outputFileName := fmt.Sprint("iter1_", i+1)
		outputFile, err := os.Create(outputFileName)

		if err != nil {
			fmt.Println(err)
		}
		indexes := find_index_of_ones(availabilities[i])
		// check(availabilities[i], len(availabilities[i]))

		for j := 0; j <= int(math.Pow(float64(2), float64(len(indexes))))-1; j++ {
			slice := (generate_slice(j, len(indexes)))
			replacement := replace(indexes, availabilities[i], slice)

			// fmt.Println(replacement)
			// possibleColumns = append(possibleColumns, replacement)

			strReplacement := make([]string, len(replacement))
			for k, x := range replacement {
				strReplacement[k] = strconv.Itoa(x)
			}

			stringthing := strings.Join(strReplacement, "")
			outputFile.WriteString(stringthing)
			outputFile.WriteString("\n")

		}

	}
}

func find_index_of_ones(av []int) []int {
	temp := make([]int, 0)
	for i := 0; i <= len(av)-1; i++ {
		if av[i] == 1 {
			temp = append(temp, i)
		}
	}
	return temp
}

func replace(indexes []int, av []int, slice []int) []int {
	for i := 0; i <= len(indexes)-1; i++ {

		index_to_change := indexes[i]

		av[index_to_change] = slice[i]

	}
	// fmt.Println("what was returned:", av)
	return av
}

func generate_slice(value int, size int) []int {

	powers := []uint64{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648, 4294967296, 8589934592, 17179869184}
	// this returns columns, eg column with [1 0 1 0] corresponds to h = 10

	slice := make([]int, size)
	valueAsUint := uint64(value)
	initialValue := valueAsUint

	for countdown := size - 1; countdown >= 0; countdown-- {
		holder := powers[countdown]

		if initialValue >= holder {
			// then we need to remove something off it
			initialValue = initialValue - holder
			slice[size-countdown-1] = 1
		}
	}

	return slice
}

func readLines(startPosition int, endPosition int) [][]string {

	stringInterval := [][]string{}

	// https://stackoverflow.com/questions/35704948/generate-array-with-with-range-of-integers
	lo, hi := startPosition, endPosition
	s := make([]int, hi-lo+1)
	for i := range s {
		s[i] = i + lo
	}

	// we want to read from
	for i := 1; i <= 5; i++ {
		// reading := []string{}
		fileName := "reverse_iter3_" + strconv.Itoa(i)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		scanner := bufio.NewScanner(file)

		holder := []string{}
		// Keep track of the current line number
		lineNumber := 1
		fmt.Println(fileName)
		// Iterate through each line in the file
		for scanner.Scan() {
			// Check if the current line number is in the list of lines to read
			for _, line := range s {
				if lineNumber == line {
					// Print the line if it's in the list of lines to read
					holder = append(holder, scanner.Text())
					break // Move to the next line
				}
			}

			// Increment the line number
			lineNumber++
		}
		stringInterval = append(stringInterval, holder)

	}
	return stringInterval
}

func remove_halfday_columns(max_number_meets_half_day int, num_of_availabilities_processing int) {

	// grab files haphazardly

	for i := 1; i <= num_of_availabilities_processing; i++ {
		fmt.Println("adding halfday constraint:", i, "of", num_of_availabilities_processing)
		validLines := []string{}

		looking_for_file := "iter1_" + strconv.Itoa(i)
		file, err := os.Open(looking_for_file)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			num_meetings_halfday := strings.Count(line, "1")

			// if within tolernace then add it
			if num_meetings_halfday <= max_number_meets_half_day {
				validLines = append(validLines, line)
			}
		}

		// Check for scanning errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}

		outputFileName := fmt.Sprint("iter2_", i)
		outputFile, err := os.Create(outputFileName)

		for j := range validLines {
			outputFile.WriteString(validLines[j])
			outputFile.WriteString("\n")
		}

	}

}

func make_full_days(max_number_meets_full_day int, num_of_availabilities_processing int) {

	// read the entire file in
	for i := 0; i < (num_of_availabilities_processing / 2); i++ {
		fmt.Println("adding fullday constraint:", i, "of", num_of_availabilities_processing/2)
		morningFilename := "iter2_" + strconv.Itoa(2*i+1)
		afternoonFilename := "iter2_" + strconv.Itoa(2*i+2)
		contents_morning_file := openFileReadAll(morningFilename)
		contents_afternoon_file := openFileReadAll(afternoonFilename)

		validFulldayColumn := []string{}

		for j := 0; j < len(contents_morning_file); j++ {
			for k := 0; k < len(contents_afternoon_file); k++ {

				possibleFulldayColumns := contents_morning_file[j] + contents_afternoon_file[k]
				num_meetings_fullday := strings.Count(possibleFulldayColumns, "1")

				if num_meetings_fullday <= max_number_meets_full_day {
					validFulldayColumn = append(validFulldayColumn, possibleFulldayColumns)
				}

			}
		}
		outputFileName := fmt.Sprint("iter3_", i+1)
		outputFile, _ := os.Create(outputFileName)

		//NOTE: THIS KEEPS FLUSHING SO IT'S VERY BAD!!!
		for j := range validFulldayColumn {
			outputFile.WriteString(validFulldayColumn[j])
			outputFile.WriteString("\n")
		}
	}

}

//=====

func make2dayWeeks_part1(max_number_meets_full_week int, num_of_availabilities_processing int) {
	// each participant creates sheets [iter3_1 to iter3_5]
	// we know these are monday to friday
	// we can play around with set covering to remove most of them :)

	number_of_people := num_of_availabilities_processing / 10

	for i := 0; i < number_of_people; i++ {
		monFilename := "iter3_" + strconv.Itoa(2*i+1)
		tueFilename := "iter3_" + strconv.Itoa(2*i+2)
		contents_mon := openFileReadAll(monFilename)
		contents_tue := openFileReadAll(tueFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_ab_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_ab_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_mon); j1++ {
			for j2 := 0; j2 < len(contents_tue); j2++ {

				possibleMonTues := contents_mon[j1] + contents_tue[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_mon[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_tue[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4ab...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		monFilename := "iter3_" + strconv.Itoa(2*i+1)
		wedFilename := "iter3_" + strconv.Itoa(2*i+3)
		contents_mon := openFileReadAll(monFilename)
		contents_wed := openFileReadAll(wedFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_ac_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_ac_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_mon); j1++ {
			for j2 := 0; j2 < len(contents_wed); j2++ {

				possibleMonTues := contents_mon[j1] + contents_wed[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_mon[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_wed[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4ac...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		monFilename := "iter3_" + strconv.Itoa(2*i+1)
		thuFilename := "iter3_" + strconv.Itoa(2*i+4)
		contents_mon := openFileReadAll(monFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_ad_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_ad_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_mon); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_mon[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_mon[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4ad...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		monFilename := "iter3_" + strconv.Itoa(2*i+1)
		friFilename := "iter3_" + strconv.Itoa(2*i+5)
		contents_mon := openFileReadAll(monFilename)
		contents_fri := openFileReadAll(friFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_ae_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_ae_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_mon); j1++ {
			for j2 := 0; j2 < len(contents_fri); j2++ {

				possibleMonTues := contents_mon[j1] + contents_fri[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_mon[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_fri[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4ae...")
		w1.Flush()
		w2.Flush()
	}
}

func make2dayWeeks_part2(max_number_meets_full_week int, num_of_availabilities_processing int) {
	// each participant creates sheets [iter3_1 to iter3_5]
	// we know these are monday to friday
	// we can play around with set covering to remove most of them :)

	number_of_people := num_of_availabilities_processing / 10

	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+2)
		wedFilename := "iter3_" + strconv.Itoa(2*i+3)
		contents_tue := openFileReadAll(tueFilename)
		contents_wed := openFileReadAll(wedFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_bc_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_bc_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_wed); j2++ {

				possibleMonTues := contents_tue[j1] + contents_wed[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_wed[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4bc...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+2)
		thuFilename := "iter3_" + strconv.Itoa(2*i+4)
		contents_tue := openFileReadAll(tueFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_bd_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_bd_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_tue[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4bd...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+2)
		thuFilename := "iter3_" + strconv.Itoa(2*i+5)
		contents_tue := openFileReadAll(tueFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_be_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_be_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_tue[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4be...")
		w1.Flush()
		w2.Flush()
	}
}

func make2dayWeeks_part3(max_number_meets_full_week int, num_of_availabilities_processing int) {
	// each participant creates sheets [iter3_1 to iter3_5]
	// we know these are monday to friday
	// we can play around with set covering to remove most of them :)

	number_of_people := num_of_availabilities_processing / 10
	/// for c
	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+2)
		thuFilename := "iter3_" + strconv.Itoa(2*i+4)
		contents_tue := openFileReadAll(tueFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_cd_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_cd_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_tue[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4cd...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+3)
		thuFilename := "iter3_" + strconv.Itoa(2*i+5)
		contents_tue := openFileReadAll(tueFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_ce_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_ce_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()
		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_tue[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4ce...")
		w1.Flush()
		w2.Flush()
	}

	for i := 0; i < number_of_people; i++ {
		tueFilename := "iter3_" + strconv.Itoa(2*i+4)
		thuFilename := "iter3_" + strconv.Itoa(2*i+5)
		contents_tue := openFileReadAll(tueFilename)
		contents_thu := openFileReadAll(thuFilename)

		output1FileName := fmt.Sprint("iter4_", i+1, "_de_1")
		output2FileName := fmt.Sprint("iter4_", i+1, "_de_2")
		outputFile1, _ := os.Create(output1FileName)
		outputFile2, _ := os.Create(output2FileName)
		w1 := bufio.NewWriter(outputFile1)
		w2 := bufio.NewWriter(outputFile2)

		outputFile1.Sync()
		outputFile2.Sync()

		// validFullweekColumn := []string{}
		for j1 := 0; j1 < len(contents_tue); j1++ {
			for j2 := 0; j2 < len(contents_thu); j2++ {

				possibleMonTues := contents_tue[j1] + contents_thu[j2]
				num_meetings_week := strings.Count(possibleMonTues, "1")
				if num_meetings_week <= max_number_meets_full_week {
					possible1, _ := strconv.ParseInt(contents_tue[j1], 2, 0)
					possible1_1 := strconv.FormatInt(possible1, 36)
					possible2, _ := strconv.ParseInt(contents_thu[j2], 2, 0)
					possible2_1 := strconv.FormatInt(possible2, 36)
					w1.WriteString(possible1_1)
					w1.WriteString("\n")
					w2.WriteString(possible2_1)
					w2.WriteString("\n")
				}
			}
		}
		fmt.Println("started writing 4de...")
		w1.Flush()
		w2.Flush()
	}

}

func link_4ab_with_3_3(max_number_meets_full_week int, num_of_availabilities_processing int) {
	number_of_people := num_of_availabilities_processing / 10
	for i := 0; i < number_of_people; i++ {
		// iter4abPart1_Filename := "iter4_" + strconv.Itoa(2*i+1) + "_ab_1"
		iter4abPart2_Filename := "iter4_" + strconv.Itoa(2*i+1) + "_ab_2"
		iter4bcPart1_Filename := "iter4_" + strconv.Itoa(2*i+1) + "_bc_1"
		// iter4acPart2_Filename := "iter4_" + strconv.Itoa(2*i+1) + "_bc_2"

		// contents_4ab_pt1 := openFileReadAll(iter4abPart1_Filename)
		contents_4ab_pt2 := openFileReadAll(iter4abPart2_Filename)
		contents_4bc_pt1 := openFileReadAll(iter4bcPart1_Filename)
		// contents_4bc_pt2 := openFileReadAll(iter4acPart2_Filename)

		// outputFileName := fmt.Sprint("iter4_", i+1, "_ab")
		// outputFile, _ := os.Create(outputFileName)
		// w := bufio.NewWriter(outputFile)

		// outputFile.Sync()

		// validMorning := [][]string{}

		// fmt.Println(result)

		validB := []int{}
		for j1 := 0; j1 < len(contents_4ab_pt2); j1++ {
			fmt.Println("j1 is", j1, "of", len(contents_4ab_pt2))
			for j2 := 0; j2 < len(contents_4bc_pt1); j2++ {

				if contents_4ab_pt2[j1] == contents_4bc_pt1[j2] {
					validB = append(validB, j1)
				}

				// fmt.Println("j1", j1+1, "of", len(contents_4ab_pt1))
				// fmt.Println("j2", j2+1, "of", len(contents_4bc_pt1))

			}
		}
		fmt.Println(validB)
	}
}

// func greedy()

func openFileReadAll(fileName string) []string {
	fileContents := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileContents = append(fileContents, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	return fileContents
}

func import_availabilities() [][]int {

	c1_mon_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0}
	c1_mon_2 := []int{0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c1_tue_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c1_tue_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c1_wed_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0}
	c1_wed_2 := []int{0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c1_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c1_thu_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c1_fri_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c1_fri_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c2_mon_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	c2_mon_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	c2_tue_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	c2_tue_2 := []int{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	c2_wed_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	c2_wed_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	c2_thu_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	c2_thu_2 := []int{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	c2_fri_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	c2_fri_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	c3_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	c3_mon_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c3_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c3_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c3_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	c3_wed_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c3_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c3_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c3_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c3_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c4_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c4_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c4_tue_1 := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1}
	c4_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c4_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c4_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c4_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c4_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c4_fri_1 := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1}
	c4_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c5_mon_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c5_mon_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c5_tue_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c5_tue_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c5_wed_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c5_wed_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c5_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	c5_thu_2 := []int{0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c5_fri_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c5_fri_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c6_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c6_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c6_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c6_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c6_wed_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	c6_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c6_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c6_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c6_fri_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	c6_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c7_mon_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c7_mon_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c7_tue_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c7_tue_2 := []int{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	c7_wed_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c7_wed_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c7_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c7_thu_2 := []int{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
	c7_fri_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	c7_fri_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	c8_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c8_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c8_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c8_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c8_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c8_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c8_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	c8_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c8_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c8_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c9_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c9_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c9_tue_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	c9_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c9_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c9_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c9_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c9_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c9_fri_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	c9_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c10_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	c10_mon_2 := []int{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c10_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c10_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c10_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c10_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c10_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	c10_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c10_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	c10_fri_2 := []int{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m1_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m1_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m1_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m1_tue_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m1_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m1_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m1_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m1_thu_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m1_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m1_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m2_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m2_mon_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m2_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m2_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m2_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m2_wed_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m2_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m2_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m2_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m2_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m3_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m3_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m3_tue_1 := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1}
	m3_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m3_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m3_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m3_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m3_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m3_fri_1 := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1}
	m3_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m4_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1}
	m4_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m4_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m4_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m4_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m4_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m4_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1}
	m4_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m4_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m4_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m5_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m5_mon_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m5_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m5_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m5_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m5_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m5_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m5_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m5_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m5_fri_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m6_mon_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m6_mon_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1}
	m6_tue_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m6_tue_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m6_wed_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m6_wed_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m6_thu_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m6_thu_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m6_fri_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m6_fri_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m7_mon_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}
	m7_mon_2 := []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m7_tue_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m7_tue_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m7_wed_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m7_wed_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m7_thu_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m7_thu_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m7_fri_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m7_fri_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m8_mon_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}
	m8_mon_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m8_tue_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m8_tue_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m8_wed_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}
	m8_wed_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m8_thu_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m8_thu_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m8_fri_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m8_fri_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m9_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m9_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m9_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m9_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m9_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m9_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m9_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m9_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m9_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m9_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m10_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m10_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m10_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m10_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m10_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m10_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m10_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m10_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m10_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1}
	m10_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m11_mon_1 := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0}
	m11_mon_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m11_tue_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m11_tue_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m11_wed_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m11_wed_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m11_thu_1 := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0}
	m11_thu_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m11_fri_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	m11_fri_2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	m12_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m12_mon_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m12_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m12_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m12_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m12_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m12_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m12_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m12_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m12_fri_2 := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m13_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m13_mon_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m13_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m13_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m13_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m13_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m13_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m13_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m13_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m13_fri_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m14_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m14_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m14_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m14_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m14_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}
	m14_wed_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m14_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m14_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m14_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m14_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m15_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m15_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m15_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m15_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m15_wed_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m15_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m15_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m15_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m15_fri_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m15_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m16_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m16_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m16_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m16_tue_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m16_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m16_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m16_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m16_thu_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m16_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m16_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m17_mon_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m17_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m17_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m17_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m17_wed_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m17_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m17_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m17_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m17_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m17_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m18_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m18_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m18_tue_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m18_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m18_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m18_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m18_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m18_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m18_fri_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1}
	m18_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m19_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m19_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m19_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m19_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m19_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m19_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m19_thu_1 := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m19_thu_2 := []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m19_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m19_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m20_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m20_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m20_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m20_tue_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m20_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m20_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m20_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m20_thu_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m20_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m20_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m21_mon_1 := []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m21_mon_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m21_tue_1 := []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m21_tue_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m21_wed_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m21_wed_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m21_thu_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m21_thu_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m21_fri_1 := []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m21_fri_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m22_mon_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	m22_mon_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	m22_tue_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	m22_tue_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	m22_wed_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	m22_wed_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	m22_thu_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m22_thu_2 := []int{0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	m22_fri_1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1}
	m22_fri_2 := []int{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	m23_mon_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}
	m23_mon_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m23_tue_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0}
	m23_tue_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m23_wed_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m23_wed_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m23_thu_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m23_thu_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m23_fri_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m23_fri_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m24_mon_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m24_mon_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m24_tue_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m24_tue_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1}
	m24_wed_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m24_wed_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1}
	m24_thu_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m24_thu_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m24_fri_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m24_fri_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m25_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m25_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m25_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m25_tue_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m25_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m25_wed_2 := []int{1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m25_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m25_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m25_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m25_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m26_mon_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m26_mon_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m26_tue_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m26_tue_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m26_wed_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0}
	m26_wed_2 := []int{0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m26_thu_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m26_thu_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m26_fri_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	m26_fri_2 := []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	m27_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	m27_mon_2 := []int{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m27_tue_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m27_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m27_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m27_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m27_thu_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m27_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m27_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	m27_fri_2 := []int{0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m28_mon_1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m28_mon_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m28_tue_1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m28_tue_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m28_wed_1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m28_wed_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m28_thu_1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m28_thu_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m28_fri_1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}
	m28_fri_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m29_mon_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m29_mon_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m29_tue_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m29_tue_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m29_wed_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m29_wed_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m29_thu_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	m29_thu_2 := []int{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m29_fri_1 := []int{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0}
	m29_fri_2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}
	m30_mon_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m30_mon_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m30_tue_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m30_tue_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m30_wed_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m30_wed_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m30_thu_1 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m30_thu_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m30_fri_1 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}
	m30_fri_2 := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	slice := [][]int{c1_mon_1, c1_mon_2, c1_tue_1, c1_tue_2, c1_wed_1, c1_wed_2, c1_thu_1, c1_thu_2, c1_fri_1, c1_fri_2, c2_mon_1, c2_mon_2, c2_tue_1, c2_tue_2, c2_wed_1, c2_wed_2, c2_thu_1, c2_thu_2, c2_fri_1, c2_fri_2, c3_mon_1, c3_mon_2, c3_tue_1, c3_tue_2, c3_wed_1, c3_wed_2, c3_thu_1, c3_thu_2, c3_fri_1, c3_fri_2, c4_mon_1, c4_mon_2, c4_tue_1, c4_tue_2, c4_wed_1, c4_wed_2, c4_thu_1, c4_thu_2, c4_fri_1, c4_fri_2, c5_mon_1, c5_mon_2, c5_tue_1, c5_tue_2, c5_wed_1, c5_wed_2, c5_thu_1, c5_thu_2, c5_fri_1, c5_fri_2, c6_mon_1, c6_mon_2, c6_tue_1, c6_tue_2, c6_wed_1, c6_wed_2, c6_thu_1, c6_thu_2, c6_fri_1, c6_fri_2, c7_mon_1, c7_mon_2, c7_tue_1, c7_tue_2, c7_wed_1, c7_wed_2, c7_thu_1, c7_thu_2, c7_fri_1, c7_fri_2, c8_mon_1, c8_mon_2, c8_tue_1, c8_tue_2, c8_wed_1, c8_wed_2, c8_thu_1, c8_thu_2, c8_fri_1, c8_fri_2, c9_mon_1, c9_mon_2, c9_tue_1, c9_tue_2, c9_wed_1, c9_wed_2, c9_thu_1, c9_thu_2, c9_fri_1, c9_fri_2, c10_mon_1, c10_mon_2, c10_tue_1, c10_tue_2, c10_wed_1, c10_wed_2, c10_thu_1, c10_thu_2, c10_fri_1, c10_fri_2, m1_mon_1, m1_mon_2, m1_tue_1, m1_tue_2, m1_wed_1, m1_wed_2, m1_thu_1, m1_thu_2, m1_fri_1, m1_fri_2, m2_mon_1, m2_mon_2, m2_tue_1, m2_tue_2, m2_wed_1, m2_wed_2, m2_thu_1, m2_thu_2, m2_fri_1, m2_fri_2, m3_mon_1, m3_mon_2, m3_tue_1, m3_tue_2, m3_wed_1, m3_wed_2, m3_thu_1, m3_thu_2, m3_fri_1, m3_fri_2, m4_mon_1, m4_mon_2, m4_tue_1, m4_tue_2, m4_wed_1, m4_wed_2, m4_thu_1, m4_thu_2, m4_fri_1, m4_fri_2, m5_mon_1, m5_mon_2, m5_tue_1, m5_tue_2, m5_wed_1, m5_wed_2, m5_thu_1, m5_thu_2, m5_fri_1, m5_fri_2, m6_mon_1, m6_mon_2, m6_tue_1, m6_tue_2, m6_wed_1, m6_wed_2, m6_thu_1, m6_thu_2, m6_fri_1, m6_fri_2, m7_mon_1, m7_mon_2, m7_tue_1, m7_tue_2, m7_wed_1, m7_wed_2, m7_thu_1, m7_thu_2, m7_fri_1, m7_fri_2, m8_mon_1, m8_mon_2, m8_tue_1, m8_tue_2, m8_wed_1, m8_wed_2, m8_thu_1, m8_thu_2, m8_fri_1, m8_fri_2, m9_mon_1, m9_mon_2, m9_tue_1, m9_tue_2, m9_wed_1, m9_wed_2, m9_thu_1, m9_thu_2, m9_fri_1, m9_fri_2, m10_mon_1, m10_mon_2, m10_tue_1, m10_tue_2, m10_wed_1, m10_wed_2, m10_thu_1, m10_thu_2, m10_fri_1, m10_fri_2, m11_mon_1, m11_mon_2, m11_tue_1, m11_tue_2, m11_wed_1, m11_wed_2, m11_thu_1, m11_thu_2, m11_fri_1, m11_fri_2, m12_mon_1, m12_mon_2, m12_tue_1, m12_tue_2, m12_wed_1, m12_wed_2, m12_thu_1, m12_thu_2, m12_fri_1, m12_fri_2, m13_mon_1, m13_mon_2, m13_tue_1, m13_tue_2, m13_wed_1, m13_wed_2, m13_thu_1, m13_thu_2, m13_fri_1, m13_fri_2, m14_mon_1, m14_mon_2, m14_tue_1, m14_tue_2, m14_wed_1, m14_wed_2, m14_thu_1, m14_thu_2, m14_fri_1, m14_fri_2, m15_mon_1, m15_mon_2, m15_tue_1, m15_tue_2, m15_wed_1, m15_wed_2, m15_thu_1, m15_thu_2, m15_fri_1, m15_fri_2, m16_mon_1, m16_mon_2, m16_tue_1, m16_tue_2, m16_wed_1, m16_wed_2, m16_thu_1, m16_thu_2, m16_fri_1, m16_fri_2, m17_mon_1, m17_mon_2, m17_tue_1, m17_tue_2, m17_wed_1, m17_wed_2, m17_thu_1, m17_thu_2, m17_fri_1, m17_fri_2, m18_mon_1, m18_mon_2, m18_tue_1, m18_tue_2, m18_wed_1, m18_wed_2, m18_thu_1, m18_thu_2, m18_fri_1, m18_fri_2, m19_mon_1, m19_mon_2, m19_tue_1, m19_tue_2, m19_wed_1, m19_wed_2, m19_thu_1, m19_thu_2, m19_fri_1, m19_fri_2, m20_mon_1, m20_mon_2, m20_tue_1, m20_tue_2, m20_wed_1, m20_wed_2, m20_thu_1, m20_thu_2, m20_fri_1, m20_fri_2, m21_mon_1, m21_mon_2, m21_tue_1, m21_tue_2, m21_wed_1, m21_wed_2, m21_thu_1, m21_thu_2, m21_fri_1, m21_fri_2, m22_mon_1, m22_mon_2, m22_tue_1, m22_tue_2, m22_wed_1, m22_wed_2, m22_thu_1, m22_thu_2, m22_fri_1, m22_fri_2, m23_mon_1, m23_mon_2, m23_tue_1, m23_tue_2, m23_wed_1, m23_wed_2, m23_thu_1, m23_thu_2, m23_fri_1, m23_fri_2, m24_mon_1, m24_mon_2, m24_tue_1, m24_tue_2, m24_wed_1, m24_wed_2, m24_thu_1, m24_thu_2, m24_fri_1, m24_fri_2, m25_mon_1, m25_mon_2, m25_tue_1, m25_tue_2, m25_wed_1, m25_wed_2, m25_thu_1, m25_thu_2, m25_fri_1, m25_fri_2, m26_mon_1, m26_mon_2, m26_tue_1, m26_tue_2, m26_wed_1, m26_wed_2, m26_thu_1, m26_thu_2, m26_fri_1, m26_fri_2, m27_mon_1, m27_mon_2, m27_tue_1, m27_tue_2, m27_wed_1, m27_wed_2, m27_thu_1, m27_thu_2, m27_fri_1, m27_fri_2, m28_mon_1, m28_mon_2, m28_tue_1, m28_tue_2, m28_wed_1, m28_wed_2, m28_thu_1, m28_thu_2, m28_fri_1, m28_fri_2, m29_mon_1, m29_mon_2, m29_tue_1, m29_tue_2, m29_wed_1, m29_wed_2, m29_thu_1, m29_thu_2, m29_fri_1, m29_fri_2, m30_mon_1, m30_mon_2, m30_tue_1, m30_tue_2, m30_wed_1, m30_wed_2, m30_thu_1, m30_thu_2, m30_fri_1, m30_fri_2}

	return slice

}
