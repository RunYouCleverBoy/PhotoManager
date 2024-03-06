import org.gradle.api.Plugin
import org.gradle.api.Project
import java.io.File



class ImportModelsFromGo : Plugin<Project> {
    override fun apply(target: Project) {
        target.task("importModelsFromGo") {
            doLast {
                val goModels = File("../../Server/models")
                val kotlinModels =
                    File("app/src/main/java/com/photomanager/photomanager/main/home/api")
                goModels.listFiles()?.forEach { goFile ->
                    if (!goFile.name.endsWith(".go")) return@forEach

                    val outputFileName = goFile.name.largeCamel().replace(".go", ".kt")
                    val kotlinFile = File(kotlinModels, outputFileName)
                    if (kotlinFile.exists()) {
                        kotlinFile.delete()
                    }

                    val lines = mutableListOf(
                        "package com.photomanager.photomanager.main.home.api",
                        "",
                        "import kotlinx.serialization.Serializable",
                        "import kotlinx.serialization.SerialName",
                        "",
                    )
                    lines.addAll(compile(goFile))
                    kotlinFile.bufferedWriter().use { out ->
                        lines.forEach {
                            out.write(it)
                            out.newLine()
                        }
                    }
                }
            }
        }
    }

    private fun String.largeCamel(): String {
        return mapIndexed { i, c -> if (i == 0) c.uppercase() else c }.joinToString("")
    }

    private fun String.smallCamel(): String {
        return mapIndexed { i, c -> if (i == 0) c.lowercase() else c }.joinToString("")
    }

    sealed class LineType(val kotlinLine: String) {
        object SerializableAnnotation : LineType("@Serializable")
        sealed class ClassDeclaration(val name: String, text: String) : LineType(text) {
            class Data(name: String) : ClassDeclaration(name, "data class $name(")
            class Regular(name: String) : ClassDeclaration(name, "class $name(")
        }

        class Field(line: String) : LineType(line)
        class Comment(line: String) : LineType(line)
        object ClassDeclarationEnd : LineType(")")
    }

    private fun compile(goFile: File): List<String> {
        val result: MutableList<String> = mutableListOf()
        val group: MutableList<LineType> = mutableListOf()
        goFile.readLines().forEach {
            val line = it.trim()
            val fieldRegex = Regex("^(\\w+)\\s+\\*?([\\w.\\[\\]]+)\\s+`json:\"(\\w+)(\" |,|\"`).*")
            when {
                line.matches(Regex("^type\\s+\\w+\\s+struct\\s+\\{")) -> {
                    if (group.isNotEmpty()) {
                        group.clear()
                    }
                    val name = line.split(Regex("\\s+"))[1]
                    group.add(LineType.SerializableAnnotation)
                    group.add(LineType.ClassDeclaration.Data(name.largeCamel()))
                }

                fieldRegex.matches(line) -> {
                    val v = fieldRegex.find(line)?.groupValues!!
                    val name = v[1].smallCamel()
                    val type = v[2]
                    val jsonName = v[3]
                    val addedLine = when (type) {
                        "int" -> "@SerialName(\"$jsonName\") val $name: Int = 0"
                        "int64" -> "@SerialName(\"$jsonName\") val $name: Long = 0L"
                        "string" -> "@SerialName(\"$jsonName\") val $name: String = \"\""
                        "bool" -> "@SerialName(\"$jsonName\") val $name: Boolean = false"
                        "float64" -> "@SerialName(\"$jsonName\") val $name: Double = 0.0"
                        "float32" -> "@SerialName(\"$jsonName\") val $name: Float = 0f"
                        "[]byte" -> "@SerialName(\"$jsonName\") val $name: ByteArray = ByteArray(0)"
                        "[]string" -> "@SerialName(\"$jsonName\") val $name: List<String> = emptyList()"
                        "[]primitive.ObjectID" -> "@SerialName(\"$jsonName\") val $name: List<String> = emptyList()"
                        "time.Time" -> "@SerialName(\"$jsonName\") val $name: Long = 0L"
                        "primitive.ObjectID" -> "@SerialName(\"$jsonName\") val $name: String"
                        else -> {
                            if (type.startsWith("[]")) {
                                val innerType = type.removePrefix("[]")
                                "@SerialName(\"$jsonName\") val $name: List<$innerType> = emptyList()"
                            } else {
                                "@SerialName(\"$jsonName\") val $name: $type"
                            }
                        }
                    }
                    group.add(LineType.Field(addedLine))
                }

                line.matches(Regex("^}")) -> {
                    group.add(LineType.ClassDeclarationEnd)
                    val groupIsOk = group.verify()
                    if (groupIsOk) {
                        val reassembledLines = group.synthesize()
                        result.addAll(reassembledLines)
                    } else {
                        println("Group is not ok. File ${goFile.name} group fails")
                    }
                    group.clear()
                }

                else -> group.add(LineType.Comment("// $line"))
            }
        }
        return result
    }

    private fun List<LineType>.verify(): Boolean {
        data class Rule(val error: String, val condition: (List<LineType>) -> Boolean)

        infix fun String.onFail(condition: (List<LineType>) -> Boolean) = Rule(this, condition)
        val rules = listOf(
            "Must have a Serialization annotation" onFail { group -> group.getOrNull(0) is LineType.SerializableAnnotation },
            "Must have a class declaration" onFail { group -> group.getOrNull(1) is LineType.ClassDeclaration },
            "Must have a closing line" onFail { group -> group.lastOrNull() is LineType.ClassDeclarationEnd },
            "Inside the declaration there must be only fields" onFail { group ->
                group.none { item -> item is LineType.Field } || (2 until lastIndex).run { all { i -> group[i] is LineType.Field } }
            })

        val withoutComments = filterNot { it is LineType.Comment }
        return rules.all { rule ->
            rule.condition(withoutComments).also {
                if (!it) {
                    println("Error: ${rule.error}")
                    println(withoutComments.joinToString("\n") { x -> x.kotlinLine })
                }
            }
        }
    }

    private fun List<LineType>.synthesize(): List<String> {
        var fieldCount = count { it is LineType.Field }
        return map { lineType ->
            when (lineType) {
                is LineType.SerializableAnnotation -> lineType.kotlinLine
                is LineType.ClassDeclaration.Data -> {
                    if (fieldCount == 0) {
                        LineType.ClassDeclaration.Regular(lineType.name).kotlinLine
                    } else {
                        LineType.ClassDeclaration.Data(lineType.name).kotlinLine
                    }
                }

                is LineType.Field -> {
                    var comma = ","
                    fieldCount--
                    if (fieldCount < 0) {
                        comma = ""
                    }
                    "   ${lineType.kotlinLine}$comma"
                }

                is LineType.Comment -> lineType.kotlinLine
                is LineType.ClassDeclarationEnd -> lineType.kotlinLine
                else -> {
                    throw IllegalStateException("Unknown line type")
                }
            }
        }
    }
}

