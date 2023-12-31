import json

# dictionary in python dumps into json format
superHeroes = {
  "squadName": "Super hero squad",
  "homeTown": "Metro City",
  "formed": 2016,
  "secretBase": "Super tower",
  "active": True,
  "members": [
    {
      "name": "Molecule Man",
      "age": 29,
      "secretIdentity": "Dan Jukes",
      "powers": [
        "Radiation resistance",
        "Turning tiny",
        "Radiation blast"
      ]
    },
    {
      "name": "Madame Uppercut",
      "age": 39,
      "secretIdentity": "Jane Wilson",
      "powers": [
        "Million tonne punch",
        "Damage resistance",
        "Superhuman reflexes"
      ]
    },
    {
      "name": "Eternal Flame",
      "age": 1000000,
      "secretIdentity": "Unknown",
      "powers": [
        "Immortality",
        "Heat Immunity",
        "Inferno",
        "Teleportation",
        "Interdimensional travel"
      ]
    }
  ]
}

print(superHeroes['homeTown'])
print(superHeroes['active'])
print(superHeroes['members'][1]['powers'][2])

with open('./Lecture_06_HTTP_1_1_Source_Code_v3/lec-06-prg-04-json-example/lec-06-prg-04-json-example2.json', 'w') as json_out_file:
    json.dump(superHeroes, json_out_file, indent='\t')
