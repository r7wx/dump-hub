package esapi

const entryMapping = `
{
  "settings": {
    "number_of_shards": 5,
    "number_of_replicas": 0,
    "refresh_interval" : "30s",
    "codec": "best_compression" 
  },
  "mappings": {
    "dynamic_templates": [
      {
        "all_text": {
          "match_mapping_type": "string",
          "mapping": {
            "copy_to": "_all",
            "type": "text"
          }
        }
      }
    ],
    "properties": {
      "_all": {
        "type": "text"
      }
    }
  }
}
`

const statusMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "properties": {
      "date": {"type": "keyword" }, 
      "filename": { "type": "keyword" }, 
      "status": { "type": "integer" }
    }
  }
}
`
