import XCTest
@testable import MongoLiteWrapper

final class MongoLiteTests: XCTestCase {

    override func setUp() {
        super.setUp()
        // Clean up the data file before each test
        let fileManager = FileManager.default
        let dataFilePath = "mongo.local"
        if fileManager.fileExists(atPath: dataFilePath) {
            try? fileManager.removeItem(atPath: dataFilePath)
        }
    }

    func testInsertDocument() {
        let json = "{\"name\": \"test\", \"value\": 123}"
        let result = MongoLite.insertDocument(json: json)
        XCTAssertEqual(result, "InsertDocument: document successfully inserted")

        let searchResult = MongoLite.searchDocuments(query: "{\"name\": \"test\"}")
        let expected = "[{\"name\":\"test\",\"value\":123}]"
        XCTAssertEqual(searchResult, expected)
    }

    func testSearchDocuments() {
        MongoLite.insertDocument(json: "{\"name\": \"test1\", \"value\": 123}")
        MongoLite.insertDocument(json: "{\"name\": \"test2\", \"value\": 456}")

        let searchResult = MongoLite.searchDocuments(query: "{\"value\": 123}")
        let expected = "[{\"name\":\"test1\",\"value\":123}]"
        XCTAssertEqual(searchResult, expected)
    }

    func testDeleteDocuments() {
        MongoLite.insertDocument(json: "{\"name\": \"test1\", \"value\": 123}")
        MongoLite.insertDocument(json: "{\"name\": \"test2\", \"value\": 456}")

        let deleteResult = MongoLite.deleteDocuments(query: "{\"value\": 123}")
        XCTAssertEqual(deleteResult, "DeleteDocument: 1 documents deleted")

        let searchResult = MongoLite.searchDocuments(query: "{\"value\": 123}")
        let expected = "[]"
        XCTAssertEqual(searchResult, expected)
    }
}