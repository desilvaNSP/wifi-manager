package dao

type NameValue struct{
	Name string  `db:"name"json:"name"`
	Value int  `db:"value"json:"value"`
}

type User struct {
	Username string `db:"username"json:"username"`
	Password string `db:"password"json:"password"`
	Email string `db:"email"json:"email"`
}

type RadiusUser struct{
	Id      int64  `db:"id"json:"id"`
	Username string `db:"username"json:"username"`
	Password string `json:"password"`
	Acctstarttime string `db:"acctstarttime"json:"acctstarttime"`
	Acctlastupdatedtime string `db:"acctlastupdatedtime"json:"acctlastupdatedtime"`
	Acctactivationtime string `db:"acctactivationtime"json:"acctactivationtime"`
	Acctstoptime string `db:"acctstoptime"json:"acctstoptime"`
	Location string  `db:"location"json:"location"`
	Visits int64     `db:"visits"json:"visits"`
}

type Role struct{
   BitMask int `json:"bitMask"`
   Title string `json:"title"`
}

type AuthUser struct {
	Username string `json:"username"`
	Role Role   `json:"role"`
}

type Constrains struct {
	From string `json:"from"`
	To string   `json:"to"`
	LocationId string    `json:"locationid"`
}

type Location struct {
	LocationId string    `json:"locationid"`
	LocationName string  `json:"locationname"`
	NasIP string         `json:"nasip"`
	IPFrom string        `json:"ipfrom"`
	IPTo string          `json:"ipto"`

}

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type DBConfigs struct{
	Username string
	Password string
	DBName string
	Host string
	Port int
}

type ServerConfigs struct{
	Sample string
	HttpPort int
	HttpsPort int
	ReadTimeOut int
	WriteTimeOut int
}