package main

import (
	"fmt"
)

var (
	input   = InputData{}
	foods   = Foods{}
	results = ResultsPerActivities{}
	mode = ActivityMode{}
	specialActivity float32 = 1.8
)

func main() {
	fmt.Println("Starting development my project nutrition calculator...")
	fillData()
	results = calculateDayFoodSize(mode.LowActivity, input.Weight, input.Age)

	fmt.Println("Activity", results.ActivityMode.Special)
	fmt.Println("KKal in proteins", results.KkalInProteins)
	fmt.Println("Kkal in fats", results.KkalInFats)
	fmt.Println("Kkal in carbohydrates", results.KkalInCarbohydrates)
	fmt.Println("kkal", results.KKalAll)

	results = calculateDayFoodSize(mode.AverageActivity, input.Weight, input.Age)

	fmt.Println("Activity", results.ActivityMode.Special)
	fmt.Println("KKal in proteins", results.KkalInProteins)
	fmt.Println("Kkal in fats", results.KkalInFats)
	fmt.Println("Kkal in carbohydrates", results.KkalInCarbohydrates)
	fmt.Println("kkal", results.KKalAll)
}

func calculateDayFoodSize(activityType float32, weight int32, age int32) ResultsPerActivities {
	results.ActivityMode.Special = activityType
	input.Weight = weight
	switch age := age;  {
	case age >= 18 && age <= 30:
		results.KKalAll = rounding((0.063*float32(input.Weight) + 2.8957) * 240 * activityType)
	case age >= 31 && age <= 60:
		results.KKalAll = rounding((0.0491*float32(input.Weight) + 2.4587) * 240)
	}
	results.KkalInProteins = results.KKalAll / 6
	results.KkalInFats = results.KKalAll / 6
	results.KkalInCarbohydrates = results.KKalAll / 6 * 4

	return results
}

func rounding(a float32) int32 {
	return int32(a)
}

func fillData() {
	var arr []Foods
	foods = Foods{"cheekenFile", 113, 23.6, 1.9, 0.4, false, false, true, false, false, false, false}
	arr = append(arr, foods)
	foods := Foods{"Salmon", 142, 19.8, 6.3, 0, false, false, false, true, false, false, false}
	arr = append(arr, foods)

	mode = ActivityMode{1, 1.3 , 1.5, specialActivity}
	input = InputData{92, 194, 30, mode}
}

type ResultsPerActivities struct {
	ActivityMode        ActivityMode
	KKalAll             int32
	KkalInProteins      int32
	KkalInFats          int32
	KkalInCarbohydrates int32
}

type InputData struct {
	Weight       int32
	Heigh        int32
	Age          int32
	ActivityMode ActivityMode
}

type ActivityMode struct {
	LowActivity     float32
	AverageActivity float32
	HighActivity    float32
	Special         float32
}
type TotalInAllFood struct {
	KKal          int
	Proteins      float32
	Fats          float32
	Carbohydrates float32
}

type Foods struct {
	Name          string
	KKal          int
	Proteins      float32
	Fats          float32
	Carbohydrates float32
	Fruits        bool
	Vegetables    bool
	Meat          bool
	Fish          bool
	Cereals       bool
	DairyProducts bool
	Eggs          bool
}

/*
Chicken fillet
Salmon
Boiled veal
Chicken egg, 1 pc.
Royal Cheese
Cottage cheese, 5%
Cereals

Hard Pasta
Buckwheat
Pollock fillet
Chinese cabbage
Beet
White cabbage
The apples
Oranges
Grapefruit
*/
