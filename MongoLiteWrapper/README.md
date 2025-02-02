# MongoLiteWrapper

MongoLiteWrapper is a Swift wrapper around the MongoLite library, providing a simple interface for interacting with MongoDB-like document storage in Swift applications.

## Features

- **Insert Documents**: Easily insert JSON documents into the storage.
- **Search Documents**: Query documents based on key-value pairs.
- **Delete Documents**: Remove documents that match specified criteria.

## Installation

To use MongoLiteWrapper in your project, add it as a dependency in your `Package.swift` file:

```swift
dependencies: [
    .package(url: "https://github.com/yourfavDev/mongoLite.git", from: "1.0.0")
]
```

## Usage

### Importing the Library

To use the MongoLite wrapper in your Swift files, import the module:

```swift
import MongoLiteWrapper
```

### Inserting a Document

You can insert a document by calling the `insertDocument` method:

```swift
let result = MongoLite.insertDocument(json: "{\"name\": \"test\", \"value\": 123}")
print(result)
```

### Searching for Documents

To search for documents, use the `searchDocuments` method:

```swift
let results = MongoLite.searchDocuments(query: "{\"value\": 123}")
print(results)
```

### Deleting Documents

To delete documents that match a query, call the `deleteDocuments` method:

```swift
let deleteResult = MongoLite.deleteDocuments(query: "{\"value\": 123}")
print(deleteResult)
```

## Testing

Unit tests are provided in the `Tests` directory. You can run the tests using the following command:

```bash
swift test
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.