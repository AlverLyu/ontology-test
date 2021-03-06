package runtime

import (
	"time"

	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System;
using System.ComponentModel;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static void Main(string msg)
    {
        Runtime.Log(msg);
    }
}
*/

func TestRuntimLog(ctx *testframework.TestFrameworkContext) bool {
	code := "51c56b6c766b00527ac4616c766b00c361681253797374656d2e52756e74696d652e4c6f6761616c7566"
	codeAddr, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()

	if err != nil {
		ctx.LogError("TestRuntimLog - GetDefaultAccount error: %s", err)
		return false
	}

	_, err = ctx.Ont.Rpc.DeploySmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		code,
		"TestRuntimLog",
		"",
		"",
		"",
		"")

	//等待出块
	_, err = ctx.Ont.Rpc.WaitForGenerateBlock(30*time.Second, 2)

	if err != nil {
		ctx.LogError("TestRuntimLog WaitForGenerateBlock error:%s", err)
		return false
	}

	txHash, err := ctx.Ont.Rpc.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddr,
		[]interface{}{"ontology"})

	if err != nil {
		ctx.LogError("TestRuntimLog InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogInfo("TestRuntimLog invoke Tx:%s\n", txHash.ToHexString())
	events, err := ctx.Ont.Rpc.GetSmartContractEvent(txHash)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}

	ctx.LogInfo("======events log===== %+v", events)

	return true

	//transfer := events[0].States
	//ctx.LogInfo("%+v", transfer)

	//notify , ok := res.(map[string]interface{})
	//if !ok {
	//	ctx.LogError("TestLog res asset to map[string]interface{} error:%s", err)
	//	return false
	//}
	//
	//log := notify["Message"]
	//if input !=log {
	//	ctx.LogError("TestLog log error %s != %s", input, log)
	//	return false
	//}
	return true
}
