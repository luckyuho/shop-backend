package logInfo

import (
	"basic/db"
	"context"

	"github.com/sirupsen/logrus"
)

func InsertLog(
	header string,
	err error,
) {
	// 插入文档
	_, errMsg := db.GetMongo().InsertOne(context.TODO(), NewLogInfo(header, err))
	if errMsg != nil {
		logrus.Error("log insert error: ", err)
	}
}

// func UpdateLog() {
// 	// 更新文档
// 	filter := bson.D{{"name", "John"}}
// 	update := bson.D{{"$set", bson.D{{"age", 91}}}}

// 	updateResult, err := db.GetMongo().UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Updated", updateResult.ModifiedCount, "document(s)")
// }
