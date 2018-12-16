package actions

import (
	"encoding/hex"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gobuffalo/buffalo"
)

func PingUpdateHandler(c buffalo.Context) error {
	rawFileID := c.Param("fileId")
	fileID, _ := strconv.Atoi(rawFileID)
	pkRaw := c.Param("pk")

	bkp, _ := hex.DecodeString(pkRaw)
	pk, _ := crypto.ToECDSA(bkp)

	PingFile(fileID, pk)
	return nil
}
