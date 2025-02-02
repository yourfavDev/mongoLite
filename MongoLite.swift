import Foundation

@objc class MongoLite: NSObject {
    @objc static func insertDocument(json: String) -> String {
        let cjson = strdup(json)
        defer { free(cjson) }
        let result = InsertDocument(cjson)
        return String(cString: result!)
    }

    @objc static func searchDocuments(query: String) -> String {
        let cquery = strdup(query)
        defer { free(cquery) }
        let result = SearchDocument(cquery)
        return String(cString: result!)
    }

    @objc static func deleteDocuments(query: String) -> String {
        let cquery = strdup(query)
        defer { free(cquery) }
        let result = DeleteDocument(cquery)
        return String(cString: result!)
    }
}
