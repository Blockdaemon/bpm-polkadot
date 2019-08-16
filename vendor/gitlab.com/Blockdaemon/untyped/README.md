# Description

untyped is a Go library to deal with nested maps and list as untyped data. It helps when manipulating unstructured JSON-like data.

## Why?

I was writing code like this:

	config["genesisBlock"].(map[string]interface{})["engine"].(map[string]interface{})["authorityRound"].(map[string]interface{})["params"].(map[string]interface{})["validators"].(map[string]interface{})["list"] = validators

which is super brittle (it panics if a type assertion fails) and ugly.

## How?

	const testDataJSON = `
	{
		"level1": {
			"list1": ["x", "y", "z"],
		},
		"list0": ["a", "b", [1, 2, 3], {"level4": "foo"}]
	}`

	var testData interface{}
	json.Unmarshal([]byte(testDataJSON), &testData)

	level1, err := GetValue(testData, "level1")
	list1, err := GetValue(testData, "level1.list1")
	secondElement, err := GetValue(testData, "level1.list1.[1]")
	level4, err := GetValue(testData, "list0.[3].level4")

or setting data:

	err = SetValue(testData, "list0.[3]", "abcd")
	err = SetValue(testData, "level1.list1", "abcd")

## Thanks

Thanks to Vladimir Nguyen for the initial implementation on which this library is based!
