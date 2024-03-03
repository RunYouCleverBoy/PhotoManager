package com.photomanager.photomanager.main.home.db

import androidx.room.Database
import androidx.room.RoomDatabase


@Database(entities = [FootageEntity::class, CollectionEntity::class], version = 1)
abstract class PhotoDatabase : RoomDatabase() {
    abstract fun footageDao(): FootageDao
    abstract fun collectionDao(): CollectionDao
}

