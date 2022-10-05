# go-tagger
A simple tagging library that can use a set of defined rules to return a list of applicable textual tags for any given struct. 

It's currently used in the wild for an e-commerce order processing pipeline. It tags thousands of orders to be used downstream for fraud detection, prioritisation, efficient dispatch, filtering and more.

## Usage

Here is a basic snippet of Go code and a sample rule file. This will return the `disposable-email` and `multiple-items` tags.

```go
  order := Order{
    Email: "grant@email.com",
    Address: OrderAddress{
        Postcode: "PE20 1PH",
    },
  }
  
  tagger := NewTagger()
  tagger.LoadRules("resources/rules.json")
  tags := tagger.Tag(order)
  fmt.Println(tags)
```

```json
{
  "Email": [
    {
      "type": "contains",
      "value": "@kvhrr.com",
      "tag": "disposable-email"
    }
  ],
  "Items": [
    {
      "type": "count",
      "op": "gteq",
      "value": 2,
      "tag": "multiple-items"
    }
  ],
  "Value": [
    {
      "type": "value",
      "op": "gteq",
      "value": 1000,
      "tag": "high-value"
    }
  ]
}
```

## Todo

- Use reflection to decide how to process the value in a rule  
  https://pkg.go.dev/reflect
- Load rules from json
- Allow users to store rules in a DB and load them into Tagger rather than load from json file.

## Ideas

- Regex in rules