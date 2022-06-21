# Vault-API-Filecoin
 The HTTP API to get unlockable content from the Varda Vault filecoin immutable DB
 
 [Detailed docs for each coding language](https://jilt.github.io/Vault-API-Filecoin/)
 
 HTTP Endpoints for checking ownership:
 
 `https://filecoin-dlubh5ly6a-uc.a.run.app/owned/{user}`
 
 Gets all mintbase NFTs owned by the user: id, title, image and description
 
 `https://filecoin-dlubh5ly6a-uc.a.run.app/owned/{user}/{store}`
  
 Gets all mintbase owned NFT within the same smart contract
 
 `https://filecoin-dlubh5ly6a-uc.a.run.app/owners/{tokenid}`
 
 Get all owners for a given NFT (each editions) on mintbase
 
  `https://filecoin-dlubh5ly6a-uc.a.run.app/owned-paras/{user}`
 
 Gets all paras NFTs owned by the user: id, title, image and description
 
 `https://filecoin-dlubh5ly6a-uc.a.run.app/owned-paras/{user}/{collection}`
  
 Gets all paras owned NFT within the same collection (not yet implemented)
 
 `https://filecoin-dlubh5ly6a-uc.a.run.app/owners-paras/{tokenid}`
 
 Get all owners for a given NFT (each editions) on paras

`https://filecoin-dlubh5ly6a-uc.a.run.app/unlockables/{tokenid}`

Gets filecoin unlockable content for the given NFT token id
