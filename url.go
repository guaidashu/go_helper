/**
  create by yy on 2019-09-26
*/

package go_helper

import (
	"net/url"
	"strings"
)

/**
This method can encode the url query params.
If the params have some special character This encode method maybe will return a error result,
please check your params and use with caution.
return string, error
*/
func EncodeGetUrlParam(urlOrigin string) (string, string, error) {
	var head string
	if strings.Contains(urlOrigin, "https://") {
		head = "https://"
	} else {
		head = "http://"
	}
	urlParse, err := url.Parse(urlOrigin)
	if err != nil {
		return "", "", err
	}
	query := urlParse.Query().Encode()
	urlOrigin = head + urlParse.Host + urlParse.Path + "?" + query
	return query, urlOrigin, err
}

/**
This method can encode the url query params.
urlParams example:
	urlParams := url.Values{}
	urlParams.Add("key", "value")
	query := EncodeGetUrlParamValues(urlParams)
return string
*/
func EncodeGetUrlParamValues(urlParams url.Values) string {
	body := urlParams.Encode()
	return body
}
