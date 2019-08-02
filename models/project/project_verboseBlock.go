package models

import "gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models"

type VerboseBlock struct {
    models.Proto
    Hash string `json:"hash"`
    Info map[string]interface{} `json:"info"`
}
