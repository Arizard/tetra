package tetra

import "log"

// TransformCSV runs multiple transforms on a csv, as defined in a Config.
func TransformCSV(cfg Config, csv string) string {
	newCsv := csv

	for _, transform := range cfg.Transforms {
		tempCsv, err := operate(transform, newCsv)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		newCsv = tempCsv
	}

	return newCsv
}
