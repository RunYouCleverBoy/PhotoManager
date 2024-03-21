package com.photomanager.photomanager.main.home.db

import android.content.Context
import dagger.hilt.android.qualifiers.ApplicationContext
import javax.inject.Inject

class DatabaseHolderImpl @Inject constructor(@ApplicationContext private val appContext: Context) :
    DatabaseHolder {
    override var database: PhotoDatabase = createDatabase()
        get() = field.takeIf { it.isOpen } ?: createDatabase().also { field = it }

    override fun closeDatabase() {
        database.close()
    }

    private fun createDatabase(): PhotoDatabase {
        return androidx.room.Room.databaseBuilder(
            appContext,
            PhotoDatabase::class.java,
            DATABASE_NAME
        ).build()
    }

    companion object {
        const val DATABASE_NAME: String = "photos_database"
    }
}