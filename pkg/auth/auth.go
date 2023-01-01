package auth

import (
	"echo-auth/pkg/config"
	"echo-auth/pkg/models"
	"echo-auth/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	config.LoadEnv()
}

var (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
	jwtSecretKey           = os.Getenv("JWT_SECRET")
)

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
	ExpiresAt    string `json:"expires_at"`
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
func GenerateTokensAndSetCookies(user *models.User, c echo.Context) Token {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		log.Println("Token Error", time.Now(), err)
		return Token{}
	}
	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)
	// We generate here a new refresh token and saving it to the cookie.
	refreshToken, _, refErr := generateRefreshToken(user)
	if refErr != nil {
		log.Println("Refresh Token Error", time.Now(), err)
		return Token{}
	}
	setTokenCookie(refreshTokenCookieName, refreshToken, exp, c)
	//return accessToken, exp, err
	return Token{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username, ExpiresAt: exp.String()}
}

func generateAccessToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time.
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("", time.Now(), err)
		return "", time.Now(), err
	}
	return tokenString, expirationTime, nil
}

func generateRefreshToken(user *models.User) (string, time.Time, error) {
	// Declare the expiration time of the token - 24 hours.
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

// Here we are creating a new cookie, which will store the valid JWT token.
func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// Purpose of this cookie is to store the user's name.
func setUserCookie(user *models.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Username
	cookie.Expires = expiration
	cookie.Path = "/"

	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	//return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("login"))
	rs := utils.ResponseStruct{StatusCode: http.StatusUnauthorized, MessageEn: "Unauthorized"}
	return rs.WriteToResponse(c, nil, "en")
}
