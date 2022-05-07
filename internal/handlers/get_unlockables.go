package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	"github.com/jilt/Vault-API-Filecoin/internal/models"
	"github.com/web3-storage/go-w3s-client"
)

type UnlockableHandler struct {
	Token string
	CID   string
}

func (h *UnlockableHandler) Handle(c *gin.Context) {
	var unlockableParameter models.UnlockableParameter

	if err := c.ShouldBindUri(&unlockableParameter); err != nil {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "provide a valid token ID",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	sess, err := h.ConnectToW3Storage()
	if err != nil {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "could not connect to w3 storage",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	h.Token = unlockableParameter.TokenId
	//h.CID = "bafybeigk7ttzgm7i3j7tsfg6ql7rv3ykaratezm63gwub6iw3vqonkm5yy"
	h.CID = h.getLatestW3Upload(sess)
	_, tokenMap, err := h.DownloadCidJson()
	if err != nil {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "failed to fetch data",
		}
		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	//check if tokenid exists in cid json struct
	if _, ok := tokenMap[h.Token]; ok {
		c.JSON(http.StatusOK, tokenMap[h.Token])
		return
	} else {
		e := models.BasicError{
			Code:    models.InvalidTokenID.String(),
			Message: "token id does not exist",
		}
		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}
}

func (h *UnlockableHandler) ConnectToW3Storage() (w3s.Client, error) {
	wc, err := w3s.NewClient(
		w3s.WithEndpoint(os.Getenv("W3_STROAGE_TOKEN")),
	)
	if err != nil {
		panic(err)
	}

	return wc, err
}

func (h *UnlockableHandler) getLatestW3Upload(c w3s.Client) string {
	//Get Total count of uploads
	var latestCid cid.Cid
	uploads, err := c.List(context.Background(), w3s.WithMaxResults(1))
	if err != nil {
		log.Println("Error: while list w3c storage ", err)
		panic(err)
	}

	for {
		u, err := uploads.Next()
		if err != nil {
			// finished successfully
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%s	%s	Size: %d	Deals: %d	Pins: %d\n", u.Created.Format("2006-01-02 15:04:05"), u.Cid, u.DagSize, len(u.Deals), len(u.Pins))
		latestCid = u.Cid
	}

	return latestCid.String()
}

func (h *UnlockableHandler) DownloadCidJson() (models.UnlockableCidJson, map[string]interface{}, error) {
	var (
		unlockableJson models.UnlockableCidJson
		URL            = fmt.Sprintf("https://%s.ipfs.dweb.link/root.json", h.CID)
		myClient       = &http.Client{Timeout: 60 * time.Second}
		tokenDataMap   = make(map[string]interface{})
	)

	r, err := myClient.Get(URL)
	if err != nil {
		return unlockableJson, nil, err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&unlockableJson)
	if err != nil {
		return unlockableJson, nil, err
	}

	log.Println(unlockableJson)

	for _, v := range unlockableJson {
		tokenDataMap[v.Name] = v
	}

	return unlockableJson, tokenDataMap, nil
}
