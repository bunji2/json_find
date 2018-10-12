# json_find

## Usage: json_find json_file pattern

## Sample:

```
% cat file.json
[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
]

% json_find file.json #0.Name
Platypus

% json_find file.json #1.Order
Dasyuromorphia

% json_find file.json #2.Name
(not found)
```
