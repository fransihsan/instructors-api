package validators

import (
	"errors"
	"regexp"
)

func ValidateCreateInstructor(name, email, password string) error {
	namePattern, _ := regexp.Compile(`^([A-Za-z]+ ?[A-Za-z]*){3,29}[A-Z-a-z]$`)
	if len(name) > 30 || !namePattern.MatchString(name) {
		return errors.New("input nama tidak sesuai (alfabet; tanpa simbol; boleh ada spasi di antara kata; total karakter: minimal 4, maksimal 30;)\ncontoh: \"Yusuf Nur Wahid\"")
	}

	emailPattern, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailPattern.MatchString(email) {
		return errors.New("input email tidak sesuai (username: lowercase alfanumerik, simbol yang diperbolehkan hanya '_' dan '-'; pemisah: simbol '@'; domain name: lowercase alfanumerik, simbol yang diperbolehkan hanya '-'; pemisah: simbol '.'; domain: lowercase alfabet, minimal 2 karakter, maksimal 4 karakter;)\ncontoh: \"yusuf@mail.com\"")
	}

	passwordPattern, _ := regexp.Compile(`^[a-zA-Z0-9!@#$&()\\\-\x60.+,/\"]{5,8}$`)
	if !passwordPattern.MatchString(password) {
		return errors.New("input password tidak sesuai (alfanumerik; boleh menggunakan simbol; total karakter: minimal 5, maksimal 8;)")
	}
	return nil
}

func ValidateUpdateInstructor(name, email, password string) error {
	namePattern, _ := regexp.Compile(`^([A-Za-z]+ ?[A-Za-z]*){3,29}[A-Z-a-z]$`)
	if len(name) != 0 && (len(name) > 30 || !namePattern.MatchString(name)) {
		return errors.New("input nama tidak sesuai (alfabet; tanpa simbol; boleh ada spasi di antara kata; total karakter: minimal 4, maksimal 30;)\ncontoh: \"Yusuf Nur Wahid\"")
	}

	emailPattern, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if len(email) != 0 && !emailPattern.MatchString(email) {
		return errors.New("input email tidak sesuai (username: lowercase alfanumerik, simbol yang diperbolehkan hanya '_' dan '-'; pemisah: simbol '@'; domain name: lowercase alfanumerik, simbol yang diperbolehkan hanya '-'; pemisah: simbol '.'; domain: lowercase alfabet, minimal 2 karakter, maksimal 4 karakter;)\ncontoh: \"yusuf@mail.com\"")
	}

	passwordPattern, _ := regexp.Compile(`^[a-zA-Z0-9!@#$&()\\\-\x60.+,/\"]{5,8}$`)
	if len(password) != 0 && !passwordPattern.MatchString(password) {
		return errors.New("input password tidak sesuai (alfanumerik; boleh menggunakan simbol; total karakter: minimal 5, maksimal 8;)")
	}
	return nil
}
