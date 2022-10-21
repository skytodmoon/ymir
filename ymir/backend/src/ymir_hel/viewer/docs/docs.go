// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users/{userID}/repo/{repoID}/branch/{branchID}/assets": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query single or set of assets.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Offset, default is 0",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit, default is 1",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. class_ids=1,3,7",
                        "name": "class_ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. annotation_types=GT,PRED",
                        "name": "annotation_types",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. current_asset_id=xxxyyyzzz",
                        "name": "current_asset_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. cm_types=0,1,2,3 NotSet=0,TP=1,FP=2,FN=3,TN=4,Unknown=5,MTP=11,IGNORED=12",
                        "name": "cm_types",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ck pairs, e.g. cks=xxx,xxx:,xxx:yyy, e.g. camera_id:1",
                        "name": "cks",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "tag pairs, e.g. cks=xxx,xxx:,xxx:yyy, e.g. camera_id:1",
                        "name": "tags",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'result': constants.QueryAssetsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userID}/repo/{repoID}/branch/{branchID}/dataset_meta_count": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset info, lightweight api.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'result': constants.QueryDatasetStatsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userID}/repo/{repoID}/branch/{branchID}/dataset_stats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset Stats.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "e.g. class_ids=1,3,7",
                        "name": "class_ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. require_assets_hist",
                        "name": "require_assets_hist",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "e.g. require_annos_hist",
                        "name": "require_annos_hist",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'result': constants.QueryDatasetStatsResult",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userID}/repo/{repoID}/branch/{branchID}/model_info": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query model info.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Branch ID",
                        "name": "branchID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'result':  constants.MirdataModel",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userID}/repo/{repoID}/dataset_duplication": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query dataset dups.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Repo ID",
                        "name": "repoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "e.g. candidate_dataset_ids=xxx,yyy",
                        "name": "candidate_dataset_ids",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "dataset_ids to be corroded",
                        "name": "corrodee_dataset_ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "'code': 0, 'msg': 'Success', 'Success': true, 'result': 'duplication: 50, total_count: {xxx: 100, yyy: 200}'",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
