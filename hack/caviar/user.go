package caviar

import "net/url"

// SignIn authenticates a user for the current session.
func (s *Session) SignIn(username, password string) error {
	data := make(url.Values)
	data.Add("user[email]", username)
	data.Add("user[password]", password)

	_, err := s.PostForm("/users/sign_in", data)
	if err != nil {
		return err
	}
	return nil
}
