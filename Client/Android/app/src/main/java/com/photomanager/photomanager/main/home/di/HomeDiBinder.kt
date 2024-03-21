package com.photomanager.photomanager.main.home.di

import com.photomanager.photomanager.main.home.api.ImagesApi
import com.photomanager.photomanager.main.home.api.ImagesApiImpl
import com.photomanager.photomanager.main.home.db.DatabaseHolder
import com.photomanager.photomanager.main.home.db.DatabaseHolderImpl
import com.photomanager.photomanager.main.home.db.dao.PhotoDao
import com.photomanager.photomanager.main.home.ktor.KtorFactory
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepo
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepoImpl
import com.photomanager.photomanager.main.home.repository.PhotoRepo
import com.photomanager.photomanager.main.home.repository.PhotoRepoImpl
import com.photomanager.photomanager.main.home.ui.HomeTabRepo
import com.photomanager.photomanager.main.home.ui.HomeTabRepoImpl
import com.photomanager.photomanager.utils.GeoLocationUtils
import com.photomanager.photomanager.utils.GeoLocationUtilsImpl
import dagger.Binds
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ViewModelComponent
import io.ktor.client.HttpClient

@Module
@InstallIn(ViewModelComponent::class)
abstract class HomeDiBinder {
    @Binds
    abstract fun provideWorkImagesRepo(photoRepo: PhotoRepoImpl): PhotoRepo

    @Binds
    abstract fun provideImageProcessor(imageProcessor: ImageProcessorRepoImpl): ImageProcessorRepo

    @Binds
    abstract fun provideDatabase(database: DatabaseHolderImpl): DatabaseHolder

    @Binds
    abstract fun provideServerApi(serverApi: ImagesApiImpl): ImagesApi

    @Binds
    abstract fun provideGeoLocationUtils(geoLocationUtils: GeoLocationUtilsImpl): GeoLocationUtils

    @Binds
    abstract fun provideTabRepo(tabRepo: HomeTabRepoImpl): HomeTabRepo
}

@Module
@InstallIn(ViewModelComponent::class)
class HomeDiProviders {
    @Provides
    fun providePhotoDao(databaseHolder: DatabaseHolder): PhotoDao =
        databaseHolder.database.photoDao()

    @Provides
    fun provideKtorClient(): HttpClient = KtorFactory.createKtorClient()

    @Provides
    fun provideHttpConfiguration(): KtorFactory.Configuration =
        KtorFactory.Configuration("http://localhost:8080")

}