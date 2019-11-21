//the idea is taken from https://gist.github.com/bosgood/d9687f5d64c0f4038bc19c478d51fa64
// but that code throw exception and does not compile
// I just refined that and fix the problem and convert it to Int version
// same can be done for other nullable types 

// LeadItem struct
type LeadItem struct {
	CustomerId sql.NullInt64
}


//the names of fields should 100% match  struct field name
var Lead = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Lead",
		Fields: graphql.Fields{
			"CustomerId": &graphql.Field{
					Type: graphql.Int,
					Resolve:  resolveNullableInt,
			},
		},
	},
)


func resolveNullableInt(p graphql.ResolveParams) (interface{}, error) {
	v := reflect.ValueOf(p.Source)

	if !v.IsValid() {
		return nil, fmt.Errorf("Invalid elem: %v", p.Source)
	}

	fieldName := p.Info.FieldName
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		return nil, fmt.Errorf("Missing field: %v", fieldName)
	}

	nullString := f.Interface().(sql.NullInt64)
	if nullString.Valid {
		return nullString.Int64, nil
	}

	return nil, nil
}
