package gen

const TempConfig = `
project:
  host: "0.0.0.0"
  port: 8080
  port_pub: 8080
  port_pri: 8081
mysql:
  username: "TempDBUser"
  password: "TempDBPwd"
  db_host: "TempDBAddr"
  db_port: TempDBPort
  db_name: "TempDBName"
  charset: "utf8"
  max_open_conn: 50
  max_idle_conn: 10
  conn_max_lifetime: 500

`
