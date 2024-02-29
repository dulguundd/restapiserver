package driving

import (
	"errors"
	"net"
	"net/http"
	"restAPIServer/app/dto"
	"strings"
)

func newUserRequestValidation(req dto.NewUserRequest) (bool, *dto.NewUserValidationError) {
	var incorrectRequestBool bool = false
	var newUserValidationError dto.NewUserValidationError
	if len([]rune(req.Username)) < 6 {
		incorrectRequestBool = true
		newUserValidationError.Username = "username too short"
	}
	if strings.Index(req.Email, "@") == -1 {
		incorrectRequestBool = true
		newUserValidationError.Email = "mail is incorrect"
	}
	if len([]rune(req.Password)) < 6 {
		incorrectRequestBool = true
		newUserValidationError.Password = "password too short"
	}
	if incorrectRequestBool == true {
		return incorrectRequestBool, &newUserValidationError
	} else {
		return incorrectRequestBool, nil
	}
}

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}
	return "", errors.New("IP not found")
}
