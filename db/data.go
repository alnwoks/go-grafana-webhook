package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//AlertDocument models document for grafana messages collection
type AlertDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	RuleName    string             `bson:"ruleName,omitempty"`
	State       string             `bson:"state,omitempty"`
	Message     string             `bson:"message,omitempty"`
	RuleID      int                `bson:"ruleId,omitempty"`
	RuleURL     string             `bson:"ruleUrl,omitempty"`
	ImageURL    string             `bson:"immageUrl,omitempty"`
	EvalMatches []alertItem        `bson:"evalMatches,omitempty"`
}

//AlertItem defines the properties of the evalmatches in the alert body
type alertItem struct {
	Metric string   `bson:"metric"`
	Tags   alertTag `bson:"tags"`
	Value  int      `bson:"value"`
}

//AlertTag describes the properties for the tags in the aleritem
type alertTag struct {
	Name, Value string
}
