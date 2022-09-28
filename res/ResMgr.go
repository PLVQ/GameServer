package res

import (
	"io/ioutil"

	"gameServer/log"
	"gameServer/res/respb"

	"google.golang.org/protobuf/proto"
)

var ActivityOpenConfigListData respb.ActivityOpenConfigList

func init() {
	LoadPbBin("ActivityOpenConfig", &ActivityOpenConfigListData)
}

func LoadPbBin(pFile string, pConfig proto.Message) {
	if pConfig == nil {
		log.Log.Fatal("Config Is Null!")
	}

	pf, err := ioutil.ReadFile("./resource/" + pFile + ".bin")
	if err != nil {
		log.Log.WithField("error", err.Error()).Fatal()
	}

	err = proto.Unmarshal(pf, pConfig)
	if err != nil {
		log.Log.WithField("error", err.Error()).Fatal()
	}

	log.Log.WithField("Config", pFile).Info("Load Succ!")
}
