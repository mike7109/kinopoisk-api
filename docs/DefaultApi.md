# \DefaultApi

All URIs are relative to *https://api.kinopoisk.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageControllerFindManyV14**](DefaultApi.md#ImageControllerFindManyV14) | **Get** /v1.4/image | Поиск картинок
[**KeywordControllerFindManyV14**](DefaultApi.md#KeywordControllerFindManyV14) | **Get** /v1.4/keyword | Поиск ключевых слов
[**ListControllerFindManyV14**](DefaultApi.md#ListControllerFindManyV14) | **Get** /v1.4/list | Поиск коллекций
[**ListControllerFindOneV14**](DefaultApi.md#ListControllerFindOneV14) | **Get** /v1.4/list/{slug} | Поиск коллекции по slug
[**MovieControllerFindManyAwardsV14**](DefaultApi.md#MovieControllerFindManyAwardsV14) | **Get** /v1.4/movie/awards | Награды тайтлов
[**MovieControllerFindManyByQueryV14**](DefaultApi.md#MovieControllerFindManyByQueryV14) | **Get** /v1.4/movie | Универсальный поиск с фильтрами
[**MovieControllerFindOneV14**](DefaultApi.md#MovieControllerFindOneV14) | **Get** /v1.4/movie/{id} | Поиск по id
[**MovieControllerGetPossibleValuesByFieldName**](DefaultApi.md#MovieControllerGetPossibleValuesByFieldName) | **Get** /v1/movie/possible-values-by-field | Получить список стран, жанров, и т.д.
[**MovieControllerGetRandomMovieV14**](DefaultApi.md#MovieControllerGetRandomMovieV14) | **Get** /v1.4/movie/random | Получить рандомный тайтл из базы
[**MovieControllerSearchMovieV14**](DefaultApi.md#MovieControllerSearchMovieV14) | **Get** /v1.4/movie/search | Поиск фильмов по названию
[**PersonControllerFindManyAwardsV14**](DefaultApi.md#PersonControllerFindManyAwardsV14) | **Get** /v1.4/person/awards | Награды актеров
[**PersonControllerFindManyV14**](DefaultApi.md#PersonControllerFindManyV14) | **Get** /v1.4/person | Универсальный поиск с фильтрами
[**PersonControllerFindOneV14**](DefaultApi.md#PersonControllerFindOneV14) | **Get** /v1.4/person/{id} | Поиск по id
[**PersonControllerSearchPersonV14**](DefaultApi.md#PersonControllerSearchPersonV14) | **Get** /v1.4/person/search | Поиск актеров, режиссеров, и т.д по имени
[**ReviewControllerFindManyV14**](DefaultApi.md#ReviewControllerFindManyV14) | **Get** /v1.4/review | Универсальный поиск с фильтрами
[**SeasonControllerFindManyV14**](DefaultApi.md#SeasonControllerFindManyV14) | **Get** /v1.4/season | Поиск сезонов
[**StudioControllerFindManyV14**](DefaultApi.md#StudioControllerFindManyV14) | **Get** /v1.4/studio | Поиск студий



## ImageControllerFindManyV14

> ImageDocsResponseDtoV14 ImageControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).Type_(type_).Language(language).Height(height).Width(width).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Поиск картинок



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    movieId := []string{"Inner_example"} // []string | Поиск картинок по id фильма (пример: `\"666\", \"!666\"`) (optional)
    type_ := []string{"Type_example"} // []string | Поиск картинок по типу (пример: `\"cover\", \"!cover\"`) (optional)
    language := []string{"Language_example"} // []string | Поиск картинок по языку (пример: `\"en\", \"!de\"`) (optional)
    height := []string{"Inner_example"} // []string | Поиск картинок по высоте (пример: `\"1920\", \"360-1920\"`) (optional)
    width := []string{"Inner_example"} // []string | Поиск картинок по ширине (пример: `\"1080\", \"320-1080\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImageControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).Type_(type_).Language(language).Height(height).Width(width).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImageControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageControllerFindManyV14`: ImageDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImageControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImageControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **movieId** | **[]string** | Поиск картинок по id фильма (пример: &#x60;\&quot;666\&quot;, \&quot;!666\&quot;&#x60;) | 
 **type_** | **[]string** | Поиск картинок по типу (пример: &#x60;\&quot;cover\&quot;, \&quot;!cover\&quot;&#x60;) | 
 **language** | **[]string** | Поиск картинок по языку (пример: &#x60;\&quot;en\&quot;, \&quot;!de\&quot;&#x60;) | 
 **height** | **[]string** | Поиск картинок по высоте (пример: &#x60;\&quot;1920\&quot;, \&quot;360-1920\&quot;&#x60;) | 
 **width** | **[]string** | Поиск картинок по ширине (пример: &#x60;\&quot;1080\&quot;, \&quot;320-1080\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**ImageDocsResponseDtoV14**](ImageDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeywordControllerFindManyV14

> KeywordDocsResponseDtoV14 KeywordControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Title(title).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Поиск ключевых слов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    id := []string{"Inner_example"} // []string | Поиск ключевого слова по id (пример: `\"666\", \"!666\"`) (optional)
    moviesId := []string{"Inner_example"} // []string | Поиск ключевых слов по id фильма (пример: `\"666\", \"!666\"`) (optional)
    title := []string{"Inner_example"} // []string | Поиск ключевых слов по наименованию (пример: `\"1980-е\", \"!1980-е\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.KeywordControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Title(title).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.KeywordControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeywordControllerFindManyV14`: KeywordDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.KeywordControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeywordControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **id** | **[]string** | Поиск ключевого слова по id (пример: &#x60;\&quot;666\&quot;, \&quot;!666\&quot;&#x60;) | 
 **moviesId** | **[]string** | Поиск ключевых слов по id фильма (пример: &#x60;\&quot;666\&quot;, \&quot;!666\&quot;&#x60;) | 
 **title** | **[]string** | Поиск ключевых слов по наименованию (пример: &#x60;\&quot;1980-е\&quot;, \&quot;!1980-е\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**KeywordDocsResponseDtoV14**](KeywordDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListControllerFindManyV14

> ListDocsResponseDtoV14 ListControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Slug(slug).Category(category).MoviesCount(moviesCount).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Поиск коллекций



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    slug := []string{"Inner_example"} // []string | Поиск slug (пример: `\"!top250\", \"top250\"`) (optional)
    category := []string{"Category_example"} // []string | Поиск по категории (пример: `\"Фильмы\", \"!Фильмы\"`) (optional)
    moviesCount := []string{"Inner_example"} // []string | Поиск по количеству фильмов (пример: `\"1-200\", \"10\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Slug(slug).Category(category).MoviesCount(moviesCount).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListControllerFindManyV14`: ListDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **slug** | **[]string** | Поиск slug (пример: &#x60;\&quot;!top250\&quot;, \&quot;top250\&quot;&#x60;) | 
 **category** | **[]string** | Поиск по категории (пример: &#x60;\&quot;Фильмы\&quot;, \&quot;!Фильмы\&quot;&#x60;) | 
 **moviesCount** | **[]string** | Поиск по количеству фильмов (пример: &#x60;\&quot;1-200\&quot;, \&quot;10\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**ListDocsResponseDtoV14**](ListDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListControllerFindOneV14

> List ListControllerFindOneV14(ctx, slug).Execute()

Поиск коллекции по slug



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListControllerFindOneV14(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListControllerFindOneV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListControllerFindOneV14`: List
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListControllerFindOneV14`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListControllerFindOneV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**List**](List.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerFindManyAwardsV14

> MovieAwardDocsResponseDto MovieControllerFindManyAwardsV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).NominationTitle(nominationTitle).NominationAwardTitle(nominationAwardTitle).NominationAwardYear(nominationAwardYear).Winning(winning).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Награды тайтлов

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    movieId := []string{"Inner_example"} // []string | Поиск по ID фильма (пример: `\"666\", \"555\", \"!666\"`) (optional)
    nominationTitle := []string{"Inner_example"} // []string | Поиск по номинациям (пример: `\"Оскар\", \"Золотой глобус\"`) (optional)
    nominationAwardTitle := []string{"Inner_example"} // []string | Поиск по наградам (пример: `\"Лучший фильм\", \"Лучший актер\"`) (optional)
    nominationAwardYear := []string{"Inner_example"} // []string | Поиск по году награды (пример: `\"2019\", \"2020\"`) (optional)
    winning := []string{"Inner_example"} // []string | Поиск по победам (пример: `\"true\", \"false\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerFindManyAwardsV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).NominationTitle(nominationTitle).NominationAwardTitle(nominationAwardTitle).NominationAwardYear(nominationAwardYear).Winning(winning).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerFindManyAwardsV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerFindManyAwardsV14`: MovieAwardDocsResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerFindManyAwardsV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerFindManyAwardsV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **movieId** | **[]string** | Поиск по ID фильма (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **nominationTitle** | **[]string** | Поиск по номинациям (пример: &#x60;\&quot;Оскар\&quot;, \&quot;Золотой глобус\&quot;&#x60;) | 
 **nominationAwardTitle** | **[]string** | Поиск по наградам (пример: &#x60;\&quot;Лучший фильм\&quot;, \&quot;Лучший актер\&quot;&#x60;) | 
 **nominationAwardYear** | **[]string** | Поиск по году награды (пример: &#x60;\&quot;2019\&quot;, \&quot;2020\&quot;&#x60;) | 
 **winning** | **[]string** | Поиск по победам (пример: &#x60;\&quot;true\&quot;, \&quot;false\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**MovieAwardDocsResponseDto**](MovieAwardDocsResponseDto.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerFindManyByQueryV14

> MovieDocsResponseDtoV14 MovieControllerFindManyByQueryV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).ExternalIdImdb(externalIdImdb).ExternalIdTmdb(externalIdTmdb).ExternalIdKpHD(externalIdKpHD).Type_(type_).TypeNumber(typeNumber).IsSeries(isSeries).Status(status).Year(year).ReleaseYearsStart(releaseYearsStart).ReleaseYearsEnd(releaseYearsEnd).RatingKp(ratingKp).RatingImdb(ratingImdb).RatingTmdb(ratingTmdb).RatingMpaa(ratingMpaa).AgeRating(ageRating).VotesKp(votesKp).VotesImdb(votesImdb).VotesTmdb(votesTmdb).VotesFilmCritics(votesFilmCritics).VotesRussianFilmCritics(votesRussianFilmCritics).VotesAwait(votesAwait).BudgetValue(budgetValue).AudienceCount(audienceCount).MovieLength(movieLength).SeriesLength(seriesLength).TotalSeriesLength(totalSeriesLength).GenresName(genresName).CountriesName(countriesName).TicketsOnSale(ticketsOnSale).NetworksItemsName(networksItemsName).PersonsId(personsId).PersonsProfession(personsProfession).PersonsEnProfession(personsEnProfession).FeesWorld(feesWorld).FeesUsa(feesUsa).FeesRussia(feesRussia).PremiereWorld(premiereWorld).PremiereUsa(premiereUsa).PremiereRussia(premiereRussia).PremiereDigital(premiereDigital).PremiereCinema(premiereCinema).PremiereCountry(premiereCountry).SimilarMoviesId(similarMoviesId).SequelsAndPrequelsId(sequelsAndPrequelsId).WatchabilityItemsName(watchabilityItemsName).Lists(lists).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Универсальный поиск с фильтрами



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    id := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk (пример: `\"666\", \"555\", \"!666\"`) (optional)
    externalIdImdb := []string{"Inner_example"} // []string | Поиск по IMDB ID (пример: `\"tt666\", \"tt555\", \"!tt666\"`) (optional)
    externalIdTmdb := []float32{float32(123)} // []float32 | Поиск по TMDB ID (пример: `666, 555, !666`) (optional)
    externalIdKpHD := []string{"Inner_example"} // []string | Поиск по id KinoPoisk HD (пример: `\"48e8d0acb0f62d8585101798eaeceec5\", \"!48e8d0acb0f62d8585101798eaeceec5\"`) (optional)
    type_ := []string{"Type_example"} // []string | Поиск по типу фильма (пример: `\"movie\", \"tv-series\", \"!anime\"`) (optional)
    typeNumber := []string{"Inner_example"} // []string | Поиск по номеру типа фильма (пример: `1, 2, !3`). Список типов: 1 (movie), 2 (tv-series), 3 (cartoon), 4 (anime), 5 (animated-series). (optional)
    isSeries := []string{"Inner_example"} // []string | Поиск по индикатору сериала (пример: `true, false`) (optional)
    status := []string{"Status_example"} // []string | Поиск по статусу фильма (пример: `\"announced\", \"completed\", \"!filming\"`) (optional)
    year := []string{"Inner_example"} // []string | Поиск по году (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    releaseYearsStart := []string{"Inner_example"} // []string | Поиск по года начала релиза (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    releaseYearsEnd := []string{"Inner_example"} // []string | Поиск по года окончания релиза (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    ratingKp := []string{"Inner_example"} // []string | Поиск по рейтингу Кинопоиск (пример: `7, 10, 7.2-10`) (optional)
    ratingImdb := []string{"Inner_example"} // []string | Поиск по рейтингу IMDB (пример: `7, 10, 7.2-10`) (optional)
    ratingTmdb := []string{"Inner_example"} // []string | Поиск по рейтингу TMDB (пример: `7, 10, 7.2-10`) (optional)
    ratingMpaa := []string{"Inner_example"} // []string | Поиск по рейтингу MPAA (пример: `\"G\", \"NC-17\", \"!R\"`) (optional)
    ageRating := []string{"Inner_example"} // []string | Поиск по возрастному рейтингу (пример: `12, !18, 12-18`) (optional)
    votesKp := []string{"Inner_example"} // []string | Поиск по количеству голосов на KP (пример: `1000-6666666`) (optional)
    votesImdb := []string{"Inner_example"} // []string | Поиск по количеству голосов на IMDB (пример: `1000-6666666`) (optional)
    votesTmdb := []string{"Inner_example"} // []string | Поиск по количеству голосов на TMDB (пример: `1000-6666666`) (optional)
    votesFilmCritics := []string{"Inner_example"} // []string | Поиск по количеству голосов кинокритиков (пример: `1000-6666666`) (optional)
    votesRussianFilmCritics := []string{"Inner_example"} // []string | Поиск по количеству голосов кинокритиков из России (пример: `1000-6666666`) (optional)
    votesAwait := []string{"Inner_example"} // []string | Поиск по количеству голосов ожидания на Кинопоиске (пример: `1000-6666666`) (optional)
    budgetValue := []string{"Inner_example"} // []string | Поиск по бюджету фильма (пример: `1000-6666666`) (optional)
    audienceCount := []string{"Inner_example"} // []string | Поиск по количеству аудитории (пример: `1000-6666666`) (optional)
    movieLength := []string{"Inner_example"} // []string | Поиск по продолжительности фильма (пример: `100-120`) (optional)
    seriesLength := []string{"Inner_example"} // []string | Поиск по всей продолжительности одной серии (пример: `20-60`) (optional)
    totalSeriesLength := []string{"Inner_example"} // []string | Поиск по всей продолжительности сериала (пример: `100-120`) (optional)
    genresName := []string{"Inner_example"} // []string | Поиск по жанрам (пример: `\"драма\", \"комедия\", \"!мелодрама\", \"+ужасы\"`) (optional)
    countriesName := []string{"Inner_example"} // []string | Поиск по странам (пример: `\"США\", \"Россия\", \"!Франция\" , \"+Великобритания\"`) (optional)
    ticketsOnSale := []string{"Inner_example"} // []string | Поиск по наличию билетов в продаже (пример: `true, false`) (optional)
    networksItemsName := []string{"Inner_example"} // []string | Поиск по сетям производства фильма (пример: `\"HBO\", \"Netflix\", \"!Amazon\"`) (optional)
    personsId := []string{"Inner_example"} // []string | Поиск по ID персон (пример: `666, 555, !666`) (optional)
    personsProfession := []string{"Inner_example"} // []string | Поиск по профессиям персон (пример: `\"актер\", \"режиссер\", \"!сценарист\"`) (optional)
    personsEnProfession := []string{"Inner_example"} // []string | Поиск по английским профессиям персон (пример: `\"actor\", \"director\", \"!writer\"`) (optional)
    feesWorld := []string{"Inner_example"} // []string | Поиск по сборам в мире (пример: `1000-6666666`) (optional)
    feesUsa := []string{"Inner_example"} // []string | Поиск по сборам в США (пример: `1000-6666666`) (optional)
    feesRussia := []string{"Inner_example"} // []string | Поиск по сборам в России (пример: `1000-6666666`) (optional)
    premiereWorld := []string{"Inner_example"} // []string | Поиск по дате премьеры в мире (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereUsa := []string{"Inner_example"} // []string | Поиск по дате премьеры в США (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereRussia := []string{"Inner_example"} // []string | Поиск по дате премьеры в России (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereDigital := []string{"Inner_example"} // []string | Поиск по дате премьеры в стриминговых сервисах (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereCinema := []string{"Inner_example"} // []string | Поиск по дате премьеры в кинотеатрах (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereCountry := []string{"Inner_example"} // []string | Поиск по стране премьеры (пример: `\"США\", \"Россия\", \"!Франция\" , \"+Великобритания\"`) (optional)
    similarMoviesId := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk из списка похожих фильмов (пример: `666, 555, !666`) (optional)
    sequelsAndPrequelsId := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk из списка сиквелов и преквелов (пример: `666, 555, !666`) (optional)
    watchabilityItemsName := []string{"Inner_example"} // []string | Поиск по доуступным платформам для просмотра (пример: `\"ivi\", \"okko\", \"!megogo\"`) (optional)
    lists := []string{"Inner_example"} // []string | Поиск по коллекциям из KinoPoisk (пример: `\"top250\", \"top-100-indian-movies\", \"!top-100-movies\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerFindManyByQueryV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).ExternalIdImdb(externalIdImdb).ExternalIdTmdb(externalIdTmdb).ExternalIdKpHD(externalIdKpHD).Type_(type_).TypeNumber(typeNumber).IsSeries(isSeries).Status(status).Year(year).ReleaseYearsStart(releaseYearsStart).ReleaseYearsEnd(releaseYearsEnd).RatingKp(ratingKp).RatingImdb(ratingImdb).RatingTmdb(ratingTmdb).RatingMpaa(ratingMpaa).AgeRating(ageRating).VotesKp(votesKp).VotesImdb(votesImdb).VotesTmdb(votesTmdb).VotesFilmCritics(votesFilmCritics).VotesRussianFilmCritics(votesRussianFilmCritics).VotesAwait(votesAwait).BudgetValue(budgetValue).AudienceCount(audienceCount).MovieLength(movieLength).SeriesLength(seriesLength).TotalSeriesLength(totalSeriesLength).GenresName(genresName).CountriesName(countriesName).TicketsOnSale(ticketsOnSale).NetworksItemsName(networksItemsName).PersonsId(personsId).PersonsProfession(personsProfession).PersonsEnProfession(personsEnProfession).FeesWorld(feesWorld).FeesUsa(feesUsa).FeesRussia(feesRussia).PremiereWorld(premiereWorld).PremiereUsa(premiereUsa).PremiereRussia(premiereRussia).PremiereDigital(premiereDigital).PremiereCinema(premiereCinema).PremiereCountry(premiereCountry).SimilarMoviesId(similarMoviesId).SequelsAndPrequelsId(sequelsAndPrequelsId).WatchabilityItemsName(watchabilityItemsName).Lists(lists).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerFindManyByQueryV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerFindManyByQueryV14`: MovieDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerFindManyByQueryV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerFindManyByQueryV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **id** | **[]string** | Поиск по ID KinoPoisk (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **externalIdImdb** | **[]string** | Поиск по IMDB ID (пример: &#x60;\&quot;tt666\&quot;, \&quot;tt555\&quot;, \&quot;!tt666\&quot;&#x60;) | 
 **externalIdTmdb** | **[]float32** | Поиск по TMDB ID (пример: &#x60;666, 555, !666&#x60;) | 
 **externalIdKpHD** | **[]string** | Поиск по id KinoPoisk HD (пример: &#x60;\&quot;48e8d0acb0f62d8585101798eaeceec5\&quot;, \&quot;!48e8d0acb0f62d8585101798eaeceec5\&quot;&#x60;) | 
 **type_** | **[]string** | Поиск по типу фильма (пример: &#x60;\&quot;movie\&quot;, \&quot;tv-series\&quot;, \&quot;!anime\&quot;&#x60;) | 
 **typeNumber** | **[]string** | Поиск по номеру типа фильма (пример: &#x60;1, 2, !3&#x60;). Список типов: 1 (movie), 2 (tv-series), 3 (cartoon), 4 (anime), 5 (animated-series). | 
 **isSeries** | **[]string** | Поиск по индикатору сериала (пример: &#x60;true, false&#x60;) | 
 **status** | **[]string** | Поиск по статусу фильма (пример: &#x60;\&quot;announced\&quot;, \&quot;completed\&quot;, \&quot;!filming\&quot;&#x60;) | 
 **year** | **[]string** | Поиск по году (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **releaseYearsStart** | **[]string** | Поиск по года начала релиза (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **releaseYearsEnd** | **[]string** | Поиск по года окончания релиза (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **ratingKp** | **[]string** | Поиск по рейтингу Кинопоиск (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingImdb** | **[]string** | Поиск по рейтингу IMDB (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingTmdb** | **[]string** | Поиск по рейтингу TMDB (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingMpaa** | **[]string** | Поиск по рейтингу MPAA (пример: &#x60;\&quot;G\&quot;, \&quot;NC-17\&quot;, \&quot;!R\&quot;&#x60;) | 
 **ageRating** | **[]string** | Поиск по возрастному рейтингу (пример: &#x60;12, !18, 12-18&#x60;) | 
 **votesKp** | **[]string** | Поиск по количеству голосов на KP (пример: &#x60;1000-6666666&#x60;) | 
 **votesImdb** | **[]string** | Поиск по количеству голосов на IMDB (пример: &#x60;1000-6666666&#x60;) | 
 **votesTmdb** | **[]string** | Поиск по количеству голосов на TMDB (пример: &#x60;1000-6666666&#x60;) | 
 **votesFilmCritics** | **[]string** | Поиск по количеству голосов кинокритиков (пример: &#x60;1000-6666666&#x60;) | 
 **votesRussianFilmCritics** | **[]string** | Поиск по количеству голосов кинокритиков из России (пример: &#x60;1000-6666666&#x60;) | 
 **votesAwait** | **[]string** | Поиск по количеству голосов ожидания на Кинопоиске (пример: &#x60;1000-6666666&#x60;) | 
 **budgetValue** | **[]string** | Поиск по бюджету фильма (пример: &#x60;1000-6666666&#x60;) | 
 **audienceCount** | **[]string** | Поиск по количеству аудитории (пример: &#x60;1000-6666666&#x60;) | 
 **movieLength** | **[]string** | Поиск по продолжительности фильма (пример: &#x60;100-120&#x60;) | 
 **seriesLength** | **[]string** | Поиск по всей продолжительности одной серии (пример: &#x60;20-60&#x60;) | 
 **totalSeriesLength** | **[]string** | Поиск по всей продолжительности сериала (пример: &#x60;100-120&#x60;) | 
 **genresName** | **[]string** | Поиск по жанрам (пример: &#x60;\&quot;драма\&quot;, \&quot;комедия\&quot;, \&quot;!мелодрама\&quot;, \&quot;+ужасы\&quot;&#x60;) | 
 **countriesName** | **[]string** | Поиск по странам (пример: &#x60;\&quot;США\&quot;, \&quot;Россия\&quot;, \&quot;!Франция\&quot; , \&quot;+Великобритания\&quot;&#x60;) | 
 **ticketsOnSale** | **[]string** | Поиск по наличию билетов в продаже (пример: &#x60;true, false&#x60;) | 
 **networksItemsName** | **[]string** | Поиск по сетям производства фильма (пример: &#x60;\&quot;HBO\&quot;, \&quot;Netflix\&quot;, \&quot;!Amazon\&quot;&#x60;) | 
 **personsId** | **[]string** | Поиск по ID персон (пример: &#x60;666, 555, !666&#x60;) | 
 **personsProfession** | **[]string** | Поиск по профессиям персон (пример: &#x60;\&quot;актер\&quot;, \&quot;режиссер\&quot;, \&quot;!сценарист\&quot;&#x60;) | 
 **personsEnProfession** | **[]string** | Поиск по английским профессиям персон (пример: &#x60;\&quot;actor\&quot;, \&quot;director\&quot;, \&quot;!writer\&quot;&#x60;) | 
 **feesWorld** | **[]string** | Поиск по сборам в мире (пример: &#x60;1000-6666666&#x60;) | 
 **feesUsa** | **[]string** | Поиск по сборам в США (пример: &#x60;1000-6666666&#x60;) | 
 **feesRussia** | **[]string** | Поиск по сборам в России (пример: &#x60;1000-6666666&#x60;) | 
 **premiereWorld** | **[]string** | Поиск по дате премьеры в мире (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereUsa** | **[]string** | Поиск по дате премьеры в США (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereRussia** | **[]string** | Поиск по дате премьеры в России (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereDigital** | **[]string** | Поиск по дате премьеры в стриминговых сервисах (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereCinema** | **[]string** | Поиск по дате премьеры в кинотеатрах (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereCountry** | **[]string** | Поиск по стране премьеры (пример: &#x60;\&quot;США\&quot;, \&quot;Россия\&quot;, \&quot;!Франция\&quot; , \&quot;+Великобритания\&quot;&#x60;) | 
 **similarMoviesId** | **[]string** | Поиск по ID KinoPoisk из списка похожих фильмов (пример: &#x60;666, 555, !666&#x60;) | 
 **sequelsAndPrequelsId** | **[]string** | Поиск по ID KinoPoisk из списка сиквелов и преквелов (пример: &#x60;666, 555, !666&#x60;) | 
 **watchabilityItemsName** | **[]string** | Поиск по доуступным платформам для просмотра (пример: &#x60;\&quot;ivi\&quot;, \&quot;okko\&quot;, \&quot;!megogo\&quot;&#x60;) | 
 **lists** | **[]string** | Поиск по коллекциям из KinoPoisk (пример: &#x60;\&quot;top250\&quot;, \&quot;top-100-indian-movies\&quot;, \&quot;!top-100-movies\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**MovieDocsResponseDtoV14**](MovieDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerFindOneV14

> MovieDtoV14 MovieControllerFindOneV14(ctx, id).Execute()

Поиск по id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := float32(8.14) // float32 | ID из кинопоиска

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerFindOneV14(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerFindOneV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerFindOneV14`: MovieDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerFindOneV14`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID из кинопоиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerFindOneV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieDtoV14**](MovieDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerGetPossibleValuesByFieldName

> []PossibleValueDto MovieControllerGetPossibleValuesByFieldName(ctx).Field(field).Execute()

Получить список стран, жанров, и т.д.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    field := "field_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerGetPossibleValuesByFieldName(context.Background()).Field(field).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerGetPossibleValuesByFieldName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerGetPossibleValuesByFieldName`: []PossibleValueDto
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerGetPossibleValuesByFieldName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerGetPossibleValuesByFieldNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | **string** |  | 

### Return type

[**[]PossibleValueDto**](PossibleValueDto.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerGetRandomMovieV14

> MovieDtoV14 MovieControllerGetRandomMovieV14(ctx).NotNullFields(notNullFields).Id(id).ExternalIdImdb(externalIdImdb).ExternalIdTmdb(externalIdTmdb).ExternalIdKpHD(externalIdKpHD).Type_(type_).TypeNumber(typeNumber).IsSeries(isSeries).Status(status).Year(year).ReleaseYearsStart(releaseYearsStart).ReleaseYearsEnd(releaseYearsEnd).RatingKp(ratingKp).RatingImdb(ratingImdb).RatingTmdb(ratingTmdb).RatingMpaa(ratingMpaa).AgeRating(ageRating).VotesKp(votesKp).VotesImdb(votesImdb).VotesTmdb(votesTmdb).VotesFilmCritics(votesFilmCritics).VotesRussianFilmCritics(votesRussianFilmCritics).VotesAwait(votesAwait).BudgetValue(budgetValue).AudienceCount(audienceCount).MovieLength(movieLength).SeriesLength(seriesLength).TotalSeriesLength(totalSeriesLength).GenresName(genresName).CountriesName(countriesName).TicketsOnSale(ticketsOnSale).NetworksItemsName(networksItemsName).PersonsId(personsId).PersonsProfession(personsProfession).PersonsEnProfession(personsEnProfession).FeesWorld(feesWorld).FeesUsa(feesUsa).FeesRussia(feesRussia).PremiereWorld(premiereWorld).PremiereUsa(premiereUsa).PremiereRussia(premiereRussia).PremiereDigital(premiereDigital).PremiereCinema(premiereCinema).PremiereCountry(premiereCountry).SimilarMoviesId(similarMoviesId).SequelsAndPrequelsId(sequelsAndPrequelsId).WatchabilityItemsName(watchabilityItemsName).Lists(lists).Execute()

Получить рандомный тайтл из базы



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    id := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk (пример: `\"666\", \"555\", \"!666\"`) (optional)
    externalIdImdb := []string{"Inner_example"} // []string | Поиск по IMDB ID (пример: `\"tt666\", \"tt555\", \"!tt666\"`) (optional)
    externalIdTmdb := []float32{float32(123)} // []float32 | Поиск по TMDB ID (пример: `666, 555, !666`) (optional)
    externalIdKpHD := []string{"Inner_example"} // []string | Поиск по id KinoPoisk HD (пример: `\"48e8d0acb0f62d8585101798eaeceec5\", \"!48e8d0acb0f62d8585101798eaeceec5\"`) (optional)
    type_ := []string{"Type_example"} // []string | Поиск по типу фильма (пример: `\"movie\", \"tv-series\", \"!anime\"`) (optional)
    typeNumber := []string{"Inner_example"} // []string | Поиск по номеру типа фильма (пример: `1, 2, !3`). Список типов: 1 (movie), 2 (tv-series), 3 (cartoon), 4 (anime), 5 (animated-series). (optional)
    isSeries := []string{"Inner_example"} // []string | Поиск по индикатору сериала (пример: `true, false`) (optional)
    status := []string{"Status_example"} // []string | Поиск по статусу фильма (пример: `\"announced\", \"completed\", \"!filming\"`) (optional)
    year := []string{"Inner_example"} // []string | Поиск по году (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    releaseYearsStart := []string{"Inner_example"} // []string | Поиск по года начала релиза (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    releaseYearsEnd := []string{"Inner_example"} // []string | Поиск по года окончания релиза (пример: `1874, 2050, !2020, 2020-2024`) (optional)
    ratingKp := []string{"Inner_example"} // []string | Поиск по рейтингу Кинопоиск (пример: `7, 10, 7.2-10`) (optional)
    ratingImdb := []string{"Inner_example"} // []string | Поиск по рейтингу IMDB (пример: `7, 10, 7.2-10`) (optional)
    ratingTmdb := []string{"Inner_example"} // []string | Поиск по рейтингу TMDB (пример: `7, 10, 7.2-10`) (optional)
    ratingMpaa := []string{"Inner_example"} // []string | Поиск по рейтингу MPAA (пример: `\"G\", \"NC-17\", \"!R\"`) (optional)
    ageRating := []string{"Inner_example"} // []string | Поиск по возрастному рейтингу (пример: `12, !18, 12-18`) (optional)
    votesKp := []string{"Inner_example"} // []string | Поиск по количеству голосов на KP (пример: `1000-6666666`) (optional)
    votesImdb := []string{"Inner_example"} // []string | Поиск по количеству голосов на IMDB (пример: `1000-6666666`) (optional)
    votesTmdb := []string{"Inner_example"} // []string | Поиск по количеству голосов на TMDB (пример: `1000-6666666`) (optional)
    votesFilmCritics := []string{"Inner_example"} // []string | Поиск по количеству голосов кинокритиков (пример: `1000-6666666`) (optional)
    votesRussianFilmCritics := []string{"Inner_example"} // []string | Поиск по количеству голосов кинокритиков из России (пример: `1000-6666666`) (optional)
    votesAwait := []string{"Inner_example"} // []string | Поиск по количеству голосов ожидания на Кинопоиске (пример: `1000-6666666`) (optional)
    budgetValue := []string{"Inner_example"} // []string | Поиск по бюджету фильма (пример: `1000-6666666`) (optional)
    audienceCount := []string{"Inner_example"} // []string | Поиск по количеству аудитории (пример: `1000-6666666`) (optional)
    movieLength := []string{"Inner_example"} // []string | Поиск по продолжительности фильма (пример: `100-120`) (optional)
    seriesLength := []string{"Inner_example"} // []string | Поиск по всей продолжительности одной серии (пример: `20-60`) (optional)
    totalSeriesLength := []string{"Inner_example"} // []string | Поиск по всей продолжительности сериала (пример: `100-120`) (optional)
    genresName := []string{"Inner_example"} // []string | Поиск по жанрам (пример: `\"драма\", \"комедия\", \"!мелодрама\", \"+ужасы\"`) (optional)
    countriesName := []string{"Inner_example"} // []string | Поиск по странам (пример: `\"США\", \"Россия\", \"!Франция\" , \"+Великобритания\"`) (optional)
    ticketsOnSale := []string{"Inner_example"} // []string | Поиск по наличию билетов в продаже (пример: `true, false`) (optional)
    networksItemsName := []string{"Inner_example"} // []string | Поиск по сетям производства фильма (пример: `\"HBO\", \"Netflix\", \"!Amazon\"`) (optional)
    personsId := []string{"Inner_example"} // []string | Поиск по ID персон (пример: `666, 555, !666`) (optional)
    personsProfession := []string{"Inner_example"} // []string | Поиск по профессиям персон (пример: `\"актер\", \"режиссер\", \"!сценарист\"`) (optional)
    personsEnProfession := []string{"Inner_example"} // []string | Поиск по английским профессиям персон (пример: `\"actor\", \"director\", \"!writer\"`) (optional)
    feesWorld := []string{"Inner_example"} // []string | Поиск по сборам в мире (пример: `1000-6666666`) (optional)
    feesUsa := []string{"Inner_example"} // []string | Поиск по сборам в США (пример: `1000-6666666`) (optional)
    feesRussia := []string{"Inner_example"} // []string | Поиск по сборам в России (пример: `1000-6666666`) (optional)
    premiereWorld := []string{"Inner_example"} // []string | Поиск по дате премьеры в мире (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereUsa := []string{"Inner_example"} // []string | Поиск по дате премьеры в США (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereRussia := []string{"Inner_example"} // []string | Поиск по дате премьеры в России (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereDigital := []string{"Inner_example"} // []string | Поиск по дате премьеры в стриминговых сервисах (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereCinema := []string{"Inner_example"} // []string | Поиск по дате премьеры в кинотеатрах (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    premiereCountry := []string{"Inner_example"} // []string | Поиск по стране премьеры (пример: `\"США\", \"Россия\", \"!Франция\" , \"+Великобритания\"`) (optional)
    similarMoviesId := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk из списка похожих фильмов (пример: `666, 555, !666`) (optional)
    sequelsAndPrequelsId := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk из списка сиквелов и преквелов (пример: `666, 555, !666`) (optional)
    watchabilityItemsName := []string{"Inner_example"} // []string | Поиск по доуступным платформам для просмотра (пример: `\"ivi\", \"okko\", \"!megogo\"`) (optional)
    lists := []string{"Inner_example"} // []string | Поиск по коллекциям из KinoPoisk (пример: `\"top250\", \"top-100-indian-movies\", \"!top-100-movies\"`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerGetRandomMovieV14(context.Background()).NotNullFields(notNullFields).Id(id).ExternalIdImdb(externalIdImdb).ExternalIdTmdb(externalIdTmdb).ExternalIdKpHD(externalIdKpHD).Type_(type_).TypeNumber(typeNumber).IsSeries(isSeries).Status(status).Year(year).ReleaseYearsStart(releaseYearsStart).ReleaseYearsEnd(releaseYearsEnd).RatingKp(ratingKp).RatingImdb(ratingImdb).RatingTmdb(ratingTmdb).RatingMpaa(ratingMpaa).AgeRating(ageRating).VotesKp(votesKp).VotesImdb(votesImdb).VotesTmdb(votesTmdb).VotesFilmCritics(votesFilmCritics).VotesRussianFilmCritics(votesRussianFilmCritics).VotesAwait(votesAwait).BudgetValue(budgetValue).AudienceCount(audienceCount).MovieLength(movieLength).SeriesLength(seriesLength).TotalSeriesLength(totalSeriesLength).GenresName(genresName).CountriesName(countriesName).TicketsOnSale(ticketsOnSale).NetworksItemsName(networksItemsName).PersonsId(personsId).PersonsProfession(personsProfession).PersonsEnProfession(personsEnProfession).FeesWorld(feesWorld).FeesUsa(feesUsa).FeesRussia(feesRussia).PremiereWorld(premiereWorld).PremiereUsa(premiereUsa).PremiereRussia(premiereRussia).PremiereDigital(premiereDigital).PremiereCinema(premiereCinema).PremiereCountry(premiereCountry).SimilarMoviesId(similarMoviesId).SequelsAndPrequelsId(sequelsAndPrequelsId).WatchabilityItemsName(watchabilityItemsName).Lists(lists).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerGetRandomMovieV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerGetRandomMovieV14`: MovieDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerGetRandomMovieV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerGetRandomMovieV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **id** | **[]string** | Поиск по ID KinoPoisk (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **externalIdImdb** | **[]string** | Поиск по IMDB ID (пример: &#x60;\&quot;tt666\&quot;, \&quot;tt555\&quot;, \&quot;!tt666\&quot;&#x60;) | 
 **externalIdTmdb** | **[]float32** | Поиск по TMDB ID (пример: &#x60;666, 555, !666&#x60;) | 
 **externalIdKpHD** | **[]string** | Поиск по id KinoPoisk HD (пример: &#x60;\&quot;48e8d0acb0f62d8585101798eaeceec5\&quot;, \&quot;!48e8d0acb0f62d8585101798eaeceec5\&quot;&#x60;) | 
 **type_** | **[]string** | Поиск по типу фильма (пример: &#x60;\&quot;movie\&quot;, \&quot;tv-series\&quot;, \&quot;!anime\&quot;&#x60;) | 
 **typeNumber** | **[]string** | Поиск по номеру типа фильма (пример: &#x60;1, 2, !3&#x60;). Список типов: 1 (movie), 2 (tv-series), 3 (cartoon), 4 (anime), 5 (animated-series). | 
 **isSeries** | **[]string** | Поиск по индикатору сериала (пример: &#x60;true, false&#x60;) | 
 **status** | **[]string** | Поиск по статусу фильма (пример: &#x60;\&quot;announced\&quot;, \&quot;completed\&quot;, \&quot;!filming\&quot;&#x60;) | 
 **year** | **[]string** | Поиск по году (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **releaseYearsStart** | **[]string** | Поиск по года начала релиза (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **releaseYearsEnd** | **[]string** | Поиск по года окончания релиза (пример: &#x60;1874, 2050, !2020, 2020-2024&#x60;) | 
 **ratingKp** | **[]string** | Поиск по рейтингу Кинопоиск (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingImdb** | **[]string** | Поиск по рейтингу IMDB (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingTmdb** | **[]string** | Поиск по рейтингу TMDB (пример: &#x60;7, 10, 7.2-10&#x60;) | 
 **ratingMpaa** | **[]string** | Поиск по рейтингу MPAA (пример: &#x60;\&quot;G\&quot;, \&quot;NC-17\&quot;, \&quot;!R\&quot;&#x60;) | 
 **ageRating** | **[]string** | Поиск по возрастному рейтингу (пример: &#x60;12, !18, 12-18&#x60;) | 
 **votesKp** | **[]string** | Поиск по количеству голосов на KP (пример: &#x60;1000-6666666&#x60;) | 
 **votesImdb** | **[]string** | Поиск по количеству голосов на IMDB (пример: &#x60;1000-6666666&#x60;) | 
 **votesTmdb** | **[]string** | Поиск по количеству голосов на TMDB (пример: &#x60;1000-6666666&#x60;) | 
 **votesFilmCritics** | **[]string** | Поиск по количеству голосов кинокритиков (пример: &#x60;1000-6666666&#x60;) | 
 **votesRussianFilmCritics** | **[]string** | Поиск по количеству голосов кинокритиков из России (пример: &#x60;1000-6666666&#x60;) | 
 **votesAwait** | **[]string** | Поиск по количеству голосов ожидания на Кинопоиске (пример: &#x60;1000-6666666&#x60;) | 
 **budgetValue** | **[]string** | Поиск по бюджету фильма (пример: &#x60;1000-6666666&#x60;) | 
 **audienceCount** | **[]string** | Поиск по количеству аудитории (пример: &#x60;1000-6666666&#x60;) | 
 **movieLength** | **[]string** | Поиск по продолжительности фильма (пример: &#x60;100-120&#x60;) | 
 **seriesLength** | **[]string** | Поиск по всей продолжительности одной серии (пример: &#x60;20-60&#x60;) | 
 **totalSeriesLength** | **[]string** | Поиск по всей продолжительности сериала (пример: &#x60;100-120&#x60;) | 
 **genresName** | **[]string** | Поиск по жанрам (пример: &#x60;\&quot;драма\&quot;, \&quot;комедия\&quot;, \&quot;!мелодрама\&quot;, \&quot;+ужасы\&quot;&#x60;) | 
 **countriesName** | **[]string** | Поиск по странам (пример: &#x60;\&quot;США\&quot;, \&quot;Россия\&quot;, \&quot;!Франция\&quot; , \&quot;+Великобритания\&quot;&#x60;) | 
 **ticketsOnSale** | **[]string** | Поиск по наличию билетов в продаже (пример: &#x60;true, false&#x60;) | 
 **networksItemsName** | **[]string** | Поиск по сетям производства фильма (пример: &#x60;\&quot;HBO\&quot;, \&quot;Netflix\&quot;, \&quot;!Amazon\&quot;&#x60;) | 
 **personsId** | **[]string** | Поиск по ID персон (пример: &#x60;666, 555, !666&#x60;) | 
 **personsProfession** | **[]string** | Поиск по профессиям персон (пример: &#x60;\&quot;актер\&quot;, \&quot;режиссер\&quot;, \&quot;!сценарист\&quot;&#x60;) | 
 **personsEnProfession** | **[]string** | Поиск по английским профессиям персон (пример: &#x60;\&quot;actor\&quot;, \&quot;director\&quot;, \&quot;!writer\&quot;&#x60;) | 
 **feesWorld** | **[]string** | Поиск по сборам в мире (пример: &#x60;1000-6666666&#x60;) | 
 **feesUsa** | **[]string** | Поиск по сборам в США (пример: &#x60;1000-6666666&#x60;) | 
 **feesRussia** | **[]string** | Поиск по сборам в России (пример: &#x60;1000-6666666&#x60;) | 
 **premiereWorld** | **[]string** | Поиск по дате премьеры в мире (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereUsa** | **[]string** | Поиск по дате премьеры в США (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereRussia** | **[]string** | Поиск по дате премьеры в России (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereDigital** | **[]string** | Поиск по дате премьеры в стриминговых сервисах (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereCinema** | **[]string** | Поиск по дате премьеры в кинотеатрах (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **premiereCountry** | **[]string** | Поиск по стране премьеры (пример: &#x60;\&quot;США\&quot;, \&quot;Россия\&quot;, \&quot;!Франция\&quot; , \&quot;+Великобритания\&quot;&#x60;) | 
 **similarMoviesId** | **[]string** | Поиск по ID KinoPoisk из списка похожих фильмов (пример: &#x60;666, 555, !666&#x60;) | 
 **sequelsAndPrequelsId** | **[]string** | Поиск по ID KinoPoisk из списка сиквелов и преквелов (пример: &#x60;666, 555, !666&#x60;) | 
 **watchabilityItemsName** | **[]string** | Поиск по доуступным платформам для просмотра (пример: &#x60;\&quot;ivi\&quot;, \&quot;okko\&quot;, \&quot;!megogo\&quot;&#x60;) | 
 **lists** | **[]string** | Поиск по коллекциям из KinoPoisk (пример: &#x60;\&quot;top250\&quot;, \&quot;top-100-indian-movies\&quot;, \&quot;!top-100-movies\&quot;&#x60;) | 

### Return type

[**MovieDtoV14**](MovieDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieControllerSearchMovieV14

> SearchMovieResponseDtoV14 MovieControllerSearchMovieV14(ctx).Query(query).Page(page).Limit(limit).Execute()

Поиск фильмов по названию



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    query := "query_example" // string | Поисковый запрос
    page := float32(8.14) // float32 | Страница выборки (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MovieControllerSearchMovieV14(context.Background()).Query(query).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MovieControllerSearchMovieV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MovieControllerSearchMovieV14`: SearchMovieResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MovieControllerSearchMovieV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovieControllerSearchMovieV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Поисковый запрос | 
 **page** | **float32** | Страница выборки | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]

### Return type

[**SearchMovieResponseDtoV14**](SearchMovieResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonControllerFindManyAwardsV14

> PersonAwardDocsResponseDto PersonControllerFindManyAwardsV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).PersonId(personId).NominationTitle(nominationTitle).NominationAwardTitle(nominationAwardTitle).NominationAwardYear(nominationAwardYear).Winning(winning).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Награды актеров

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    personId := []string{"Inner_example"} // []string | Поиск по ID персоны (пример: `\"666\", \"555\", \"!666\"`) (optional)
    nominationTitle := []string{"Inner_example"} // []string | Поиск по номинациям (пример: `\"Оскар\", \"Золотой глобус\"`) (optional)
    nominationAwardTitle := []string{"Inner_example"} // []string | Поиск по наградам (пример: `\"Лучший фильм\", \"Лучший актер\"`) (optional)
    nominationAwardYear := []string{"Inner_example"} // []string | Поиск по году награды (пример: `\"2019\", \"2020\"`) (optional)
    winning := "winning_example" // string | Поиск по победам (пример: `\"true\", \"false\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PersonControllerFindManyAwardsV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).PersonId(personId).NominationTitle(nominationTitle).NominationAwardTitle(nominationAwardTitle).NominationAwardYear(nominationAwardYear).Winning(winning).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PersonControllerFindManyAwardsV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PersonControllerFindManyAwardsV14`: PersonAwardDocsResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PersonControllerFindManyAwardsV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPersonControllerFindManyAwardsV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **personId** | **[]string** | Поиск по ID персоны (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **nominationTitle** | **[]string** | Поиск по номинациям (пример: &#x60;\&quot;Оскар\&quot;, \&quot;Золотой глобус\&quot;&#x60;) | 
 **nominationAwardTitle** | **[]string** | Поиск по наградам (пример: &#x60;\&quot;Лучший фильм\&quot;, \&quot;Лучший актер\&quot;&#x60;) | 
 **nominationAwardYear** | **[]string** | Поиск по году награды (пример: &#x60;\&quot;2019\&quot;, \&quot;2020\&quot;&#x60;) | 
 **winning** | **string** | Поиск по победам (пример: &#x60;\&quot;true\&quot;, \&quot;false\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**PersonAwardDocsResponseDto**](PersonAwardDocsResponseDto.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonControllerFindManyV14

> PersonDocsResponseDtoV14 PersonControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Sex(sex).Growth(growth).Birthday(birthday).Death(death).Age(age).BirthPlaceValue(birthPlaceValue).DeathPlaceValue(deathPlaceValue).SpousesId(spousesId).SpousesDivorced(spousesDivorced).SpousesSex(spousesSex).CountAwards(countAwards).ProfessionValue(professionValue).MoviesRating(moviesRating).MoviesEnProfession(moviesEnProfession).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Универсальный поиск с фильтрами



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    id := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk (пример: `\"111\", \"222\", \"!666\"`) (optional)
    moviesId := []string{"Inner_example"} // []string | Поиск по ID фильма (пример: `\"666\", \"555\", \"!666\"`) (optional)
    sex := []string{"Sex_example"} // []string | Поиск по гендеру (пример: `Женский, Мужской`) (optional)
    growth := []string{"Inner_example"} // []string | Поиск по росту (пример: `170-180, 180`) (optional)
    birthday := []string{"Inner_example"} // []string | Поиск по дате рождения (пример: `01.01.2000-01.01.2001, 01.01.2000`) (optional)
    death := []string{"Inner_example"} // []string | Поиск по дате смерти (пример: `01.01.2000-01.01.2001, 01.01.2000`) (optional)
    age := []string{"Inner_example"} // []string | Поиск по возрасту (пример: `18-25, 25`) (optional)
    birthPlaceValue := []string{"Inner_example"} // []string | Поиск по месту рождения (пример: `Москва, Санкт-Петербург`) (optional)
    deathPlaceValue := []string{"Inner_example"} // []string | Поиск по месту смерти (пример: `Москва, Санкт-Петербург`) (optional)
    spousesId := []string{"Inner_example"} // []string | Поиск по ID супруги(супруга) (пример: `111, 222`) (optional)
    spousesDivorced := "spousesDivorced_example" // string | Поиск по статусу развода (пример: `true, false`) (optional)
    spousesSex := []string{"SpousesSex_example"} // []string | Поиск по гендеру супруги(супруга) (пример: `Женский, Мужской`) (optional)
    countAwards := []string{"Inner_example"} // []string | Поиск по количеству наград (пример: `1-10, 10`) (optional)
    professionValue := []string{"ProfessionValue_example"} // []string | Поиск по профессии (пример: `Актер, Режиссер`) (optional)
    moviesRating := []string{"Inner_example"} // []string | Поиск по рейтингу фильма (пример: `1-10, 10`) (optional)
    moviesEnProfession := []string{"MoviesEnProfession_example"} // []string | Поиск по профессии в фильмах на английском (пример: `actor, director`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PersonControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Sex(sex).Growth(growth).Birthday(birthday).Death(death).Age(age).BirthPlaceValue(birthPlaceValue).DeathPlaceValue(deathPlaceValue).SpousesId(spousesId).SpousesDivorced(spousesDivorced).SpousesSex(spousesSex).CountAwards(countAwards).ProfessionValue(professionValue).MoviesRating(moviesRating).MoviesEnProfession(moviesEnProfession).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PersonControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PersonControllerFindManyV14`: PersonDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PersonControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPersonControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **id** | **[]string** | Поиск по ID KinoPoisk (пример: &#x60;\&quot;111\&quot;, \&quot;222\&quot;, \&quot;!666\&quot;&#x60;) | 
 **moviesId** | **[]string** | Поиск по ID фильма (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **sex** | **[]string** | Поиск по гендеру (пример: &#x60;Женский, Мужской&#x60;) | 
 **growth** | **[]string** | Поиск по росту (пример: &#x60;170-180, 180&#x60;) | 
 **birthday** | **[]string** | Поиск по дате рождения (пример: &#x60;01.01.2000-01.01.2001, 01.01.2000&#x60;) | 
 **death** | **[]string** | Поиск по дате смерти (пример: &#x60;01.01.2000-01.01.2001, 01.01.2000&#x60;) | 
 **age** | **[]string** | Поиск по возрасту (пример: &#x60;18-25, 25&#x60;) | 
 **birthPlaceValue** | **[]string** | Поиск по месту рождения (пример: &#x60;Москва, Санкт-Петербург&#x60;) | 
 **deathPlaceValue** | **[]string** | Поиск по месту смерти (пример: &#x60;Москва, Санкт-Петербург&#x60;) | 
 **spousesId** | **[]string** | Поиск по ID супруги(супруга) (пример: &#x60;111, 222&#x60;) | 
 **spousesDivorced** | **string** | Поиск по статусу развода (пример: &#x60;true, false&#x60;) | 
 **spousesSex** | **[]string** | Поиск по гендеру супруги(супруга) (пример: &#x60;Женский, Мужской&#x60;) | 
 **countAwards** | **[]string** | Поиск по количеству наград (пример: &#x60;1-10, 10&#x60;) | 
 **professionValue** | **[]string** | Поиск по профессии (пример: &#x60;Актер, Режиссер&#x60;) | 
 **moviesRating** | **[]string** | Поиск по рейтингу фильма (пример: &#x60;1-10, 10&#x60;) | 
 **moviesEnProfession** | **[]string** | Поиск по профессии в фильмах на английском (пример: &#x60;actor, director&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**PersonDocsResponseDtoV14**](PersonDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonControllerFindOneV14

> Person PersonControllerFindOneV14(ctx, id).Execute()

Поиск по id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := float32(8.14) // float32 | ID из кинопоиска

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PersonControllerFindOneV14(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PersonControllerFindOneV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PersonControllerFindOneV14`: Person
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PersonControllerFindOneV14`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | ID из кинопоиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonControllerFindOneV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Person**](Person.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonControllerSearchPersonV14

> SearchPersonResponseDtoV14 PersonControllerSearchPersonV14(ctx).Query(query).Page(page).Limit(limit).Execute()

Поиск актеров, режиссеров, и т.д по имени



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    query := "query_example" // string | Поисковый запрос
    page := float32(8.14) // float32 | Страница выборки (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PersonControllerSearchPersonV14(context.Background()).Query(query).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PersonControllerSearchPersonV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PersonControllerSearchPersonV14`: SearchPersonResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PersonControllerSearchPersonV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPersonControllerSearchPersonV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Поисковый запрос | 
 **page** | **float32** | Страница выборки | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]

### Return type

[**SearchPersonResponseDtoV14**](SearchPersonResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewControllerFindManyV14

> ReviewDocsResponseDtoV14 ReviewControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MovieId(movieId).AuthorId(authorId).Author(author).Type_(type_).Date(date).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Универсальный поиск с фильтрами



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    id := []string{"Inner_example"} // []string | Поиск по ID отзыва (пример: `\"111\", \"222\", \"!666\"`) (optional)
    movieId := []string{"Inner_example"} // []string | Поиск по ID фильма (пример: `\"666\", \"555\", \"!666\"`) (optional)
    authorId := []string{"Inner_example"} // []string | Поиск отзывов по ID автора (пример: `\"666\", \"555\", \"!666\"`) (optional)
    author := []string{"Inner_example"} // []string | Поиск по имени автора отзыва (пример: `\"КиноПоиск\", \"!КиноПоиск\"`) (optional)
    type_ := []string{"Type_example"} // []string | Поиск по типу отзыва (пример: `\"!Негативный\", \"Нейтральный\", \"Позитивный\"`) (optional)
    date := []string{"Inner_example"} // []string | Поиск по дате создания отзыва (пример: `\"01.01.2021-01.01.2022\", \"01.01.2021\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ReviewControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MovieId(movieId).AuthorId(authorId).Author(author).Type_(type_).Date(date).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReviewControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReviewControllerFindManyV14`: ReviewDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReviewControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **id** | **[]string** | Поиск по ID отзыва (пример: &#x60;\&quot;111\&quot;, \&quot;222\&quot;, \&quot;!666\&quot;&#x60;) | 
 **movieId** | **[]string** | Поиск по ID фильма (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **authorId** | **[]string** | Поиск отзывов по ID автора (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **author** | **[]string** | Поиск по имени автора отзыва (пример: &#x60;\&quot;КиноПоиск\&quot;, \&quot;!КиноПоиск\&quot;&#x60;) | 
 **type_** | **[]string** | Поиск по типу отзыва (пример: &#x60;\&quot;!Негативный\&quot;, \&quot;Нейтральный\&quot;, \&quot;Позитивный\&quot;&#x60;) | 
 **date** | **[]string** | Поиск по дате создания отзыва (пример: &#x60;\&quot;01.01.2021-01.01.2022\&quot;, \&quot;01.01.2021\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**ReviewDocsResponseDtoV14**](ReviewDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeasonControllerFindManyV14

> SeasonDocsResponseDtoV14 SeasonControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).Number(number).EpisodesNumber(episodesNumber).AirDate(airDate).EpisodesAirDate(episodesAirDate).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Поиск сезонов

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    movieId := []string{"Inner_example"} // []string | Поиск по ID фильма (пример: `\"666\", \"555\", \"!666\"`) (optional)
    number := []string{"Inner_example"} // []string | Поиск по номеру сезона (пример: `\"1\", \"1-19\", \"!3\"`) (optional)
    episodesNumber := []string{"Inner_example"} // []string | Поиск по нормеру эпизода (пример: `\"1\", \"1-19\", \"!3\"`) (optional)
    airDate := []string{"Inner_example"} // []string | Поиск по дате выхода сезона (пример: `\"2020-01-01-2020-12-31\", \"2020-01-01\"`) (optional)
    episodesAirDate := []string{"Inner_example"} // []string | Поиск по дате выхода эпизода (пример: `\"2020-01-01-2020-12-31\", \"2020-01-01\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления в базе (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления в базу (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SeasonControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).MovieId(movieId).Number(number).EpisodesNumber(episodesNumber).AirDate(airDate).EpisodesAirDate(episodesAirDate).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SeasonControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeasonControllerFindManyV14`: SeasonDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SeasonControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeasonControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **movieId** | **[]string** | Поиск по ID фильма (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **number** | **[]string** | Поиск по номеру сезона (пример: &#x60;\&quot;1\&quot;, \&quot;1-19\&quot;, \&quot;!3\&quot;&#x60;) | 
 **episodesNumber** | **[]string** | Поиск по нормеру эпизода (пример: &#x60;\&quot;1\&quot;, \&quot;1-19\&quot;, \&quot;!3\&quot;&#x60;) | 
 **airDate** | **[]string** | Поиск по дате выхода сезона (пример: &#x60;\&quot;2020-01-01-2020-12-31\&quot;, \&quot;2020-01-01\&quot;&#x60;) | 
 **episodesAirDate** | **[]string** | Поиск по дате выхода эпизода (пример: &#x60;\&quot;2020-01-01-2020-12-31\&quot;, \&quot;2020-01-01\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления в базе (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления в базу (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**SeasonDocsResponseDtoV14**](SeasonDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StudioControllerFindManyV14

> StudioDocsResponseDtoV14 StudioControllerFindManyV14(ctx).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Type_(type_).SubType(subType).Title(title).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()

Поиск студий



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := float32(8.14) // float32 | Номер страницы (optional) (default to 1)
    limit := float32(8.14) // float32 | Количество элементов на странице (optional) (default to 10)
    selectFields := []string{"SelectFields_example"} // []string | Список полей требуемых в ответе из модели (optional)
    notNullFields := []string{"NotNullFields_example"} // []string | Список полей которые не должны быть null или пусты (optional)
    sortField := []string{"SortField_example"} // []string | Сортировка по полям из модели (optional)
    sortType := []string{"Inner_example"} // []string | Тип сортировки применительно к полям из sortField (пример: `\"1\", \"-1\"`) (optional)
    id := []string{"Inner_example"} // []string | Поиск по ID KinoPoisk (пример: `\"warnerbros\", \"222\", \"!666\"`) (optional)
    moviesId := []string{"Inner_example"} // []string | Поиск по ID фильма (пример: `\"666\", \"555\", \"!666\"`) (optional)
    type_ := []string{"Inner_example"} // []string | Поиск по типу студии (пример: `\"Производство\", \"Студия дубляжа\"`) (optional)
    subType := []string{"Inner_example"} // []string | Поиск по типу студии (пример: `\"company\", \"studio\"`) (optional)
    title := []string{"Inner_example"} // []string | Поиск по названию студии (пример: `\"Warner Bros.\", \"!Warner Bros.\"`) (optional)
    updatedAt := []string{"Inner_example"} // []string | Поиск по дате обновления (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)
    createdAt := []string{"Inner_example"} // []string | Поиск по дате добавления (пример: `01.01.2020, 01.01.2020-31.12.2020`) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StudioControllerFindManyV14(context.Background()).Page(page).Limit(limit).SelectFields(selectFields).NotNullFields(notNullFields).SortField(sortField).SortType(sortType).Id(id).MoviesId(moviesId).Type_(type_).SubType(subType).Title(title).UpdatedAt(updatedAt).CreatedAt(createdAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StudioControllerFindManyV14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StudioControllerFindManyV14`: StudioDocsResponseDtoV14
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StudioControllerFindManyV14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStudioControllerFindManyV14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Номер страницы | [default to 1]
 **limit** | **float32** | Количество элементов на странице | [default to 10]
 **selectFields** | **[]string** | Список полей требуемых в ответе из модели | 
 **notNullFields** | **[]string** | Список полей которые не должны быть null или пусты | 
 **sortField** | **[]string** | Сортировка по полям из модели | 
 **sortType** | **[]string** | Тип сортировки применительно к полям из sortField (пример: &#x60;\&quot;1\&quot;, \&quot;-1\&quot;&#x60;) | 
 **id** | **[]string** | Поиск по ID KinoPoisk (пример: &#x60;\&quot;warnerbros\&quot;, \&quot;222\&quot;, \&quot;!666\&quot;&#x60;) | 
 **moviesId** | **[]string** | Поиск по ID фильма (пример: &#x60;\&quot;666\&quot;, \&quot;555\&quot;, \&quot;!666\&quot;&#x60;) | 
 **type_** | **[]string** | Поиск по типу студии (пример: &#x60;\&quot;Производство\&quot;, \&quot;Студия дубляжа\&quot;&#x60;) | 
 **subType** | **[]string** | Поиск по типу студии (пример: &#x60;\&quot;company\&quot;, \&quot;studio\&quot;&#x60;) | 
 **title** | **[]string** | Поиск по названию студии (пример: &#x60;\&quot;Warner Bros.\&quot;, \&quot;!Warner Bros.\&quot;&#x60;) | 
 **updatedAt** | **[]string** | Поиск по дате обновления (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 
 **createdAt** | **[]string** | Поиск по дате добавления (пример: &#x60;01.01.2020, 01.01.2020-31.12.2020&#x60;) | 

### Return type

[**StudioDocsResponseDtoV14**](StudioDocsResponseDtoV14.md)

### Authorization

[X-API-KEY](../README.md#X-API-KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

