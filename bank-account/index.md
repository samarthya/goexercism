# API's called
```
PUT samarthya
```
Allows to create an index
```json
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "samarthya"
}
```
## Criteria
Index name should meet the following criteria
- Lowercase only
- Cannot include special character `\`, `/`,`*`,`?` etc.
- Cannot include `:`
- Cannot use `.` or `..`
- Max 255 characters

