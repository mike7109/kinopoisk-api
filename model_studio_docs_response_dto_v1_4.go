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

// StudioDocsResponseDtoV14 struct for StudioDocsResponseDtoV14
type StudioDocsResponseDtoV14 struct {
	Docs []Studio `json:"docs"`
	// Общее количество результатов
	Total float32 `json:"total"`
	// Количество результатов на странице
	Limit float32 `json:"limit"`
	// Текущая страница
	Page float32 `json:"page"`
	// Сколько страниц всего
	Pages float32 `json:"pages"`
}

// NewStudioDocsResponseDtoV14 instantiates a new StudioDocsResponseDtoV14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudioDocsResponseDtoV14(docs []Studio, total float32, limit float32, page float32, pages float32) *StudioDocsResponseDtoV14 {
	this := StudioDocsResponseDtoV14{}
	this.Docs = docs
	this.Total = total
	this.Limit = limit
	this.Page = page
	this.Pages = pages
	return &this
}

// NewStudioDocsResponseDtoV14WithDefaults instantiates a new StudioDocsResponseDtoV14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudioDocsResponseDtoV14WithDefaults() *StudioDocsResponseDtoV14 {
	this := StudioDocsResponseDtoV14{}
	return &this
}

// GetDocs returns the Docs field value
func (o *StudioDocsResponseDtoV14) GetDocs() []Studio {
	if o == nil {
		var ret []Studio
		return ret
	}

	return o.Docs
}

// GetDocsOk returns a tuple with the Docs field value
// and a boolean to check if the value has been set.
func (o *StudioDocsResponseDtoV14) GetDocsOk() ([]Studio, bool) {
	if o == nil {
		return nil, false
	}
	return o.Docs, true
}

// SetDocs sets field value
func (o *StudioDocsResponseDtoV14) SetDocs(v []Studio) {
	o.Docs = v
}

// GetTotal returns the Total field value
func (o *StudioDocsResponseDtoV14) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *StudioDocsResponseDtoV14) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *StudioDocsResponseDtoV14) SetTotal(v float32) {
	o.Total = v
}

// GetLimit returns the Limit field value
func (o *StudioDocsResponseDtoV14) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *StudioDocsResponseDtoV14) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *StudioDocsResponseDtoV14) SetLimit(v float32) {
	o.Limit = v
}

// GetPage returns the Page field value
func (o *StudioDocsResponseDtoV14) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *StudioDocsResponseDtoV14) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *StudioDocsResponseDtoV14) SetPage(v float32) {
	o.Page = v
}

// GetPages returns the Pages field value
func (o *StudioDocsResponseDtoV14) GetPages() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *StudioDocsResponseDtoV14) GetPagesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *StudioDocsResponseDtoV14) SetPages(v float32) {
	o.Pages = v
}

func (o StudioDocsResponseDtoV14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["docs"] = o.Docs
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["pages"] = o.Pages
	}
	return json.Marshal(toSerialize)
}

type NullableStudioDocsResponseDtoV14 struct {
	value *StudioDocsResponseDtoV14
	isSet bool
}

func (v NullableStudioDocsResponseDtoV14) Get() *StudioDocsResponseDtoV14 {
	return v.value
}

func (v *NullableStudioDocsResponseDtoV14) Set(val *StudioDocsResponseDtoV14) {
	v.value = val
	v.isSet = true
}

func (v NullableStudioDocsResponseDtoV14) IsSet() bool {
	return v.isSet
}

func (v *NullableStudioDocsResponseDtoV14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudioDocsResponseDtoV14(val *StudioDocsResponseDtoV14) *NullableStudioDocsResponseDtoV14 {
	return &NullableStudioDocsResponseDtoV14{value: val, isSet: true}
}

func (v NullableStudioDocsResponseDtoV14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudioDocsResponseDtoV14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


