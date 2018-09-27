package julyNet

import "net/url"

func StrToUrl(strUrl string) (urlObj *url.URL, err error) {
	urlObj, err = url.Parse(strUrl)
	return urlObj, err
}

