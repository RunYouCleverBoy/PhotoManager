package com.photomanager.photomanager.initialzers

import android.content.Context
import androidx.startup.Initializer
import timber.log.Timber

class TimberInitializer : Initializer<Unit> {
    private class DebugTree : Timber.DebugTree() {
        override fun createStackElementTag(element: StackTraceElement): String? {
            return "${super.createStackElementTag(element)}:${element.lineNumber}"
        }
    }
    override fun create(context: Context) {
        // Timber initialization
        Timber.plant(DebugTree())
    }

    override fun dependencies(): List<Class<out Initializer<*>>> = emptyList()
}
