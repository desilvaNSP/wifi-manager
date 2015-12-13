package authenticator

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
	"wislabs.wifi.manager/common"
	"wislabs.wifi.manager/redis"
	"strconv"
	"wislabs.wifi.manager/utils"
	"database/sql"
	log "github.com/Sirupsen/logrus"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationEngine() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}
	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(user *common.SystemUser) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	i, _ := strconv.Atoi(os.Getenv(common.JWT_EXPIRATION_DELTA))
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(i)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = user.Password
	token.Claims["userid"] = getUserId(user)
	sample := getUserScopes(user)
	token.Claims["scopes"] = sample
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}

func getUserId(user *common.SystemUser) int64{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	var userId sql.NullInt64
	smtOut, err := dbMap.Db.Prepare("SELECT userid FROM users WHERE username=? ANd tenantid=?")
	defer smtOut.Close()
	err = smtOut.QueryRow(user.Username, user.TenantId).Scan(&userId)
	if err != nil {
		log.Debug("User authentication failed " + user.Username)
		return -1
	}else {
		user.UserId = userId.Int64
		return userId.Int64
	}
	return -1
}

func getUserScopes(user *common.SystemUser) map[string][]string{
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()
	rows, err := dbMap.Db.Query("select name,action from permissions where permissionid in (select userpermissions.permissionid from userpermissions where userpermissions.userid = ?) order by name", user.UserId)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var(
		name string
		action string
	)

	scopes := make(map[string][]string)
	for rows.Next() {
		err := rows.Scan(&name, &action)
		if err != nil {
			log.Fatal(err)
		}
		scopes[name] = append(scopes[name],action)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scopes
}

func (backend *JWTAuthenticationBackend) Authenticate(user *common.SystemUser) bool {
	dbMap := utils.GetDBConnection("dashboard");
	defer dbMap.Db.Close()

	var hashedPassword sql.NullString
	smtOut, err := dbMap.Db.Prepare("SELECT password FROM users where username=? ANd tenantid=? and status='active'")
	defer smtOut.Close()
	err = smtOut.QueryRow(user.Username, user.TenantId).Scan(&hashedPassword)
	if err == nil && hashedPassword.Valid {
		if (len(hashedPassword.String) > 0) {
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword.String), []byte(user.Password))
			if err == nil {
				log.Debug("User authenticated successfully " + user.Username)
				return true
			}
		}
	}else {
		log.Debug("User authentication failed " + user.Username)
		return false
	}
	return false
}

func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	redisConn := redis.Connect()
	return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims["exp"]))
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	redisConn := redis.Connect()
	redisToken, _ := redisConn.GetValue(token)

	if redisToken == nil {
		return false
	}
	return true
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(os.Getenv(common.JWT_PRIVATE_KEY_PATH))
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	defer privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}
	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(os.Getenv(common.JWT_PUBLIC_KEY_PATH))
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	defer publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}