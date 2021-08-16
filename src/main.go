// Copyright 2021 Tyler Calder
package main

import (
	"flag"
	"fmt"
	"os"
)

// Maps Tax bracket lower bound to tax rate
// TODO: This is the Married Filing Jointly Tax Bracket.
// It would be best to make this something that we can
// Accept as an input so that it is easily customizable
var bracket_limits = []float64{
	628301.0,
	418851.0,
	329851.0,
	172751.0,
	81051.0,
	19901.0,
	0.0,
}

var tax_brackets = map[float64]float64{
	0.0:      .10,
	19901.0:  .12,
	81051.0:  .22,
	172751.0: .24,
	329851.0: .34,
	418851.0: .35,
	628301.0: .37,
}

var income = flag.Float64("i", 0, "Your income you want to calculate for")
var help = flag.Bool("h", false, "Display this help message")

func main() {
	// Calculates your taxable income based off of rates and
	// values and tax rates
	flag.Parse()
	if *help {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	if *income <= 0.0 {
		fmt.Printf("Your income is %.2f?!? Get a job you dirty hippy!\n", *income)
	}
	tax_inc := calc_taxable_income(*income)
	tax_burden := calc_tax_burden(tax_inc)
	fmt.Printf("Your taxable income is $%.2f\n", tax_inc)
	fmt.Printf("Your tax burden is $%.2f\n", tax_burden)
}

func calc_taxable_income(gross float64) float64 {
	// Taxable Income = income - deductions
	// TODO: For now assuming standard deduction for married filing joinlty
	return gross - 25100.0 // 2021 Standard deduction
}

func calc_tax_burden(income float64) float64 {
	// Calculates your tax burden based off of tax brackets
	var total float64
	for _, key := range bracket_limits {
		if income >= key {
			// We want to calculate the tax for the portion of the income in
			// this Bracket
			total += (income - key) * tax_brackets[key]
			income = income - (income - key)
		}
	}
	return total
}
