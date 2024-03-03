package com.photomanager.photomanager.main.home.db

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey
import androidx.room.TypeConverters
import java.util.Date

@Entity(tableName = "Footage")
@TypeConverters(DateConverter::class)
data class FootageEntity(
    @PrimaryKey(autoGenerate = true) val id: Int = 0,
    @ColumnInfo(name="uri") val uri: String,
    @ColumnInfo(name="width") val width: Int,
    @ColumnInfo(name="height") val height: Int,
    @ColumnInfo(name="orientation") val orientation: String?,
    @ColumnInfo(name="caption") val caption: String,
    @ColumnInfo(name="date") val date: Date,
    @ColumnInfo(name="lat") val lat: Double?,
    @ColumnInfo(name="lon") val lon: Double?,
    @ColumnInfo(name="camera") val camera: String,
    @ColumnInfo(name="focalLength") val focalLength: String?,
)


@Entity(tableName = "Collection")
@TypeConverters(DateConverter::class)
data class CollectionEntity(
    @PrimaryKey(autoGenerate = true) val id: Int = 0,
    @ColumnInfo(name="uri") val uri: String,
    @ColumnInfo(name="width") val width: Int,
    @ColumnInfo(name="height") val height: Int,
    @ColumnInfo(name="orientation") val orientation: String?,
    @ColumnInfo(name="caption") val caption: String,
    @ColumnInfo(name="date") val date: Date,
    @ColumnInfo(name="lat") val lat: Double,
    @ColumnInfo(name="lon") val lon: Double,
    @ColumnInfo(name="camera") val camera: String,
    @ColumnInfo(name="focalLength") val focalLength: String,
)
