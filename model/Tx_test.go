package model

import (
	"os"
	"testing"

	"github.com/btcboost/copernicus/utils"
)

var testNewTx *Tx

func TestNewTx(t *testing.T) {
	testNewTx = NewTx()
	t.Log(testNewTx)

}

func TestTx_AddTxIn(t *testing.T) {
	var buf utils.Hash
	copy(buf[:], "adbasg7wy7yswdwiuyc78sayxchwuniuhy")
	testNewTx.Hash = buf

	myOutPut := NewOutPoint(&buf, 10)
	myString := "hwd7yduncue0qe01ie8dhuscb3etde21gdahsbchqbw1y278"
	mySigscript := make([]byte, len(myString))
	copy(mySigscript, myString)
	myTxIn := NewTxIn(myOutPut, mySigscript)

	testNewTx.AddTxIn(myTxIn)
}

func TestTx_AddTxOut(t *testing.T) {

	myString := "asdqwhihnciwiqd827w7e6123cdsnvh43yt892ufimjf27rufian2yr8sacmejfgu3489utwej"
	outScript := make([]byte, len(myString))
	copy(outScript, myString)
	txOut := NewTxOut(999, outScript[:])

	testNewTx.AddTxOut(txOut)
	t.Log(testNewTx.SerializeSize())
	t.Log(testNewTx)
}

func TestTx_Copy(t *testing.T) {

	copyTx := testNewTx.Copy()
	t.Log("copyTx : ", copyTx.SerializeSize())
	t.Log("copyTx : ", copyTx)
}

func TestTx_Serialize(t *testing.T) {

	file, err := os.OpenFile("tx.txt", os.O_RDWR|os.O_CREATE, 0666)
	checkErr(err)
	defer file.Close()

	err = testNewTx.Serialize(file)
	checkErr(err)
}

func TestTx_Deserialize(t *testing.T) {

	file, err := os.OpenFile("tx.txt", os.O_RDWR|os.O_CREATE, 0666)
	checkErr(err)
	defer file.Close()
	testNewTx.Deserialize(file)
}
