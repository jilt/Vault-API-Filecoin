package adapters

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jilt/Vault-API-Filecoin/internal/models"
	"github.com/jilt/Vault-API-Filecoin/logger"
)

func GetOwners(tokenID models.OwnerParameter) (*map[string]interface{}, error) {
	queryTemplate := `
	query MyQuery {
		thing(where: {id: {_eq: "{{.TokenId}}"}}) {
		  metadata {
			thing {
			  tokens {
				ownerId
			  }
			}
		  }
		}
	  }
	  
`

	tmpl, err := template.New("queryTemplate").Parse(queryTemplate)
	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, tokenID)
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

	request, err := http.NewRequest(http.MethodPost, "https://mintbase-mainnet.hasura.app/v1/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Log.Info(err)
		return nil, models.ErrFailedFetchData
	}

	client := &http.Client{Timeout: time.Second * 100}

	response, err := client.Do(request)
	if err != nil {
		logger.Log.Info(err)
		return nil, models.ErrFailedFetchData
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Log.Info(err)
		return nil, models.ErrFailedFetchData
	}

	var resp map[string]interface{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		logger.Log.Info(err)
		return nil, models.ErrFailedFetchData
	}

	return &resp, nil
}
