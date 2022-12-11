package vo

type Password string

func NewPassword(value string) (*Password, error) {
	password := Password(value)

	return &password, nil
}
