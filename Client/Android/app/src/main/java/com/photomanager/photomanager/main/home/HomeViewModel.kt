package com.photomanager.photomanager.main.home

import androidx.lifecycle.viewModelScope
import com.photomanager.photomanager.main.home.di.WorkImagesRepo
import com.photomanager.photomanager.mvi.MVIViewModel
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class HomeViewModel @Inject constructor(
    private val workImagesRepo: WorkImagesRepo
) : MVIViewModel<HomeState, HomeEvent, HomeAction>(HomeState()) {
    override fun dispatchEvent(event: HomeEvent) {
        when (event) {
            is HomeEvent.OnImageClicked -> emit(HomeAction.MarkDone(event.uri))
            is HomeEvent.OnImagesPicked -> workImagesRepo.addImages(event.uris)
        }
    }

    init {
        viewModelScope.launch {
            workImagesRepo.images.collect { images ->
                stateMutable.update { state -> state.copy(collections = images) }
            }
        }
    }
}

