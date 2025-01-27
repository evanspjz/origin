// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package consumption

import original "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-10-01/consumption"

type AggregatedCostClient = original.AggregatedCostClient
type BalancesClient = original.BalancesClient
type BudgetsClient = original.BudgetsClient
type ChargesClient = original.ChargesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ForecastsClient = original.ForecastsClient
type MarketplacesClient = original.MarketplacesClient
type BillingFrequency = original.BillingFrequency

const (
	Month   BillingFrequency = original.Month
	Quarter BillingFrequency = original.Quarter
	Year    BillingFrequency = original.Year
)

type Bound = original.Bound

const (
	Lower Bound = original.Lower
	Upper Bound = original.Upper
)

type CategoryType = original.CategoryType

const (
	Cost  CategoryType = original.Cost
	Usage CategoryType = original.Usage
)

type ChargeType = original.ChargeType

const (
	ChargeTypeActual   ChargeType = original.ChargeTypeActual
	ChargeTypeForecast ChargeType = original.ChargeTypeForecast
)

type Datagrain = original.Datagrain

const (
	DailyGrain   Datagrain = original.DailyGrain
	MonthlyGrain Datagrain = original.MonthlyGrain
)

type Grain = original.Grain

const (
	Daily   Grain = original.Daily
	Monthly Grain = original.Monthly
	Yearly  Grain = original.Yearly
)

type OperatorType = original.OperatorType

const (
	EqualTo              OperatorType = original.EqualTo
	GreaterThan          OperatorType = original.GreaterThan
	GreaterThanOrEqualTo OperatorType = original.GreaterThanOrEqualTo
)

type TimeGrainType = original.TimeGrainType

const (
	TimeGrainTypeAnnually  TimeGrainType = original.TimeGrainTypeAnnually
	TimeGrainTypeMonthly   TimeGrainType = original.TimeGrainTypeMonthly
	TimeGrainTypeQuarterly TimeGrainType = original.TimeGrainTypeQuarterly
)

type Balance = original.Balance
type BalanceProperties = original.BalanceProperties
type BalancePropertiesAdjustmentDetailsItem = original.BalancePropertiesAdjustmentDetailsItem
type BalancePropertiesNewPurchasesDetailsItem = original.BalancePropertiesNewPurchasesDetailsItem
type Budget = original.Budget
type BudgetProperties = original.BudgetProperties
type BudgetsListResult = original.BudgetsListResult
type BudgetsListResultIterator = original.BudgetsListResultIterator
type BudgetsListResultPage = original.BudgetsListResultPage
type BudgetTimePeriod = original.BudgetTimePeriod
type ChargesListResult = original.ChargesListResult
type ChargeSummary = original.ChargeSummary
type ChargeSummaryProperties = original.ChargeSummaryProperties
type CurrentSpend = original.CurrentSpend
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Filters = original.Filters
type Forecast = original.Forecast
type ForecastProperties = original.ForecastProperties
type ForecastPropertiesConfidenceLevelsItem = original.ForecastPropertiesConfidenceLevelsItem
type ForecastsListResult = original.ForecastsListResult
type ManagementGroupAggregatedCostProperties = original.ManagementGroupAggregatedCostProperties
type ManagementGroupAggregatedCostResult = original.ManagementGroupAggregatedCostResult
type Marketplace = original.Marketplace
type MarketplaceProperties = original.MarketplaceProperties
type MarketplacesListResult = original.MarketplacesListResult
type MarketplacesListResultIterator = original.MarketplacesListResultIterator
type MarketplacesListResultPage = original.MarketplacesListResultPage
type MeterDetails = original.MeterDetails
type Notification = original.Notification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PriceSheetModel = original.PriceSheetModel
type PriceSheetProperties = original.PriceSheetProperties
type PriceSheetResult = original.PriceSheetResult
type ProxyResource = original.ProxyResource
type ReservationDetail = original.ReservationDetail
type ReservationDetailProperties = original.ReservationDetailProperties
type ReservationDetailsListResult = original.ReservationDetailsListResult
type ReservationDetailsListResultIterator = original.ReservationDetailsListResultIterator
type ReservationDetailsListResultPage = original.ReservationDetailsListResultPage
type ReservationRecommendation = original.ReservationRecommendation
type ReservationRecommendationProperties = original.ReservationRecommendationProperties
type ReservationRecommendationsListResult = original.ReservationRecommendationsListResult
type ReservationRecommendationsListResultIterator = original.ReservationRecommendationsListResultIterator
type ReservationRecommendationsListResultPage = original.ReservationRecommendationsListResultPage
type ReservationSummariesListResult = original.ReservationSummariesListResult
type ReservationSummariesListResultIterator = original.ReservationSummariesListResultIterator
type ReservationSummariesListResultPage = original.ReservationSummariesListResultPage
type ReservationSummary = original.ReservationSummary
type ReservationSummaryProperties = original.ReservationSummaryProperties
type Resource = original.Resource
type ResourceAttributes = original.ResourceAttributes
type Tag = original.Tag
type TagProperties = original.TagProperties
type TagsResult = original.TagsResult
type UsageDetail = original.UsageDetail
type UsageDetailProperties = original.UsageDetailProperties
type UsageDetailsListResult = original.UsageDetailsListResult
type UsageDetailsListResultIterator = original.UsageDetailsListResultIterator
type UsageDetailsListResultPage = original.UsageDetailsListResultPage
type OperationsClient = original.OperationsClient
type PriceSheetClient = original.PriceSheetClient
type ReservationRecommendationsClient = original.ReservationRecommendationsClient
type ReservationsDetailsClient = original.ReservationsDetailsClient
type ReservationsSummariesClient = original.ReservationsSummariesClient
type TagsClient = original.TagsClient
type UsageDetailsClient = original.UsageDetailsClient

