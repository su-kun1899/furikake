package furikake

import (
	"fmt"
	"github.com/su-kun1899/chazuke"
	"strings"
)

func ToCsv(header []string, json string) (string, error) {
	container, err := chazuke.FromJSON(json)
	if err != nil {
		return "", err
	}

	containers, err := container.Array()
	if err != nil {
		return "", err
	}

	values := []string{format(header)}
	for _, val := range containers {
		rows := make([]string, 0)

		for _, key := range header {
			v, err := val.Get(key).Value()
			if err != nil {
				return "", nil
			}
			rows = append(rows, v)
		}
		values = append(values, format(rows))
	}

	return strings.Join(values, "\n"), nil
}

func format(values []string) string {
	formatted := make([]string, len(values))
	for i, v := range values {
		// ダブルクォートをエスケープして囲う
		formatted[i] = fmt.Sprintf("\"%s\"", strings.ReplaceAll(v, "\"", "\"\""))
	}
	return strings.Join(formatted, ",")
}
