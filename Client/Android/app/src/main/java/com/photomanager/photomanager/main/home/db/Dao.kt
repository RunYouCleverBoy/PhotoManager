package com.photomanager.photomanager.main.home.db

import androidx.room.Dao
import androidx.room.Delete
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import androidx.room.Transaction
import androidx.room.TypeConverters
import java.util.Date

@Dao
@TypeConverters(DateConverter::class)
interface FootageDao {
    @Query("SELECT * FROM Footage LIMIT :size OFFSET :startIndex")
    suspend fun getAllFootage(startIndex: Int, size: Int): List<FootageEntity>

    @Query("SELECT * FROM Footage WHERE id IN (:ids)")
    suspend fun getFootageByIds(ids: List<String>): List<FootageEntity>

    @Query("SELECT * FROM Footage WHERE " +
            "(:afterDate IS NULL OR date >= :afterDate) AND " +
            "(:beforeDate IS NULL OR date <= :beforeDate) AND " +
            "(:captionIncludes IS NULL OR caption LIKE '%'||:captionIncludes||'%') " +
            "LIMIT :size OFFSET :startIndex")
    suspend fun getFootageBy(
        afterDate: Date?,
        beforeDate: Date?,
        captionIncludes: String?,
        startIndex: Int,
        size: Int
    ): List<FootageEntity>

    @Query("SELECT COUNT() FROM Footage WHERE " +
            "(:afterDate IS NULL OR date >= :afterDate) AND " +
            "(:beforeDate IS NULL OR date <= :beforeDate) AND " +
            "(:captionIncludes IS NULL OR caption LIKE '%'||:captionIncludes||'%') ")
    suspend fun countFootageBy(
        afterDate: Date?,
        beforeDate: Date?,
        captionIncludes: String?
    ): Int

    @Delete
    suspend fun deleteFootage(footage: FootageEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertFootage(footage: List<FootageEntity>)
}

@Dao
@TypeConverters(DateConverter::class)
interface CollectionDao {
    @Query("SELECT * FROM Collection LIMIT :size OFFSET :startIndex")
    suspend fun getAllCollections(startIndex: Int, size: Int): List<CollectionEntity>

    @Query("SELECT * FROM Collection WHERE id = :id")
    suspend fun getCollectionById(id: Int): CollectionEntity

    @Transaction
    @Query("SELECT * FROM Collection WHERE " +
            "(:afterDate IS NULL OR date >= :afterDate) AND " +
            "(:beforeDate IS NULL OR date <= :beforeDate) AND " +
            "(:captionIncludes IS NULL OR caption LIKE '%'||:captionIncludes||'%') " +
            "LIMIT :size OFFSET :startIndex")
    suspend fun getCollectionBy(
        afterDate: Date?,
        beforeDate: Date?,
        captionIncludes: String?,
        startIndex: Int,
        size: Int
    ): List<CollectionEntity>

    @Query("SELECT COUNT() FROM Collection WHERE " +
            "(:afterDate IS NULL OR date >= :afterDate) AND " +
            "(:beforeDate IS NULL OR date <= :beforeDate) AND " +
            "(:captionIncludes IS NULL OR caption LIKE '%'||:captionIncludes||'%') ")
    suspend fun countCollectionBy(
        afterDate: Date?,
        beforeDate: Date?,
        captionIncludes: String?,
    ): Int

    @Delete
    suspend fun deleteCollection(collection: CollectionEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertCollection(collection: List<CollectionEntity>)
}