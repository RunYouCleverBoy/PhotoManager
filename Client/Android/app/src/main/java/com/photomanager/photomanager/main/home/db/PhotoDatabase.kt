package com.photomanager.photomanager.main.home.db

import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.TypeConverters
import com.photomanager.photomanager.main.home.db.dao.PhotoDao

@Database(
    entities = [Photo::class, PhotoAlbum::class, PhotosAndAlbumCross::class, PhotoAndTags::class],
    exportSchema = false,
    version = 1
)
@TypeConverters(CompositeTypeConverters::class)
abstract class PhotoDatabase : RoomDatabase() {
    abstract fun photoDao(): PhotoDao
}
