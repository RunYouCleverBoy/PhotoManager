package com.photomanager.photomanager.main.home.di

import com.photomanager.photomanager.main.home.api.ImagesApi
import com.photomanager.photomanager.main.home.api.ImagesApiImpl
import com.photomanager.photomanager.main.home.db.CollectionDao
import com.photomanager.photomanager.main.home.db.DatabaseHolder
import com.photomanager.photomanager.main.home.db.DatabaseHolderImpl
import com.photomanager.photomanager.main.home.db.FootageDao
import com.photomanager.photomanager.main.home.ktor.KtorFactory
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepo
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepoImpl
import com.photomanager.photomanager.main.home.repository.ImagesRepo
import com.photomanager.photomanager.main.home.repository.ImagesRepoImpl
import dagger.Binds
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ViewModelComponent
import io.ktor.client.HttpClient

@Module
@InstallIn(ViewModelComponent::class)
abstract class HomeProvider {
    @Binds
    abstract fun provideWorkImagesRepo(photoRepo: ImagesRepoImpl): ImagesRepo

    @Binds
    abstract fun provideDatabase(database: DatabaseHolderImpl): DatabaseHolder

    @Binds
    abstract fun provideImageProcessor(imageProcessor: ImageProcessorRepoImpl): ImageProcessorRepo

    @Binds
    abstract fun provideServerApi(serverApi: ImagesApiImpl): ImagesApi

    @Provides
    fun provideFootageDao(databaseHolder: DatabaseHolder): FootageDao = databaseHolder.database.footageDao()

    @Provides
    fun provideCollectionDao(databaseHolder: DatabaseHolder): CollectionDao = databaseHolder.database.collectionDao()

    @Provides
    fun provideKtorClient(): HttpClient = KtorFactory.createKtorClient()
}