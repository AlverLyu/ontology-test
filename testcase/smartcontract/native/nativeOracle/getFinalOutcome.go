/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package nativeOracle

import (
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/core/genesis"
	"encoding/json"
	"encoding/hex"
)

func GetFinalOutcome(ctx *testframework.TestFrameworkContext, tx string) interface{} {

	contractAddress := genesis.OracleContractAddress
	txHashHex := tx
	txHash, _:= hex.DecodeString(txHashHex)
	key := append([]byte("FinalOutcome"), txHash...)
	value, err := ctx.Ont.Rpc.GetStorage(contractAddress, []byte(key))
	if err != nil {
		ctx.LogError("GetFinalOutcome GetStorageItem key:%s error:%s", key, err)
		return false
	}
	if len(value) == 0 {
		ctx.LogError("GetFinalOutcome FinalOutcome is not set!")
		return false
	}

	result := new(interface{})
	err = json.Unmarshal(value, &result)
	if err != nil {
		ctx.LogError("GetFinalOutcome Unmarshal result error:%s", err)
		return false
	}

	return *result
}