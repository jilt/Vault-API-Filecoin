# Vault-API-Filecoin
 The HTTP API to get unlockable content from the Varda Vault immutable DB
 
 HTTP Endpoints for checking ownership:
 
 `/owned/{user}`
 
 Gets all mintbase NFTs owned by the user: id, title, image and description
 
 `/owned/{user}+{storename}`
  
 Gets all mintbase owned NFT within the same smart contract
 
 `/owners/{tokenid}`
 
 Get all owners for a given NFT (each editions) on mintbase
 
  `/owned-paras/{user}`
 
 Gets all paras NFTs owned by the user: id, title, image and description
 
 `/owned-paras/{user}+{collection}`
  
 Gets all paras owned NFT within the same collection
 
 `/owners-paras/{tokenid}`
 
 Get all owners for a given NFT (each editions) on paras

`/unlockables/{userkey}+{tokenid}`

Gets unlockable content for the given NFT token id
