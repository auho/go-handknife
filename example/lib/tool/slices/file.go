package slices

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadIntFromFileContent(c string) ([]int, error) {
	var ids []int

	idsString := strings.Split(c, "\n")
	for _, idStr := range idsString {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, fmt.Errorf("converting[%s]to int %w", idStr, err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}
