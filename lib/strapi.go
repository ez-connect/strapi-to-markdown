package lib

type Strapi struct {
	jwt string

	rest Rest
}

func (s *Strapi) SetBaseURL(value string) {
	s.rest.SetBaseURL(value)
}

func (s *Strapi) SetJWT(value string) {
	s.rest.SetJWT(value)
}

func (s *Strapi) SignIn(identifier, password string) (string, error) {
	body := map[string]string{
		"identifier": identifier,
		"password":   password,
	}

	res := map[string]string{}
	if err := s.rest.Post("auth/local", body, &res); err != nil {
		return "", err
	}

	s.jwt = body["jwt"]
	return s.jwt, nil
}

func (s *Strapi) Find(path string) ([]map[string]interface{}, error) {
	data := []map[string]interface{}{}
	err := s.rest.Get(path, &data)
	return data, err
}

func (s *Strapi) FindOne(path string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := s.rest.Get(path, &data)
	return data, err
}
