//
// Copyright (c) 2018-2019 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package unified

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/che-go-jsonrpc"
	"github.com/eclipse/che-plugin-broker/brokers/che-plugin-broker"
	"github.com/eclipse/che-plugin-broker/brokers/theia"
	"github.com/eclipse/che-plugin-broker/brokers/vscode"
	"github.com/eclipse/che-plugin-broker/common"
	"github.com/eclipse/che-plugin-broker/model"
	"github.com/eclipse/che-plugin-broker/storage"
	"github.com/eclipse/che-plugin-broker/utils"
	"net/http"
)

const ChePluginType = "Che plugin"
const EditorPluginType = "Che Editor"
const TheiaPluginType = "Theia plugin"
const VscodePluginType = "VS Code extension"

// Broker is used to process Che plugins
type Broker struct {
	common.Broker
	Storage *storage.Storage

	theiaBroker  *theia.Broker
	vscodeBroker *vscode.Broker
	cheBroker    *broker.ChePluginBroker
}

// NewBroker creates Che broker instance
func NewBroker() *Broker {
	commonBroker := common.NewBroker()
	ioUtils := utils.New()
	storageObj := storage.New()
	httpClient := &http.Client{}
	rand := common.NewRand()

	cheBroker := broker.NewBrokerWithParams(commonBroker, ioUtils, storageObj)
	theiaBroker := theia.NewBrokerWithParams(commonBroker, ioUtils, storageObj, rand)
	vscodeBroker := vscode.NewBrokerWithParams(commonBroker, ioUtils, storageObj, rand, httpClient)
	return &Broker{
		Storage: storageObj,
		Broker:  commonBroker,

		theiaBroker:  theiaBroker,
		vscodeBroker: vscodeBroker,
		cheBroker:    cheBroker,
	}
}

// Start executes plugins metas processing and sends data to Che master
func (b *Broker) Start(metas []model.PluginMeta) {
	b.PubStarted()
	b.PrintInfo("Unified Che Plugin Broker")

	b.PrintPlan(metas)

	cheMetas, theiaMetas, vscodeMetas, err := sortMetas(metas)
	if err != nil {
		b.PubFailed(err.Error())
		b.PrintFatal(err.Error())
	}

	b.PrintInfo("Starting Che common plugins processing")
	for _, meta := range cheMetas {
		err := b.cheBroker.ProcessPlugin(meta)
		if err != nil {
			b.PubFailed(err.Error())
			b.PrintFatal(err.Error())
		}
	}
	b.PrintInfo("Starting Theia plugins processing")
	for _, meta := range theiaMetas {
		err := b.theiaBroker.ProcessPlugin(meta)
		if err != nil {
			b.PubFailed(err.Error())
			b.PrintFatal(err.Error())
		}
	}
	b.PrintInfo("Starting VS Code plugins processing")
	for _, meta := range vscodeMetas {
		err := b.vscodeBroker.ProcessPlugin(meta)
		if err != nil {
			b.PubFailed(err.Error())
			b.PrintFatal(err.Error())
		}
	}

	plugins, err := b.Storage.Plugins()
	if err != nil {
		b.PubFailed(err.Error())
		b.PrintFatal(err.Error())
	}
	pluginsBytes, err := json.Marshal(plugins)
	if err != nil {
		b.PubFailed(err.Error())
		b.PrintFatal(err.Error())
	}

	b.PrintInfo("All plugins have been successfully processed")
	result := string(pluginsBytes)
	b.PrintDebug(result)
	b.PubDone(result)
	b.CloseConsumers()
}

// PushEvents sets given tunnel as consumer of broker events.
func (b *Broker) PushEvents(tun *jsonrpc.Tunnel) {
	b.Broker.PushEvents(tun, model.BrokerStatusEventType, model.BrokerResultEventType, model.BrokerLogEventType)
}

func sortMetas(metas []model.PluginMeta) (che []model.PluginMeta, theia []model.PluginMeta, vscode []model.PluginMeta, err error) {
	vscodeMetas := make([]model.PluginMeta, 0)
	theiaMetas := make([]model.PluginMeta, 0)
	cheBrokerMetas := make([]model.PluginMeta, 0)
	for _, meta := range metas {
		switch meta.Type {
		case ChePluginType:
			fallthrough
		case EditorPluginType:
			cheBrokerMetas = append(cheBrokerMetas, meta)
		case VscodePluginType:
			vscodeMetas = append(vscodeMetas, meta)
		case TheiaPluginType:
			theiaMetas = append(theiaMetas, meta)
		default:
			return nil, nil, nil, fmt.Errorf("Type field is missing in metainformation of plugin '%s:%s'", meta.ID, meta.Version)
		}
	}

	return cheBrokerMetas, theiaMetas, vscodeMetas, nil
}
