package simphttp_test

import (
    "testing"
    "github.com/wambosa/expect"
    "github.com/wambosa/simphttp"
)

func Test_GIVEN_url_is_valid_WHEN_GetString_is_called_THEN_returns_nonEmpty_string(t *testing.T){
    
    unparsed, err := simphttp.GetString("http://google.com")
    
    if err != nil {t.Error(err)}
    
    expecting := expect.TestCase {
        T: t,
        Value: unparsed,
    }
    
    expecting.Truthy()
}

func Test_GIVEN_valid_endpoint_returns_json_by_default_WHEN_GetJson_is_called_THEN_returns_map(t *testing.T){
    
    parsedMap, err := simphttp.GetJson("https://customtaxapp.com/get.test")
    
    if err != nil {t.Error(err)}
    
    expecting := expect.TestCase {
        T: t,
        Value: parsedMap,
    }
    
    expecting.Property("ok").Truthy()
    expecting.Property("version").Truthy()
}