package common

import (
	"os"
	"strconv"
)

const KnativeUsedbyKey = "knative.dev/usedByResource"

func ShouldApplyFilter() bool {
	b, err := strconv.ParseBool(os.Getenv("INFORMER_FILTER"))
	if err != nil {
		return false
	}
	return b
}