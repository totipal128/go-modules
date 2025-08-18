package check

import "fmt"

// desc =>[=======       ]20%
func LoadProgressBarTerminal1(countProgress, total int, desc string) {
	percentage := float64(countProgress) / float64(total) * 100
	barWidth := 50
	progress := int(percentage / 100 * float64(barWidth))

	bar := ""
	for i := 0; i < progress; i++ {
		bar += "\u001B[33m=\u001B[0m"
	}
	for i := progress; i < barWidth; i++ {
		bar += " "
	}

	// \r untuk menimpa baris sebelumnya
	fmt.Printf("\r %s |[%s]\u001B[31m %.2f%% \u001B[0m", desc, bar, percentage)
}
