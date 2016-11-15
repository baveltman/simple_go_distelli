package simple_go_distelli

type SimpleStruct struct {
	FirstName     	string		`json:"firstName" bson:"firstName"`
	LastName	string		`json:"lastName" bson:"lastName"`
	Gender		string		`json:"gender" bson:"gender"`
	Email 		string		`json:"email" bson:"email"`
	Password	string		`json:"password" bson:"password"`
	Username	string		`json:"username" bson:"username"`
	DisplayUsername	string		`json:"displayUsername" bson:"displayUsername"`

}
