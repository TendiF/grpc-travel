package users

import (
	"context"
	"math/rand"

	"travel/proto"
	"travel/utils"
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
	// req.Header.Add("Authorization", "Bearer ")
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

	_, err = utils.DB.Exec(query, "register", params.CountryCode, params.Phone, params.FirstName, params.LastName, params.Gender, params.Email, params.BirthDate, params.Password, params.DocumentNumber, params.Address, verCode)

	if err != nil {
		response.Message = err.Error()
		return &response, nil
	}

	response.Message = "Success"

	return &response, nil
}

func (s *Server) Login(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {

	return &proto.UserResponse{}, nil
}
