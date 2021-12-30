package users

import (
	"context"
	"math/rand"
	"os"
	"time"

	"travel/proto"
	types "travel/types"
	"travel/utils"

	"github.com/dgrijalva/jwt-go"
)

type Server struct {
	proto.UnimplementedUsersServiceServer
}

func (s *Server) Register(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	// generate verification code
	numbers := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	max := 8
	min := 0
	verCode := numbers[rand.Intn(max-min+1)+min] + numbers[rand.Intn(max-min+1)+min] + numbers[rand.Intn(max-min+1)+min] + numbers[rand.Intn(max-min+1)+min]

	// verify empty data
	if params.CountryCode == "" || params.Email == "" || params.Phone == "" || params.FirstName == "" || params.LastName == "" {
		response.Message = "fill required data"
		return &response, nil
	}

	// CEQUENS_API_KEY :=  os.Getenv("CEQUENS_API_KEY")
	// send sms
	// client := &http.Client{}
	// postBody, _ := json.Marshal(map[string]interface{}{
	// 	"senderName":      "Yalla Pay",
	// 	"messageType":     "text",
	// 	"messageText":     "code: " + verCode,
	// 	"recipients":      params.CountryCode + params.Phone,
	// 	"acknowledgement": 0,
	// 	"flashing":        0,
	// })

	// responseBody := bytes.NewBuffer(postBody)
	// req, err := http.NewRequest("POST", "https://apis.cequens.com/sms/v1/messages", responseBody)
	// req.Header.Add("Authorization", "Bearer " + CEQUENS_API_KEY)
	// req.Header.Add("Content-Type", "application/json")
	// resp, err := client.Do(req)

	// if resp.StatusCode != 200 {
	// 	response.Message = resp.Status
	// 	return &response, nil
	// }

	// defer resp.Body.Close()

	//validate by phone number
	query := `SELECT phone, country_code FROM public.users WHERE country_code='` + params.CountryCode + `' AND phone='` + params.Phone + `'`

	rows, err := utils.DB.Query(query)

	for rows.Next() {
		var phone string
		var country_code string
		err = rows.Scan(&phone, &country_code)
		if phone != "" {
			response.Message = "already register:" + country_code + phone
			return &response, nil
		}
	}

	query = `INSERT INTO public.users (status, country_code, phone, first_name, last_name, gender, email, birth_date, password, document_number, address, verification_code)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err = utils.DB.Exec(query, "register", params.CountryCode, params.Phone, params.FirstName, params.LastName, params.Gender, params.Email, params.BirthDate, utils.HashAndSalt([]byte(params.Password)), params.DocumentNumber, params.Address, verCode)

	if err != nil {
		response.Message = err.Error()
		return &response, err
	}

	response.Message = "Success"

	return &response, nil
}

func (s *Server) Login(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	query := `SELECT id, password, email FROM users 
	WHERE (country_code=$1 AND phone=$2) 
	OR (email=$3)`

	rows, err := utils.DB.Query(query, params.CountryCode, params.Phone, params.Email)

	if err != nil {
		response.Message = err.Error()
		return &response, err
	}

	for rows.Next() {
		var id string
		var password string
		var email string
		err = rows.Scan(&id, &password, &email)
		if utils.ComparePasswords(password, []byte(params.Password)) {
			claims := &types.Claims{
				Uid: id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))

			if err != nil {
				response.Message = err.Error()
				return &response, err
			}

			response.Message = tokenString
		} else {
			response.Message = "wrong email/phone or password"
		}

		return &response, nil
	}

	response.Message = "Account not found"

	return &response, nil
}

func (s *Server) VerifyCode(ctx context.Context, params *proto.VerifyReqeust) (*proto.UserResponse, error) {
	var response proto.UserResponse

	var claims types.Claims
	token := params.Token

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})

	if err != nil || !tkn.Valid || claims == (types.Claims{}) {
		response.Message = err.Error()
		return &response, err
	}

	query := `SELECT verification_code, status FROM users WHERE id=$1`
	rows, err := utils.DB.Query(query, claims.Uid)

	if err != nil {
		response.Message = err.Error()
		return &response, err
	}

	for rows.Next() {
		var verification_code string
		var status string
		err = rows.Scan(&verification_code, &status)
		if err != nil {
			response.Message = err.Error()
			return &response, err
		}

		if status == "active" {
			response.Message = "already active"
			return &response, nil
		}

		if verification_code == params.VerificationCode {
			query := `UPDATE users SET verification_code = '', status='active' WHERE id=$1`
			_, err := utils.DB.Exec(query, claims.Uid)
			if err != nil {
				response.Message = err.Error()
				return &response, err
			}

			response.Message = "Success Verify"
			return &response, nil
		}
	}

	response.Message = "Fail verify"

	return &response, nil
}
