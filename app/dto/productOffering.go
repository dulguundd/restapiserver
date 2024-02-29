package dto

type ProductOffering struct {
	Id                          string `json:"_id"`
	BaseType                    string `json:"@baseType"`
	SchemaLocation              string `json:"@schemaLocation"`
	Type                        string `json:"@type"`
	Agreement                   string `json:"agreement"`
	Attachment                  string `json:"attachment"`
	BundledProductOffering      string `json:"bundledProductOffering"`
	Category                    string `json:"category"`
	Channel                     string `json:"channel"`
	Description                 string `json:"description"`
	Href                        string `json:"href"`
	Id1                         string `json:"id"`
	IsBundle                    string `json:"isBundle"`
	IsSellable                  string `json:"isSellable"`
	LastUpdate                  string `json:"lastUpdate"`
	LifecycleStatus             string `json:"lifecycleStatus"`
	MarketSegment               string `json:"marketSegment"`
	Name                        string `json:"name"`
	Place                       string `json:"place"`
	ProdSpecCharValueUse        string `json:"prodSpecCharValueUse"`
	ProductOfferingPrice        string `json:"productOfferingPrice"`
	ProductOfferingRelationship string `json:"productOfferingRelationship"`
	ProductOfferingTerm         string `json:"productOfferingTerm"`
	ProductSpecification        string `json:"productSpecification"`
	ResourceCandidate           string `json:"resourceCandidate"`
	ServiceCandidate            string `json:"serviceCandidate"`
	ServiceLevelAgreement       string `json:"serviceLevelAgreement"`
	StatusReason                string `json:"statusReason"`
	ValidFor                    string `json:"validFor"`
	Version                     string `json:"version"`
}
