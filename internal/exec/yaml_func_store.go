package exec

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/cloudposse/atmos/pkg/schema"
	u "github.com/cloudposse/atmos/pkg/utils"
)

func processTagStore(atmosConfig schema.AtmosConfiguration, input string, currentStack string) any {
	log.Debug("Executing Atmos YAML function store", "input", input)

	str, err := getStringAfterTag(input, u.AtmosYamlFuncStore)

	if err != nil {
		u.LogErrorAndExit(atmosConfig, err)
	}

	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		u.LogErrorAndExit(atmosConfig, fmt.Errorf("invalid Atmos Store YAML function execution:: %s\nexactly two parameters are required: store_name, key", input))
	}

	storeName := strings.TrimSpace(parts[0])
	key := strings.TrimSpace(parts[1])

	store := atmosConfig.Stores[storeName]

	if store == nil {
		u.LogErrorAndExit(atmosConfig, fmt.Errorf("invalid Atmos Store YAML function execution:: %s\nstore '%s' not found", input, storeName))
	}

	value, err := store.Get(key)
	if err != nil {
		u.LogErrorAndExit(atmosConfig, fmt.Errorf("invalid Atmos Store YAML function execution:: %s\nkey '%s' not found in store '%s'", input, key, storeName))
	}

	return value
}
