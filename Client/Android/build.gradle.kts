// Top-level build file where you can add configuration options common to all sub-projects/modules.

val ktorVersion by extra { "2.3.7" }
val navVersion by extra { "2.7.7" }

plugins {
    id("com.android.application") version "8.2.2" apply false
    id("org.jetbrains.kotlin.android") version "1.9.20" apply false
    id("org.jetbrains.kotlin.jvm") version "1.9.20" apply false
    id("com.google.devtools.ksp") version "1.9.20-1.0.14" apply false
    id("com.google.dagger.hilt.android") version "2.49" apply false
}

