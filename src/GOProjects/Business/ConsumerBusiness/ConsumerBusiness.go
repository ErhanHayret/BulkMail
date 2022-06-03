package ConsumerBusiness

import(
	"bulkmail/packages/DataAccess/MongoDb"
	"bulkmail/packages/Data/Models"
)

var collection, dbResponse = MongoDb.GetClient("MailDb", "Mail")

func Insert(mail Models.MailModel) Models.ResultModel{
	if dbResponse.Status == false {
		return dbResponse
	}
	result := MongoDb.InsretMail(collection, mail)
	return result
}