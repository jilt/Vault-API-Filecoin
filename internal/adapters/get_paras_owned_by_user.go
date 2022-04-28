package adapters

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/jilt/Vault-API-Filecoin/internal/models"
	"github.com/jilt/Vault-API-Filecoin/logger"
)

func GetParasOwnedByUser(user models.UserIDParameter) (*map[string]interface{}, error) {

	// create the query
	queryTemplate := `
	query MyQuery {
  nftBuys(where: {owner_id: "{{.User}}"}) {
    token_series_id
    token_id
  },
  nftTransfers(where: {new_owner_id: "{{.User}}"}) {
    token_series_id
    token_id
  }
}
`

	tmpl, err := template.New("queryTemplate").Parse(queryTemplate)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, user)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	query := b.String()

	jsonData := map[string]string{
		"query": query,
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.thegraph.com/subgraphs/name/jilt/parasubgraph", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 100}

	response, err := client.Do(request)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	var resp map[string]interface{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	return &resp, nil
}
