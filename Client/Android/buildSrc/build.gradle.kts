plugins {
    `java-gradle-plugin`
    `kotlin-dsl`
}

gradlePlugin {
    plugins {
        create("com.gradle.plugin.import-models-from-go") {
            id = "com.plugins.import-models-from-go"
            implementationClass = "ImportModelsFromGo"
        }
    }
}

