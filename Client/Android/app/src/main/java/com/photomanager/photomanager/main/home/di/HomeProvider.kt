package com.photomanager.photomanager.main.home.di

import com.photomanager.photomanager.main.home.api.PhotosApi
import com.photomanager.photomanager.main.home.api.PhotosApiImpl
import com.photomanager.photomanager.main.home.db.DatabaseHolder
import com.photomanager.photomanager.main.home.db.DatabaseHolderImpl
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepo
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepoImpl
import com.photomanager.photomanager.main.home.repository.ImagesRepo
import com.photomanager.photomanager.main.home.repository.ImagesRepoImpl
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ViewModelComponent

@Module
@InstallIn(ViewModelComponent::class)
abstract class HomeProvider {
    @Binds
    abstract fun provideWorkImagesRepo(photoRepo: ImagesRepoImpl): ImagesRepo

    @Binds
    abstract fun provideCollectionApi(collectionApi: PhotosApiImpl): PhotosApi

    @Binds
    abstract fun provideDatabase(database: DatabaseHolderImpl): DatabaseHolder

    @Binds
    abstract fun provideImageProcessor(imageProcessor: ImageProcessorRepoImpl): ImageProcessorRepo
}