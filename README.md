# go-tagger

> ⚠️ go-tagger is currently in early development. APIs and features will likely change in the full release.  **Use in production at your own risk**.

A simple tagging library that uses a set of defined rules to return a list of textual tags for any given struct. 

It's currently used in the wild for an e-commerce order processing pipeline. It tags thousands of orders to be used downstream for fraud detection, prioritisation, efficient dispatch, filtering and more.

- [Usage](#usage)
- [Documentation](#docs)

## Usage

Using go-tagger is simple. Start by creating your rules file, then see the code example below on how you'd get started. 

```json
[
    {
        "field": "Email",
        "rules": [
            {
                "type": "contains",
                "value": "@gmail.com",
                "tag": "gmail"
            }
        ]
    },
    {
        "field": "Items",
        "rules": [
            {
                "type": "count",
                "op": "gteq",
                "value": 2,
                "tag": "multiple-items"
            }
        ]
    }
]
```

```go
order := Order{
    Email: "user1@gmail.com",
    Items: []Item{
        {
            Title: "10 Colors Refrigerator Magnets",
        },
        {
            Title: "Mens Jumper Classic Sweater with V-Neck and Long Sleeve",
        },
    },
}
  
t := NewTagger()
t.LoadRules("resources/rules.json")
tags := t.Tag(order)

fmt.Println(tags) 
// Output [gmail, multiple-items]
```

## Docs
See the following documentation for more advanced go-tagger usage. Have a question? Please use [GitHub Discussions](https://github.com/loxhill/go-tagger). 

- [Nested Struct Fields](#nested-struct-fields)


### Nested Struct Fields
If you have a nested struct like below, you rule groups need to specify which specific field you want to parse. Just specifying `Contact` in the below example would not work as it is a struct.

```go
order := Order{
    Contact: {
        Email: "user1@gmail.com",
    }
}
```

In your rule group, for the field attribute, use a period to specify nested struct fields like so (`Contact.Email`).

```json
{
    "field": "Contact.Email",
    "rules": [...]
}
```

