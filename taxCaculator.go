package main

func taxCaculator(name string, yearsEarning int) (float64){
	var taxRate float64 = 15.0 / 100.0
	var higherTaxRate float64 = 20.0 / 100.0

	var tax float64

	if yearsEarning > 30_000 { 
		 tax  = higherTaxRate * float64(yearsEarning)
	}
	if yearsEarning < 30_000 {
		 tax  = taxRate * float64(yearsEarning)	
	}
	return tax
}
