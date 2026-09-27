package model

import "github.com/google/uuid"

type Token struct {
	verificationToken string
	verificationSalt  string
}

func NewEmptyToken() Token {
	return Token{}
}

func NewToken(vt, vs string) Token {
	t := Token{}

	t.SetVerificationToken(vt)
	t.SetVerificationSalt(vs)

	return t

}

func (t *Token) VerificationToken() string      { return t.verificationToken }
func (t *Token) SetVerificationToken(vt string) { t.verificationToken = vt }

func (t *Token) VerificationSalt() string      { return t.verificationSalt }
func (t *Token) SetVerificationSalt(vs string) { t.verificationSalt = vs }

func (t *Token) IsValid() bool {
	return true
}

func (t *Token) RefreshVerificationToken() {
	vt := uuid.NewString()
	t.SetVerificationToken(vt)
}

func (t *Token) RefreshVerificationSalt() {
	vs := uuid.NewString()
	t.SetVerificationSalt(vs)
}
