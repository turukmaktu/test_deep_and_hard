package main

import (
	"github.com/go-redis/redis"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

/**
записать в redis событие
*/
func sendToRedis(api eventApi) (bool, resultApi){

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER") + ":6379",
		Password: "",
		DB:       0,
	})

	client.Incr("event:"+api.Id+":"+api.Label)

	return true, resultApi{Status: true, Message: "Success add"}
}

func checkInput( id string, label string) (bool, resultApi){

	result := true
	message := ""

	if len(id) == 0{
		result = false
		message += "not set id in request"
	}

	if len(label) == 0{
		result = false
		message += "not set label in request"
	}

	if result{
		return result, resultApi{Status: true, Message: "Success"}
	}else{
		return result, resultApi{Status: true, Message: message}
	}
}

func updateMysqlEventsFromRedis() (map[int]string){

	keys := map[int]string{}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER") + ":6379",
		Password: "",
		DB:       0,
	})

	res := client.Keys("event:*")

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/default")
	if err != nil {
		log.Println(err)
	}else{
		defer db.Close()
		for num, key := range  res.Val() {

			var data = strings.Split(key,":")

			count, err := client.Get(key).Result()

			if err != nil {
				panic(err)
			}
			if count != "0" {
				keys[num] = key
				var cnt int
				db.QueryRow("select count(*) from default.events where page_id = ? and label = ?", data[1], data[2]).Scan(&cnt)
				if cnt > 0 {
					db.Exec("update default.events set counter = counter + ? where page_id = ? and label = ?", count, data[1], data[2])
				}else{
					db.Exec("insert into default.events (page_id, label, counter) values (?, ?, ?)", data[1], data[2], count)
				}

				client.Set(key,0,0)
			}
		}
	}

	return keys
}