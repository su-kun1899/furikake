package furikake

import (
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

	values := []string{strings.Join(header, ",")}
	for _, val := range containers {
		row := make([]string, 0)

		for _, key := range header {
			v, err := val.Get(key).Value()
			if err != nil {
				return "", nil
			}
			row = append(row, v)
		}
		values = append(values, strings.Join(row, ","))
	}

	return strings.Join(values, "\n"), nil
}
