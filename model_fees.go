/*
Документация для неофициального API кинопоиска (kinopoisk.dev).

 <!-- Yandex.Metrika counter --> <div><img src=\"https://mc.yandex.ru/watch/62307766\" style=\"position:absolute; left:-9999px;\" alt=\"\" /></div> <!-- /Yandex.Metrika counter --> <p>Через этот API вы можете получить практически все данные из кинопоиска. Больше информации вы можете получить изучив эту документацию.</p> <h2>Как работать с документацией?</h2> <p> Для начала работы с API вам необходимо получить токен, который вы можете получить в боте <a href=\"https://t.me/kinopoiskdev_bot\">@kinopoiskdev_bot</a>. <br /> После получения токена, вам необходимо авторизоваться в документации, для этого нажмите на кнопку <strong>Authorize</strong> и введите токен в поле <strong>Value</strong>.<br /> После авторизации вы можете отправлять запросы к API, для этого нажмите на кнопку <strong>Try it out</strong> и заполните необходимые поля для составления нужного фильтра.<br /> После заполнения полей нажмите на кнопку <strong>Execute</strong> и получите ответ от API и пример запроса. </p> <h3>Как работать с API?</h3> <p> API работает по принципу REST, все запросы отправляются на адрес <code>https://api.kinopoisk.dev/</code> с указанием версии API и необходимого ресурса.<br /> Все запросы к API кинопоиска должны содержать заголовок <code>X-API-KEY</code> с вашим токеном. В противном случае вы получите ошибку <code>401</code>.<br /> При составлении запроса учитывайте, что все параметры должны быть в <code>query</code> и <code>path</code>. В зависимости от метода который вы используете. Например, вы хотите получить список фильмов за 2023 год в жанре <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?year=2023&genres.name=криминал</code>. Или вы хотите получить список фильмов с рейтингом выше 8, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?rating.imdb=8-10</code>. Документация kinopoisk api может помочь вам составить нужный запрос, для этого воспользуйтесь ее конструктором. </p> <h3>Особенности синтекса query параметров</h3> <p> Ключи в query параметрах имеют разные типы значений. В зависимости от типа значения, вы можете использовать разные операторы для фильтрации для поиска максимально релевантного фильма, сериала и т.д. в базе. <br /> Поля с типом <code>Number</code> могут принимать значения в форматах: <code>rating.kp=1-10</code>, <code>rating.kp=1</code>, <code>year=2022&year=2023</code>. <br /> Поля с типом <code>Date</code> могут принимать значения в форматах: <code>premiere.russia=dd.mm.yyyy-dd.mm.yyyy</code>, <code>premiere.russia=dd.mm.yyyy</code>. <br /> Поля с типом <code>String</code> могут принимать значения в форматах: <code>genres.name=драма</code>, <code>genres.name=криминал</code>, <code>genres.name=криминал&genres.name=драма</code> <br/> Поля с типом <code>Boolean</code> могут принимать значения в форматах: <code>isSeries=true</code>, <code>isSeries=false</code>. <br /> Параметры жанров и стран могут принимать операторы <code>+</code> и <code>!</code>, для указания включаемых и исключаемых значений. Например, вы хотите получить список фильмов в жанрах <code>драма</code> и <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=+криминал</code>. Или вы хотите получить список фильмов с жанром <code>драма</code> и без жанра <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=!криминал</code>. <br /> </p> <p> Расшифровка операторов: <ul>   <li><code>!</code> - исключить. Этот символ нужно отправлять в кодированной форме <code>%21</code></li>   <li><code>+</code> - включить. Этот символ нужно отправлять в кодированной форме <code>%2B</code></li>   <li><code>-</code> - диапазон значений, используется в качестве разделителя.</li> </ul> </p>  <p>По вопросам работы с API обращайтесь в чат <a href=\"https://t.me/+jeHPZVXiLPFhODJi\">Developer Community KinopoiskDev</a>.</p>  <p>Если вы обнаружили ошибку или у вас есть предложения по улучшению, создавайте issue на <a href=\"https://github.com/mdwitr0/kinopoiskdev\">GitHub</a>.</p>  <h3>Полезные ссылки:</h3> <ul>   <li><a href=\"https://kinopoiskdev.readme.io\">Более удобная документация</a></li>   <li><a href=\"https://github.com/OpenMovieDB/kinopoiskdev_client\">JavaScript и TypeScript клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"https://github.com/odi1n/kinopoisk_dev\">Python клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"/documentation-json\">OpenAPI Specification (JSON)</a></li>   <li><a href=\"/documentation-yaml\">OpenAPI Specification (YAML)</a></li> </ul> 

API version: 1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kinopoisk_api

import (
	"encoding/json"
)

// Fees struct for Fees
type Fees struct {
	World *CurrencyValue `json:"world,omitempty"`
	Russia *CurrencyValue `json:"russia,omitempty"`
	Usa *CurrencyValue `json:"usa,omitempty"`
}

// NewFees instantiates a new Fees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFees() *Fees {
	this := Fees{}
	return &this
}

// NewFeesWithDefaults instantiates a new Fees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesWithDefaults() *Fees {
	this := Fees{}
	return &this
}

// GetWorld returns the World field value if set, zero value otherwise.
func (o *Fees) GetWorld() CurrencyValue {
	if o == nil || o.World == nil {
		var ret CurrencyValue
		return ret
	}
	return *o.World
}

// GetWorldOk returns a tuple with the World field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fees) GetWorldOk() (*CurrencyValue, bool) {
	if o == nil || o.World == nil {
		return nil, false
	}
	return o.World, true
}

// HasWorld returns a boolean if a field has been set.
func (o *Fees) HasWorld() bool {
	if o != nil && o.World != nil {
		return true
	}

	return false
}

// SetWorld gets a reference to the given CurrencyValue and assigns it to the World field.
func (o *Fees) SetWorld(v CurrencyValue) {
	o.World = &v
}

// GetRussia returns the Russia field value if set, zero value otherwise.
func (o *Fees) GetRussia() CurrencyValue {
	if o == nil || o.Russia == nil {
		var ret CurrencyValue
		return ret
	}
	return *o.Russia
}

// GetRussiaOk returns a tuple with the Russia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fees) GetRussiaOk() (*CurrencyValue, bool) {
	if o == nil || o.Russia == nil {
		return nil, false
	}
	return o.Russia, true
}

// HasRussia returns a boolean if a field has been set.
func (o *Fees) HasRussia() bool {
	if o != nil && o.Russia != nil {
		return true
	}

	return false
}

// SetRussia gets a reference to the given CurrencyValue and assigns it to the Russia field.
func (o *Fees) SetRussia(v CurrencyValue) {
	o.Russia = &v
}

// GetUsa returns the Usa field value if set, zero value otherwise.
func (o *Fees) GetUsa() CurrencyValue {
	if o == nil || o.Usa == nil {
		var ret CurrencyValue
		return ret
	}
	return *o.Usa
}

// GetUsaOk returns a tuple with the Usa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fees) GetUsaOk() (*CurrencyValue, bool) {
	if o == nil || o.Usa == nil {
		return nil, false
	}
	return o.Usa, true
}

// HasUsa returns a boolean if a field has been set.
func (o *Fees) HasUsa() bool {
	if o != nil && o.Usa != nil {
		return true
	}

	return false
}

// SetUsa gets a reference to the given CurrencyValue and assigns it to the Usa field.
func (o *Fees) SetUsa(v CurrencyValue) {
	o.Usa = &v
}

func (o Fees) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.World != nil {
		toSerialize["world"] = o.World
	}
	if o.Russia != nil {
		toSerialize["russia"] = o.Russia
	}
	if o.Usa != nil {
		toSerialize["usa"] = o.Usa
	}
	return json.Marshal(toSerialize)
}

type NullableFees struct {
	value *Fees
	isSet bool
}

func (v NullableFees) Get() *Fees {
	return v.value
}

func (v *NullableFees) Set(val *Fees) {
	v.value = val
	v.isSet = true
}

func (v NullableFees) IsSet() bool {
	return v.isSet
}

func (v *NullableFees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFees(val *Fees) *NullableFees {
	return &NullableFees{value: val, isSet: true}
}

func (v NullableFees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


