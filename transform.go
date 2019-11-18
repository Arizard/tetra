package tetra

// TransformCSV runs multiple transforms on a csv, as defined in a Config.
func TransformCSV(cfg Config, csv string) string {
	newCsv := csv

	for _, transform := range cfg.Transforms {
		newCsv = operate(transform, newCsv)
	}

	return newCsv
}
