package com.photomanager.photomanager.main.home

import com.photomanager.photomanager.main.home.api.CollectionApi
import com.photomanager.photomanager.main.home.api.CollectionApiImpl
import com.photomanager.photomanager.main.home.di.WorkImagesRepo
import com.photomanager.photomanager.main.home.repository.WorkImagesRepoImpl
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.android.components.ViewModelComponent

@Module
@InstallIn(ViewModelComponent::class)
abstract class HomeProvider {
    @Binds
    abstract fun provideWorkImagesRepo(workImagesRepoImpl: WorkImagesRepoImpl): WorkImagesRepo

    @Binds
    abstract fun provideCollectionApi(collectionApi: CollectionApiImpl): CollectionApi
}