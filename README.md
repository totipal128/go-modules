# Go Modules

## Type Package Modules golang
- [check]()
- [conv]()
- [databases]()
- [get_env]()
- [logs]()
- [utils]()


***
***
***


## check
to check the condition of the desired data

## conv
is to convert data, for example from data interface to map[string]interface{}

## databases
melakukan koneksi dengan mudah dari databases menggunakan env <br>

several database connections including:
- postgreSQL <br>

| ENV       | Type     | Default                    | Description       |
| :-------- | :------- | :------------------------- | :------           |
| `api_key` | `string` | **Required**. Your API key |                   |

- mongoDBV2

| ENV               | Type      | Default                    | Description       |
| :--------         | :-------  | :------------------------- | :------           |
| `DB_NAME`         | `string`  | admin                      |                |
| `CLUSTER_URL`     | `string`  | cluster0.admin.mongodb.net |                   |
| `USER`            | `string`  |                            |                   |
| `PASS`            | `string`  |                            |                   |
| `MIN_POOL_SIZE`   | `int`     | 10                         |                   |
| `MAX_POOL_SIZE`   | `int`     | 100                        |                   |


## get_env
function to call env

## logs
function to call to display data logs on the terminal
## utils
additional function data code requirements