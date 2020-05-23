package main

import (
	"html/template"
)

var schema = "http"

var html = template.Must(template.New(schema).Parse(`
<html>
<head>
  <title>Grafana Webhook Listener</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:blue;">Under Construction...</h1>
</body>
</html>
`))

//BasicAction contains all the methods of the alertbody
type BasicAction interface {
	Save() string
	Alert() string
}

//GrafanaBody defines the properties of the alert body
type GrafanaBody struct {
	Title       string      `json:"title"`
	RuleName    string      `json:"ruleName"`
	State       string      `json:"state"`
	Message     string      `json:"message"`
	RuleID      int         `json:"ruleId"`
	RuleURL     string      `json:"ruleUrl"`
	ImageURL    string      `json:"immageUrl"`
	EvalMatches []alertItem `json:"evalMatches"`
}

//AlertItem defines the properties of the evalmatches in the alert body
type alertItem struct {
	Metric string   `json:"metric"`
	Tags   alertTag `json:"tags"`
	Value  int      `json:"value"`
}

//AlertTag describes the properties for the tags in the aleritem
type alertTag struct {
	Name, Value string
}

//Save method describes how the alertbody saves data
func (b *GrafanaBody) Save() (int, error) {
	err := new(error)
	status := 0
	//Initialization for DB connection
	return status, *err
}

//Alert method describes how the alertbody handles alerting
func (b *GrafanaBody) Alert() string {
	//method for sending alert
	str := "Alerting " + b.Title + "!!!"
	return str
}

//Action function lists how the alerbody methods are being used
func Action(a ...BasicAction) {
	for _, c := range a {
		c.Save()
		c.Alert()
	}

}

//Mock function generates a mock data of type alertbody
func Mock() (b GrafanaBody) {
	Bod := GrafanaBody{}
	Bod.Title = "body title"
	Bod.Message = "body Messsage"
	return Bod
}
