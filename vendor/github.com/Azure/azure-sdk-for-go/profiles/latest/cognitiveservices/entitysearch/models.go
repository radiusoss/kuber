// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package entitysearch

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/entitysearch"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type EntitiesClient = original.EntitiesClient
type AnswerType = original.AnswerType

const (
	AnswerTypeEntities AnswerType = original.AnswerTypeEntities
	AnswerTypePlaces   AnswerType = original.AnswerTypePlaces
)

type EntityQueryScenario = original.EntityQueryScenario

const (
	Disambiguation                   EntityQueryScenario = original.Disambiguation
	DominantEntity                   EntityQueryScenario = original.DominantEntity
	DominantEntityWithDisambiguation EntityQueryScenario = original.DominantEntityWithDisambiguation
	List                             EntityQueryScenario = original.List
	ListWithPivot                    EntityQueryScenario = original.ListWithPivot
)

type EntityScenario = original.EntityScenario

const (
	EntityScenarioDisambiguationItem EntityScenario = original.EntityScenarioDisambiguationItem
	EntityScenarioDominantEntity     EntityScenario = original.EntityScenarioDominantEntity
	EntityScenarioListItem           EntityScenario = original.EntityScenarioListItem
)

type EntityType = original.EntityType

const (
	EntityTypeActor               EntityType = original.EntityTypeActor
	EntityTypeAnimal              EntityType = original.EntityTypeAnimal
	EntityTypeArtist              EntityType = original.EntityTypeArtist
	EntityTypeAttorney            EntityType = original.EntityTypeAttorney
	EntityTypeAttraction          EntityType = original.EntityTypeAttraction
	EntityTypeBook                EntityType = original.EntityTypeBook
	EntityTypeCar                 EntityType = original.EntityTypeCar
	EntityTypeCity                EntityType = original.EntityTypeCity
	EntityTypeCollegeOrUniversity EntityType = original.EntityTypeCollegeOrUniversity
	EntityTypeComposition         EntityType = original.EntityTypeComposition
	EntityTypeContinent           EntityType = original.EntityTypeContinent
	EntityTypeCountry             EntityType = original.EntityTypeCountry
	EntityTypeDrug                EntityType = original.EntityTypeDrug
	EntityTypeEvent               EntityType = original.EntityTypeEvent
	EntityTypeFood                EntityType = original.EntityTypeFood
	EntityTypeGeneric             EntityType = original.EntityTypeGeneric
	EntityTypeHotel               EntityType = original.EntityTypeHotel
	EntityTypeHouse               EntityType = original.EntityTypeHouse
	EntityTypeLocalBusiness       EntityType = original.EntityTypeLocalBusiness
	EntityTypeLocality            EntityType = original.EntityTypeLocality
	EntityTypeMedia               EntityType = original.EntityTypeMedia
	EntityTypeMinorRegion         EntityType = original.EntityTypeMinorRegion
	EntityTypeMovie               EntityType = original.EntityTypeMovie
	EntityTypeMusicAlbum          EntityType = original.EntityTypeMusicAlbum
	EntityTypeMusicGroup          EntityType = original.EntityTypeMusicGroup
	EntityTypeMusicRecording      EntityType = original.EntityTypeMusicRecording
	EntityTypeNeighborhood        EntityType = original.EntityTypeNeighborhood
	EntityTypeOrganization        EntityType = original.EntityTypeOrganization
	EntityTypeOther               EntityType = original.EntityTypeOther
	EntityTypePerson              EntityType = original.EntityTypePerson
	EntityTypePlace               EntityType = original.EntityTypePlace
	EntityTypePointOfInterest     EntityType = original.EntityTypePointOfInterest
	EntityTypePostalCode          EntityType = original.EntityTypePostalCode
	EntityTypeProduct             EntityType = original.EntityTypeProduct
	EntityTypeRadioStation        EntityType = original.EntityTypeRadioStation
	EntityTypeRegion              EntityType = original.EntityTypeRegion
	EntityTypeRestaurant          EntityType = original.EntityTypeRestaurant
	EntityTypeSchool              EntityType = original.EntityTypeSchool
	EntityTypeSpeciality          EntityType = original.EntityTypeSpeciality
	EntityTypeSportsTeam          EntityType = original.EntityTypeSportsTeam
	EntityTypeState               EntityType = original.EntityTypeState
	EntityTypeStreetAddress       EntityType = original.EntityTypeStreetAddress
	EntityTypeSubRegion           EntityType = original.EntityTypeSubRegion
	EntityTypeTelevisionSeason    EntityType = original.EntityTypeTelevisionSeason
	EntityTypeTelevisionShow      EntityType = original.EntityTypeTelevisionShow
	EntityTypeTheaterPlay         EntityType = original.EntityTypeTheaterPlay
	EntityTypeTouristAttraction   EntityType = original.EntityTypeTouristAttraction
	EntityTypeTravel              EntityType = original.EntityTypeTravel
	EntityTypeVideoGame           EntityType = original.EntityTypeVideoGame
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type ResponseFormat = original.ResponseFormat

const (
	JSON   ResponseFormat = original.JSON
	JSONLd ResponseFormat = original.JSONLd
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type Type = original.Type

const (
	TypeContractualRulesAttribution        Type = original.TypeContractualRulesAttribution
	TypeContractualRulesContractualRule    Type = original.TypeContractualRulesContractualRule
	TypeContractualRulesLicenseAttribution Type = original.TypeContractualRulesLicenseAttribution
	TypeContractualRulesLinkAttribution    Type = original.TypeContractualRulesLinkAttribution
	TypeContractualRulesMediaAttribution   Type = original.TypeContractualRulesMediaAttribution
	TypeContractualRulesTextAttribution    Type = original.TypeContractualRulesTextAttribution
)

type TypeBasicResponseBase = original.TypeBasicResponseBase

const (
	TypeAirport               TypeBasicResponseBase = original.TypeAirport
	TypeAnswer                TypeBasicResponseBase = original.TypeAnswer
	TypeCivicStructure        TypeBasicResponseBase = original.TypeCivicStructure
	TypeCreativeWork          TypeBasicResponseBase = original.TypeCreativeWork
	TypeEntertainmentBusiness TypeBasicResponseBase = original.TypeEntertainmentBusiness
	TypeEntities              TypeBasicResponseBase = original.TypeEntities
	TypeErrorResponse         TypeBasicResponseBase = original.TypeErrorResponse
	TypeFoodEstablishment     TypeBasicResponseBase = original.TypeFoodEstablishment
	TypeHotel                 TypeBasicResponseBase = original.TypeHotel
	TypeIdentifiable          TypeBasicResponseBase = original.TypeIdentifiable
	TypeImageObject           TypeBasicResponseBase = original.TypeImageObject
	TypeIntangible            TypeBasicResponseBase = original.TypeIntangible
	TypeLicense               TypeBasicResponseBase = original.TypeLicense
	TypeLocalBusiness         TypeBasicResponseBase = original.TypeLocalBusiness
	TypeLodgingBusiness       TypeBasicResponseBase = original.TypeLodgingBusiness
	TypeMediaObject           TypeBasicResponseBase = original.TypeMediaObject
	TypeMovieTheater          TypeBasicResponseBase = original.TypeMovieTheater
	TypeOrganization          TypeBasicResponseBase = original.TypeOrganization
	TypePlace                 TypeBasicResponseBase = original.TypePlace
	TypePlaces                TypeBasicResponseBase = original.TypePlaces
	TypePostalAddress         TypeBasicResponseBase = original.TypePostalAddress
	TypeResponse              TypeBasicResponseBase = original.TypeResponse
	TypeResponseBase          TypeBasicResponseBase = original.TypeResponseBase
	TypeRestaurant            TypeBasicResponseBase = original.TypeRestaurant
	TypeSearchResponse        TypeBasicResponseBase = original.TypeSearchResponse
	TypeSearchResultsAnswer   TypeBasicResponseBase = original.TypeSearchResultsAnswer
	TypeStructuredValue       TypeBasicResponseBase = original.TypeStructuredValue
	TypeThing                 TypeBasicResponseBase = original.TypeThing
	TypeTouristAttraction     TypeBasicResponseBase = original.TypeTouristAttraction
)

type Airport = original.Airport
type BasicAnswer = original.BasicAnswer
type Answer = original.Answer
type BasicCivicStructure = original.BasicCivicStructure
type CivicStructure = original.CivicStructure
type BasicContractualRulesAttribution = original.BasicContractualRulesAttribution
type ContractualRulesAttribution = original.ContractualRulesAttribution
type BasicContractualRulesContractualRule = original.BasicContractualRulesContractualRule
type ContractualRulesContractualRule = original.ContractualRulesContractualRule
type ContractualRulesLicenseAttribution = original.ContractualRulesLicenseAttribution
type ContractualRulesLinkAttribution = original.ContractualRulesLinkAttribution
type ContractualRulesMediaAttribution = original.ContractualRulesMediaAttribution
type ContractualRulesTextAttribution = original.ContractualRulesTextAttribution
type BasicCreativeWork = original.BasicCreativeWork
type CreativeWork = original.CreativeWork
type BasicEntertainmentBusiness = original.BasicEntertainmentBusiness
type EntertainmentBusiness = original.EntertainmentBusiness
type Entities = original.Entities
type EntitiesEntityPresentationInfo = original.EntitiesEntityPresentationInfo
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type BasicFoodEstablishment = original.BasicFoodEstablishment
type FoodEstablishment = original.FoodEstablishment
type Hotel = original.Hotel
type BasicIdentifiable = original.BasicIdentifiable
type Identifiable = original.Identifiable
type ImageObject = original.ImageObject
type BasicIntangible = original.BasicIntangible
type Intangible = original.Intangible
type License = original.License
type BasicLocalBusiness = original.BasicLocalBusiness
type LocalBusiness = original.LocalBusiness
type BasicLodgingBusiness = original.BasicLodgingBusiness
type LodgingBusiness = original.LodgingBusiness
type BasicMediaObject = original.BasicMediaObject
type MediaObject = original.MediaObject
type MovieTheater = original.MovieTheater
type Organization = original.Organization
type BasicPlace = original.BasicPlace
type Place = original.Place
type Places = original.Places
type PostalAddress = original.PostalAddress
type QueryContext = original.QueryContext
type BasicResponse = original.BasicResponse
type Response = original.Response
type BasicResponseBase = original.BasicResponseBase
type ResponseBase = original.ResponseBase
type Restaurant = original.Restaurant
type SearchResponse = original.SearchResponse
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type SearchResultsAnswer = original.SearchResultsAnswer
type BasicStructuredValue = original.BasicStructuredValue
type StructuredValue = original.StructuredValue
type BasicThing = original.BasicThing
type Thing = original.Thing
type TouristAttraction = original.TouristAttraction

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewEntitiesClient() EntitiesClient {
	return original.NewEntitiesClient()
}
func NewEntitiesClientWithBaseURI(baseURI string) EntitiesClient {
	return original.NewEntitiesClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
