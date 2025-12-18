package prices

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TaxIncludedPriceJob struct {
	TaxRate        float64
	InputPrices    []float64
	TaxInputPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}

}

func (Job *TaxIncludedPriceJob) Process() {
	Job.LoadData()
	result := make(map[string]float64)
	// fmt.Println(Job.InputPrices)

	for _, price := range Job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + Job.TaxRate)
	}
	fmt.Println(result)

}

func (Job *TaxIncludedPriceJob) LoadData() {
	file, err := os.ReadFile("prices.txt")

	if err != nil {
		fmt.Println("error")
		return
	}
	str := string(file)

	lines := strings.Split(str, "\n")

	gg := []float64{}

	for _, val := range lines {
		val = strings.TrimSpace(val)
		temp, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return
		}
		gg = append(gg, temp)
	}
	// fmt.Println(gg)
	Job.InputPrices = gg

}
