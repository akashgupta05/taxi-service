# taxi-service
Input json:

```
[
    {
        "taxi": {
            "id": "8ab86c18-3fae-4804-bfd9-c3d6e8f66260",
            "type": "auto",
            "location": {
                "x": 42,
                "y": 37
            }
        }
    },
    {
        "taxi": {
            "id": "f26e890b-df8e-422e-a39c-7762aa0bac38",
            "type": "limousine",
            "location": {
                "x": 22,
                "y": 48
            }
        }
    },
    {
        "taxi": {
            "id": "ed0e23ef-6c2b-430c-9b90-cd4f1ff74c88",
            "type": "sedan",
            "location": {
                "x": 76,
                "y": 27
            }
        }
    },
    {
        "taxi": {
            "id": "690de6bc-163c-4345-bf6f-25dd0c58e864",
            "type": "sedan",
            "location": {
                "x": 32,
                "y": 63
            }
        }
    },
    {
        "taxi": {
            "id": "c0033410-981c-428a-954a-35dec05ef1d2",
            "type": "auto",
            "location": {
                "x": 8,
                "y": 13
            }
        }
    },
    {
        "taxi_request": {
            "customer_id": "b1286c18-3fae-4804-bfd9-c3d6e8f66260",
            "preferred_type": "auto",
            "secondary_type": "sedan",
            "location": {
                "x": 15,
                "y": 7
            },
            "destination": {
                "x": 66,
                "y": 21
            }
        }
    },
    {
        "taxi_request": {
            "customer_id": "34oe23ef-6c2b-430c-9b90-cd4f1ff74c88",
            "preferred_type": "limousine",
            "secondary_type": "auto",
            "location": {
                "x": 98,
                "y": 46
            },
            "destination": {
                "x": 55,
                "y": 89
            }
        }
    },
    {
        "trip_completion": {
            "customer_id": "34oe23ef-6c2b-430c-9b90-cd4f1ff74c88"
        }
    },
    {
        "taxi_request": {
            "customer_id": "34oe23ef-6c2b-430c-9b90-cd4f1ff74c89",
            "preferred_type": "limousine",
            "secondary_type": "auto",
            "location": {
                "x": 98,
                "y": 46
            },
            "destination": {
                "x": 55,
                "y": 89
            }
        }
    }
]
```

Output json:
```
[
	{
		"id": "c0033410-981c-428a-954a-35dec05ef1d2",
		"customer_id": "b1286c18-3fae-4804-bfd9-c3d6e8f66260"
	},
	{
		"id": "f26e890b-df8e-422e-a39c-7762aa0bac38",
		"customer_id": "34oe23ef-6c2b-430c-9b90-cd4f1ff74c88"
	},
	{
		"id": "f26e890b-df8e-422e-a39c-7762aa0bac38",
		"customer_id": "34oe23ef-6c2b-430c-9b90-cd4f1ff74c89"
	}
]
```


Steps to run the service:

`1. go build`

`2. ./taxi-service`