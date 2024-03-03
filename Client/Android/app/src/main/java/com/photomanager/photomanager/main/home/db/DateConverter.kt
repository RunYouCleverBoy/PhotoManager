package com.photomanager.photomanager.main.home.db

import androidx.room.TypeConverter
import java.util.Date

class DateConverter {
    @TypeConverter
    fun toDate(date: Long): Date {
        return Date(date)
    }

    @TypeConverter
    fun toLong(date: Date): Long {
        return date.time
    }
}