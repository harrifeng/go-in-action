package main

import "fmt"
import "github.com/harrifeng/go-in-action/pkg"

func main() {
	var qt pkg.QT
	fmt.Println(qt)
	qt = append(qt, 12)
	fmt.Println(qt)

	fmt.Println("--------------")
	var fhr pkg.FHR
	fmt.Println(fhr)
	// fhr is protected by struct !!
	// fhr.words = append(fhr.wwords, 12)
	fhr.Wwords = append(fhr.Wwords, 12)
	fmt.Println(fhr)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// []											  //
// [12]											  //
// --------------								  //
// {[] []}										  //
// {[12] []}									  //
////////////////////////////////////////////////////
