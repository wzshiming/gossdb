package ssdb

// Auth password
// Available since: 1.7.0.0
// Authenticate the connection.
// Warning: The password is sent in plain-text over the network!
func (c *Client) Auth(password string) error {
	return c.doNil("auth", password)
}

// DBSize Return the approximate size of the database, in bytes. If compression is enabled, the size will be of the compressed data.
func (c *Client) DBSize() (int64, error) {
	return c.doInt("dbsize")
}

// Info Return information about the server.
func (c *Client) Info() ([]string, error) {
	return c.doStrings("info")
}
