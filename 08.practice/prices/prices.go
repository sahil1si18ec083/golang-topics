package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate        float64                 `json:"input_prices"`
	InputPrices    []float64               `json:"tax_rate"`
	TaxInputPrices map[string]string       `json:"tax_input_prices"`
	IOFileManager  filemanager.Filemanager `json:"-"`
}

func NewTaxIncludedPriceJob(taxRate float64, f filemanager.Filemanager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate:       taxRate,
		IOFileManager: f,
	}

}

func (Job *TaxIncludedPriceJob) Process(done_chan chan bool, val_chan chan error) {
	Job.LoadData()
	result := make(map[string]string)
	// fmt.Println(Job.InputPrices)

	for _, price := range Job.InputPrices {
		r := price * (1 + Job.TaxRate)
		s := fmt.Sprintf("%.2f", r)
		result[fmt.Sprintf("%.2f", price)] = s
	}
	// path := fmt.Sprintf("%.2f_data.json", Job.TaxRate*100)
	// fmt.Println(path)

	Job.TaxInputPrices = result
	Job.IOFileManager.WriteJSON(Job)
	done_chan <- true
	val_chan <- nil

}

func (Job *TaxIncludedPriceJob) LoadData() {
	lines, err := Job.IOFileManager.ReadFile()
	if err != nil {
		panic(err)
	}
	gg, _ := conversion.StringsToFloat(lines)

	// fmt.Println(gg)
	Job.InputPrices = gg

}
