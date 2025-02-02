// swift-tools-version:5.3
import PackageDescription

let package = Package(
    name: "MongoLiteWrapper",
    products: [
        .library(
            name: "MongoLiteWrapper",
            targets: ["MongoLite"]),
    ],
    dependencies: [],
    targets: [
        .target(
            name: "MongoLite",
            dependencies: []),
        .testTarget(
            name: "MongoLiteTests",
            dependencies: ["MongoLite"]),
    ]
)