package childcliaubnig

import (
	"fmt"
	"errors"
)

func checkCliInputStringParams(stringParams string, showParams string) error {
	if stringParams == "" {
		return errors.New(fmt.Sprintf("\nYou are not setting [ %s ] exit!", showParams))
	}
	return nil
}
