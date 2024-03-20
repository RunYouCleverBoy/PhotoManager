package com.photomanager.photomanager.main.home.db

import androidx.room.Database
import androidx.room.RoomDatabase
import com.photomanager.photomanager.main.home.db.dao.PhotoDao


//abstract class PhotoDatabase : RoomDatabase() {
//    abstract fun footageDao(): FootageDao
//    abstract fun collectionDao(): CollectionDao
//}

@Database(entities = [Photo::class, PhotosAndAlbumCross::class, PhotoAndTags::class], version = 1)
abstract class PhotoDatabase : RoomDatabase() {
    abstract fun photoDao(): PhotoDao
}
