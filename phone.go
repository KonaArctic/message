package message
import "errors"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

// Sends out messages with TextNow. Probably illegal.
func ( self Message )Text( ) error {
	var err error
	var smstext string

	// TextNow Configs
	const apiurl string = "https://www.textnow.com/api/users/z_hbbq5zru9cxs7jhwih/messages"
	const cookie string = "connect.sid=s%3AYiagnsn5ON3035O732nLkjFN3BSqCQnT.Y%2FwnOdHUO%2FmM3uiHXZBErfWgg%2BCE3Dg8TNwOc97vMnM;"

	// Build message
	smstext = strings.Join( [ ]string{ self.About , self.Content , self.Link } , " " )
	smstext = strings.Replace( smstext , "\\" , "\\\\" , -1 )
	smstext = strings.Replace( smstext , "\"" , "\\\"" , -1 )
	smstext = url.QueryEscape( smstext )
	self.Receive = strings.Replace( self.Receive , "\\" , "\\\\" , -1 )
	self.Receive = strings.Replace( self.Receive , "\"" , "\\\"" , -1 )
	self.Receive = url.QueryEscape( self.Receive )
	smstext = strings.Join( [ ]string{ "json=%7B%22contact_value%22%3A%22" , self.Receive , "%22%2C%22message%22%3A%22" , smstext , "%22%7D" } , "" )

	// Send it
	var client http.Client
	var request http.Request
	request.Method = http.MethodPost
	request.URL = mu( url.Parse( apiurl ) )[ 0 ].( * url.URL )
	request.Header = map[ string ][ ]string{
		"Cookie" : { cookie } ,
		"Content-Type" : { "application/x-www-form-urlencoded" } , }
	request.Body = ioutil.NopCloser( strings.NewReader( smstext ) )
	request.ContentLength = int64( len( smstext ) )	// Official docs can be misleading
	response , err := client.Do( & request )
	if err != nil {
		return err }
	if response.StatusCode != http.StatusOK {
		return errors.New( "returned unexpected code: " + response.Status ) }

	return nil
}


