package graphql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"hcc/harp/config"
	"hcc/harp/logger"
	"hcc/harp/mysql"
	"hcc/harp/types"
	"hcc/harp/uuidgen"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CheckServerUUID : Check if server UUID is exist in violin server module
func CheckServerUUID(serverUUID string) error {
	query := fmt.Sprintf("%s", "query Select_Server {\n  server(uuid: \""+serverUUID+"\") {\n    uuid\n    subnet_id\n    os\n    server_name\n    server_disc\n    cpu\n    memory\n    disk_size\n    status\n    user_uuid\n    created_time\n }\n}\n")

	resp, err := http.PostForm("http://localhost:"+config.ViolinHTTPPort+
		"/graphql", url.Values{"query": {query},
		"variables":     {"{}"},
		"operationName": {"Select_Server"}})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	} else {
		return err
	}

	if strings.Contains(string(respBody), "null") {
		return errors.New("server UUID is not exist")
	}

	return nil
}

var mutationTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		////////////////////////////// volume ///////////////////////////////
		/* Create new volume
		http://localhost:8001/graphql?query=mutation+_{create_volume(size:1024000,type:"ext4",server_uuid:"[server_uuid]"){size,type,server_uuid}}
		*/
		"create_volume": &graphql.Field{
			Type:        volumeType,
			Description: "Create new volume",
			Args: graphql.FieldConfigArgument{
				"size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: create_volume")

				uuid, err := uuidgen.Uuidgen()
				if err != nil {
					logger.Logger.Println("Failed to generate uuid!")
					return nil, nil
				}

				volume := types.Volume{
					UUID:       uuid,
					Size:       params.Args["size"].(int),
					Type:       params.Args["type"].(string),
					ServerUUID: params.Args["server_uuid"].(string),
				}

				err = CheckServerUUID(volume.ServerUUID)
				if err != nil {
					logger.Logger.Println(err)
					return nil, nil
				}

				sql := "insert into volume(uuid, size, type, server_uuid) values (?, ?, ?, ?)"
				stmt, err := mysql.Db.Prepare(sql)
				if err != nil {
					logger.Logger.Println(err.Error())
					return nil, nil
				}
				defer stmt.Close()
				result, err2 := stmt.Exec(volume.UUID, volume.Size, volume.Type, volume.ServerUUID)
				if err2 != nil {
					logger.Logger.Println(err2)
					return nil, nil
				}
				logger.Logger.Println(result.LastInsertId())

				return volume, nil
			},
		},

		/* Update volume by uuid
		   http://localhost:8001/graphql?query=mutation+_{update_volume(uuid:"[volume_uuid]",size:10240,type:"ext4",server_uuid:"[server_uuid]"){uuid,size,type,server_uuid}}
		*/
		"update_volume": &graphql.Field{
			Type:        volumeType,
			Description: "Update volume by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"server_uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: update_volume")

				requestedUUID, _ := params.Args["uuid"].(string)
				size, sizeOk := params.Args["size"].(int)
				_type, _typeOk := params.Args["type"].(string)
				serverUUID, serverUUIDOk := params.Args["server_uuid"].(string)

				volume := new(types.Volume)

				if sizeOk && _typeOk && serverUUIDOk {
					volume.UUID = requestedUUID
					volume.Size = size
					volume.Type = _type
					volume.ServerUUID = serverUUID

					sql := "update volume set size = ?, type = ?, server_uuid = ? where uuid = ?"
					stmt, err := mysql.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(volume.Size, volume.Type, volume.ServerUUID, volume.UUID)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.LastInsertId())

					return volume, nil
				}
				return nil, nil
			},
		},

		/* Delete volume by id
		   http://localhost:8001/graphql?query=mutation+_{delete_volume(id:"test1"){id}}
		*/
		"delete_volume": &graphql.Field{
			Type:        volumeType,
			Description: "Delete volume by uuid",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				logger.Logger.Println("Resolving: delete_volume")

				requestedUUID, ok := params.Args["uuid"].(string)
				if ok {
					sql := "delete from volume where uuid = ?"
					stmt, err := mysql.Db.Prepare(sql)
					if err != nil {
						logger.Logger.Println(err.Error())
						return nil, nil
					}
					defer stmt.Close()
					result, err2 := stmt.Exec(requestedUUID)
					if err2 != nil {
						logger.Logger.Println(err2)
						return nil, nil
					}
					logger.Logger.Println(result.RowsAffected())

					return requestedUUID, nil
				}
				return nil, nil
			},
		},
	},
})
