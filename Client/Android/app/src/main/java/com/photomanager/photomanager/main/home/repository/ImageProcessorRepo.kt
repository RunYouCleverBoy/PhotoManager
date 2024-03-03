package com.photomanager.photomanager.main.home.repository

import android.net.Uri
import com.photomanager.photomanager.main.home.model.ImageDescriptor

interface ImageProcessorRepo {
    fun processImage(uri: Uri): ImageDescriptor
}