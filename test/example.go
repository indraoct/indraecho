package main
import("indraecho/config"
)



func main(){
	config.ConnectDB()
	rows, err := config.Db.Query("select id, username, password from users")
	defer rows.Close()
	config.CloseDB()
	if err != nil {
		return "", err
	}
        return


}