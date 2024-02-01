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

// YearRange struct for YearRange
type YearRange struct {
	// Год начала
	Start NullableFloat32 `json:"start,omitempty"`
	// Год окончания
	End NullableFloat32 `json:"end,omitempty"`
}

// NewYearRange instantiates a new YearRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYearRange() *YearRange {
	this := YearRange{}
	return &this
}

// NewYearRangeWithDefaults instantiates a new YearRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYearRangeWithDefaults() *YearRange {
	this := YearRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YearRange) GetStart() float32 {
	if o == nil || o.Start.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YearRange) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *YearRange) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableFloat32 and assigns it to the Start field.
func (o *YearRange) SetStart(v float32) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *YearRange) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *YearRange) UnsetStart() {
	o.Start.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YearRange) GetEnd() float32 {
	if o == nil || o.End.Get() == nil {
		var ret float32
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YearRange) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *YearRange) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableFloat32 and assigns it to the End field.
func (o *YearRange) SetEnd(v float32) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *YearRange) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *YearRange) UnsetEnd() {
	o.End.Unset()
}

func (o YearRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableYearRange struct {
	value *YearRange
	isSet bool
}

func (v NullableYearRange) Get() *YearRange {
	return v.value
}

func (v *NullableYearRange) Set(val *YearRange) {
	v.value = val
	v.isSet = true
}

func (v NullableYearRange) IsSet() bool {
	return v.isSet
}

func (v *NullableYearRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYearRange(val *YearRange) *NullableYearRange {
	return &NullableYearRange{value: val, isSet: true}
}

func (v NullableYearRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYearRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

