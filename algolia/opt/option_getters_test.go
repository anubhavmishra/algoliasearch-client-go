// Code generated by go generate. DO NOT EDIT.

package opt

import "testing"

func TestOptionGetters(t *testing.T) {
	var searchableAttributes *SearchableAttributesOption = nil
	searchableAttributes.Get()

	var attributesForFaceting *AttributesForFacetingOption = nil
	attributesForFaceting.Get()

	var unretrievableAttributes *UnretrievableAttributesOption = nil
	unretrievableAttributes.Get()

	var attributesToRetrieve *AttributesToRetrieveOption = nil
	attributesToRetrieve.Get()

	var restrictSearchableAttributes *RestrictSearchableAttributesOption = nil
	restrictSearchableAttributes.Get()

	var ranking *RankingOption = nil
	ranking.Get()

	var customRanking *CustomRankingOption = nil
	customRanking.Get()

	var replicas *ReplicasOption = nil
	replicas.Get()

	var primary *PrimaryOption = nil
	primary.Get()

	var filters *FiltersOption = nil
	filters.Get()

	var facetFilters *FacetFiltersOption = nil
	facetFilters.Get()

	var optionalFilters *OptionalFiltersOption = nil
	optionalFilters.Get()

	var numericFilters *NumericFiltersOption = nil
	numericFilters.Get()

	var tagFilters *TagFiltersOption = nil
	tagFilters.Get()

	var sumOrFiltersScores *SumOrFiltersScoresOption = nil
	sumOrFiltersScores.Get()

	var facets *FacetsOption = nil
	facets.Get()

	var maxValuesPerFacet *MaxValuesPerFacetOption = nil
	maxValuesPerFacet.Get()

	var facetingAfterDistinct *FacetingAfterDistinctOption = nil
	facetingAfterDistinct.Get()

	var sortFacetValuesBy *SortFacetValuesByOption = nil
	sortFacetValuesBy.Get()

	var attributesToHighlight *AttributesToHighlightOption = nil
	attributesToHighlight.Get()

	var attributesToSnippet *AttributesToSnippetOption = nil
	attributesToSnippet.Get()

	var highlightPreTag *HighlightPreTagOption = nil
	highlightPreTag.Get()

	var highlightPostTag *HighlightPostTagOption = nil
	highlightPostTag.Get()

	var snippetEllipsisText *SnippetEllipsisTextOption = nil
	snippetEllipsisText.Get()

	var restrictHighlightAndSnippetArrays *RestrictHighlightAndSnippetArraysOption = nil
	restrictHighlightAndSnippetArrays.Get()

	var page *PageOption = nil
	page.Get()

	var hitsPerPage *HitsPerPageOption = nil
	hitsPerPage.Get()

	var offset *OffsetOption = nil
	offset.Get()

	var length *LengthOption = nil
	length.Get()

	var paginationLimitedTo *PaginationLimitedToOption = nil
	paginationLimitedTo.Get()

	var minWordSizefor1Typo *MinWordSizefor1TypoOption = nil
	minWordSizefor1Typo.Get()

	var minWordSizefor2Typos *MinWordSizefor2TyposOption = nil
	minWordSizefor2Typos.Get()

	var typoTolerance *TypoToleranceOption = nil
	typoTolerance.Get()

	var allowTyposOnNumericTokens *AllowTyposOnNumericTokensOption = nil
	allowTyposOnNumericTokens.Get()

	var disableTypoToleranceOnAttributes *DisableTypoToleranceOnAttributesOption = nil
	disableTypoToleranceOnAttributes.Get()

	var disableTypoToleranceOnWords *DisableTypoToleranceOnWordsOption = nil
	disableTypoToleranceOnWords.Get()

	var separatorsToIndex *SeparatorsToIndexOption = nil
	separatorsToIndex.Get()

	var aroundLatLng *AroundLatLngOption = nil
	aroundLatLng.Get()

	var aroundLatLngViaIP *AroundLatLngViaIPOption = nil
	aroundLatLngViaIP.Get()

	var aroundRadius *AroundRadiusOption = nil
	aroundRadius.Get()

	var aroundPrecision *AroundPrecisionOption = nil
	aroundPrecision.Get()

	var minimumAroundRadius *MinimumAroundRadiusOption = nil
	minimumAroundRadius.Get()

	var insideBoundingBox *InsideBoundingBoxOption = nil
	insideBoundingBox.Get()

	var insidePolygon *InsidePolygonOption = nil
	insidePolygon.Get()

	var ignorePlurals *IgnorePluralsOption = nil
	ignorePlurals.Get()

	var removeStopWords *RemoveStopWordsOption = nil
	removeStopWords.Get()

	var camelCaseAttributes *CamelCaseAttributesOption = nil
	camelCaseAttributes.Get()

	var decompoundedAttributes *DecompoundedAttributesOption = nil
	decompoundedAttributes.Get()

	var keepDiacriticsOnCharacters *KeepDiacriticsOnCharactersOption = nil
	keepDiacriticsOnCharacters.Get()

	var queryLanguages *QueryLanguagesOption = nil
	queryLanguages.Get()

	var queryType *QueryTypeOption = nil
	queryType.Get()

	var removeWordsIfNoResults *RemoveWordsIfNoResultsOption = nil
	removeWordsIfNoResults.Get()

	var advancedSyntax *AdvancedSyntaxOption = nil
	advancedSyntax.Get()

	var optionalWords *OptionalWordsOption = nil
	optionalWords.Get()

	var disablePrefixOnAttributes *DisablePrefixOnAttributesOption = nil
	disablePrefixOnAttributes.Get()

	var disableExactOnAttributes *DisableExactOnAttributesOption = nil
	disableExactOnAttributes.Get()

	var exactOnSingleWordQuery *ExactOnSingleWordQueryOption = nil
	exactOnSingleWordQuery.Get()

	var alternativesAsExact *AlternativesAsExactOption = nil
	alternativesAsExact.Get()

	var advancedSyntaxFeatures *AdvancedSyntaxFeaturesOption = nil
	advancedSyntaxFeatures.Get()

	var enableRules *EnableRulesOption = nil
	enableRules.Get()

	var ruleContexts *RuleContextsOption = nil
	ruleContexts.Get()

	var enabled *EnabledOption = nil
	enabled.Get()

	var enablePersonalization *EnablePersonalizationOption = nil
	enablePersonalization.Get()

	var personalizationImpact *PersonalizationImpactOption = nil
	personalizationImpact.Get()

	var userToken *UserTokenOption = nil
	userToken.Get()

	var numericAttributesForFiltering *NumericAttributesForFilteringOption = nil
	numericAttributesForFiltering.Get()

	var allowCompressionOfIntegerArray *AllowCompressionOfIntegerArrayOption = nil
	allowCompressionOfIntegerArray.Get()

	var attributeForDistinct *AttributeForDistinctOption = nil
	attributeForDistinct.Get()

	var distinct *DistinctOption = nil
	distinct.Get()

	var getRankingInfo *GetRankingInfoOption = nil
	getRankingInfo.Get()

	var clickAnalytics *ClickAnalyticsOption = nil
	clickAnalytics.Get()

	var analytics *AnalyticsOption = nil
	analytics.Get()

	var analyticsTags *AnalyticsTagsOption = nil
	analyticsTags.Get()

	var synonyms *SynonymsOption = nil
	synonyms.Get()

	var replaceSynonymsInHighlight *ReplaceSynonymsInHighlightOption = nil
	replaceSynonymsInHighlight.Get()

	var minProximity *MinProximityOption = nil
	minProximity.Get()

	var responseFields *ResponseFieldsOption = nil
	responseFields.Get()

	var maxFacetHits *MaxFacetHitsOption = nil
	maxFacetHits.Get()

	var percentileComputation *PercentileComputationOption = nil
	percentileComputation.Get()

	var explain *ExplainOption = nil
	explain.Get()

	var advanced *AdvancedOption = nil
	advanced.Get()

	var restrictIndices *RestrictIndicesOption = nil
	restrictIndices.Get()

	var restrictSources *RestrictSourcesOption = nil
	restrictSources.Get()

	var validUntil *ValidUntilOption = nil
	validUntil.Get()

	var referers *ReferersOption = nil
	referers.Get()

	var query *QueryOption = nil
	query.Get()

	var autoGenerateObjectIDIfNotExist *AutoGenerateObjectIDIfNotExistOption = nil
	autoGenerateObjectIDIfNotExist.Get()

	var clearExistingRules *ClearExistingRulesOption = nil
	clearExistingRules.Get()

	var replaceExistingSynonyms *ReplaceExistingSynonymsOption = nil
	replaceExistingSynonyms.Get()

	var typeVar *TypeOption = nil
	typeVar.Get()

	var createIfNotExists *CreateIfNotExistsOption = nil
	createIfNotExists.Get()

	var forwardToReplicas *ForwardToReplicasOption = nil
	forwardToReplicas.Get()

	var anchoring *AnchoringOption = nil
	anchoring.Get()

	var extraHeaders *ExtraHeadersOption = nil
	extraHeaders.Get()

	var extraURLParams *ExtraURLParamsOption = nil
	extraURLParams.Get()

	var scopes *ScopesOption = nil
	scopes.Get()

	var clusterName *ClusterNameOption = nil
	clusterName.Get()

	var indexName *IndexNameOption = nil
	indexName.Get()

	var limit *LimitOption = nil
	limit.Get()

	var safe *SafeOption = nil
	safe.Get()

}
