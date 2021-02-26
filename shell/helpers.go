package shell

import "sort"

func CountCommands(history []string) map[string]float64 {
	duplicate := map[string]float64{}

	for _, item := range history {
		_, exist := duplicate[item]

		if exist {
			duplicate[item]++
		} else {
			duplicate[item] = 1
		}
	}
	return duplicate
}

func SortCommands(commands map[string]float64) (names []string, occurrences []float64) {
	keys := make([]string, 0, len(commands))
	values := make([]float64, 0, len(commands))

	for name := range commands {
		keys = append(keys, name)
	}

	sort.Slice(keys, func(i, j int) bool {
		return commands[keys[i]] > commands[keys[j]]
	})

	for _, name := range keys {
		values = append(values, commands[name])
	}

	return keys, values
}