func NewAggregatedCostClient(subscriptionID string) AggregatedCostClient {
	return original.NewAggregatedCostClient(subscriptionID)
}
func NewAggregatedCostClientWithBaseURI(baseURI string, subscriptionID string) AggregatedCostClient {
	return original.NewAggregatedCostClientWithBaseURI(baseURI, subscriptionID)
}
func NewBalancesClient(subscriptionID string) BalancesClient {
	return original.NewBalancesClient(subscriptionID)
}
func NewBalancesClientWithBaseURI(baseURI string, subscriptionID string) BalancesClient {
	return original.NewBalancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBudgetsClient(subscriptionID string) BudgetsClient {
	return original.NewBudgetsClient(subscriptionID)
}
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string) BudgetsClient {
	return original.NewBudgetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewChargesClient(subscriptionID string) ChargesClient {
	return original.NewChargesClient(subscriptionID)
}
func NewChargesClientWithBaseURI(baseURI string, subscriptionID string) ChargesClient {
	return original.NewChargesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewForecastsClient(subscriptionID string) ForecastsClient {
	return original.NewForecastsClient(subscriptionID)
}
func NewForecastsClientWithBaseURI(baseURI string, subscriptionID string) ForecastsClient {
	return original.NewForecastsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMarketplacesClient(subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClient(subscriptionID)
}
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleBillingFrequencyValues() []BillingFrequency {
	return original.PossibleBillingFrequencyValues()
}
func PossibleBoundValues() []Bound {
	return original.PossibleBoundValues()
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleChargeTypeValues() []ChargeType {
	return original.PossibleChargeTypeValues()
}
func PossibleDatagrainValues() []Datagrain {
	return original.PossibleDatagrainValues()
}
func PossibleGrainValues() []Grain {
	return original.PossibleGrainValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return original.PossibleTimeGrainTypeValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClient(subscriptionID)
}
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationRecommendationsClient(subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClient(subscriptionID)
}
func NewReservationRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClient(subscriptionID)
}
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClient(subscriptionID)
}
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClient(subscriptionID)
}
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
