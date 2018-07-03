package aubnig

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type CfgAub struct {
	IsDebug    bool
	ConfAubNig ConfAubNig
}

// use as var jsonCfg = new(aubnig.CfgAub)
// jsonCfg.InitJsonCfg(jsonPath)
func (c *CfgAub) InitJsonCfg(jsonPath string) error {
	_, err := os.Stat(jsonPath)
	if err != nil {
		if ! os.IsExist(err) {
			return err
		} else {
			return err
		}
	}
	fileJson, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(fileJson)
	if err != nil {
		return err
	}
	var configJson ConfAubNig
	err = json.Unmarshal(bytes, &configJson)
	if err != nil {
		return err
	}
	c.ConfAubNig = configJson
	if configJson.RunMode == "dev" {
		c.IsDebug = true
	} else {
		c.IsDebug = false
	}
	return nil
}

func (c CfgAub) Debug() bool {
	return c.IsDebug
}

func (c CfgAub) ReadConfig() ConfAubNig {
	return c.ConfAubNig
}
