/*
Документация для неофициального API кинопоиска (kinopoisk.dev).

 <!-- Yandex.Metrika counter --> <div><img src=\"https://mc.yandex.ru/watch/62307766\" style=\"position:absolute; left:-9999px;\" alt=\"\" /></div> <!-- /Yandex.Metrika counter --> <p>Через этот API вы можете получить практически все данные из кинопоиска. Больше информации вы можете получить изучив эту документацию.</p> <h2>Как работать с документацией?</h2> <p> Для начала работы с API вам необходимо получить токен, который вы можете получить в боте <a href=\"https://t.me/kinopoiskdev_bot\">@kinopoiskdev_bot</a>. <br /> После получения токена, вам необходимо авторизоваться в документации, для этого нажмите на кнопку <strong>Authorize</strong> и введите токен в поле <strong>Value</strong>.<br /> После авторизации вы можете отправлять запросы к API, для этого нажмите на кнопку <strong>Try it out</strong> и заполните необходимые поля для составления нужного фильтра.<br /> После заполнения полей нажмите на кнопку <strong>Execute</strong> и получите ответ от API и пример запроса. </p> <h3>Как работать с API?</h3> <p> API работает по принципу REST, все запросы отправляются на адрес <code>https://api.kinopoisk.dev/</code> с указанием версии API и необходимого ресурса.<br /> Все запросы к API кинопоиска должны содержать заголовок <code>X-API-KEY</code> с вашим токеном. В противном случае вы получите ошибку <code>401</code>.<br /> При составлении запроса учитывайте, что все параметры должны быть в <code>query</code> и <code>path</code>. В зависимости от метода который вы используете. Например, вы хотите получить список фильмов за 2023 год в жанре <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?year=2023&genres.name=криминал</code>. Или вы хотите получить список фильмов с рейтингом выше 8, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?rating.imdb=8-10</code>. Документация kinopoisk api может помочь вам составить нужный запрос, для этого воспользуйтесь ее конструктором. </p> <h3>Особенности синтекса query параметров</h3> <p> Ключи в query параметрах имеют разные типы значений. В зависимости от типа значения, вы можете использовать разные операторы для фильтрации для поиска максимально релевантного фильма, сериала и т.д. в базе. <br /> Поля с типом <code>Number</code> могут принимать значения в форматах: <code>rating.kp=1-10</code>, <code>rating.kp=1</code>, <code>year=2022&year=2023</code>. <br /> Поля с типом <code>Date</code> могут принимать значения в форматах: <code>premiere.russia=dd.mm.yyyy-dd.mm.yyyy</code>, <code>premiere.russia=dd.mm.yyyy</code>. <br /> Поля с типом <code>String</code> могут принимать значения в форматах: <code>genres.name=драма</code>, <code>genres.name=криминал</code>, <code>genres.name=криминал&genres.name=драма</code> <br/> Поля с типом <code>Boolean</code> могут принимать значения в форматах: <code>isSeries=true</code>, <code>isSeries=false</code>. <br /> Параметры жанров и стран могут принимать операторы <code>+</code> и <code>!</code>, для указания включаемых и исключаемых значений. Например, вы хотите получить список фильмов в жанрах <code>драма</code> и <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=+криминал</code>. Или вы хотите получить список фильмов с жанром <code>драма</code> и без жанра <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=!криминал</code>. <br /> </p> <p> Расшифровка операторов: <ul>   <li><code>!</code> - исключить. Этот символ нужно отправлять в кодированной форме <code>%21</code></li>   <li><code>+</code> - включить. Этот символ нужно отправлять в кодированной форме <code>%2B</code></li>   <li><code>-</code> - диапазон значений, используется в качестве разделителя.</li> </ul> </p>  <p>По вопросам работы с API обращайтесь в чат <a href=\"https://t.me/+jeHPZVXiLPFhODJi\">Developer Community KinopoiskDev</a>.</p>  <p>Если вы обнаружили ошибку или у вас есть предложения по улучшению, создавайте issue на <a href=\"https://github.com/mdwitr0/kinopoiskdev\">GitHub</a>.</p>  <h3>Полезные ссылки:</h3> <ul>   <li><a href=\"https://kinopoiskdev.readme.io\">Более удобная документация</a></li>   <li><a href=\"https://github.com/OpenMovieDB/kinopoiskdev_client\">JavaScript и TypeScript клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"https://github.com/odi1n/kinopoisk_dev\">Python клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"/documentation-json\">OpenAPI Specification (JSON)</a></li>   <li><a href=\"/documentation-yaml\">OpenAPI Specification (YAML)</a></li> </ul> 

API version: 1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kinopoisk_api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
	ContextHttpSignatureAuth = contextKey("httpsignature")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "https://api.kinopoisk.dev",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("Index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("The variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
