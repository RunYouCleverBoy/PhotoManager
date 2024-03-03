package com.photomanager.photomanager.main.home

import androidx.lifecycle.viewModelScope
import com.photomanager.photomanager.main.home.repository.ImageProcessorRepo
import com.photomanager.photomanager.main.home.repository.ImagesRepo
import com.photomanager.photomanager.mvi.MVIViewModel
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class HomeViewModel @Inject constructor(
    private val imagesRepo: ImagesRepo,
    private val imageProcessorRepo: ImageProcessorRepo
) : MVIViewModel<HomeState, HomeEvent, HomeAction>(HomeState()) {
    override fun dispatchEvent(event: HomeEvent) {
        when (event) {
            is HomeEvent.OnImageClicked -> emit(HomeAction.OpenImage(event.uri))
            is HomeEvent.OnImagesPicked -> imagesRepo.addToFootage(event.uris.map {
                imageProcessorRepo.processImage(it)
            })

            is HomeEvent.OnApproachingFootageWindowEnd -> appendMoreDataToFootage()
            is HomeEvent.OnApproachingCollectionWindowEnd -> appendMoreDataToCollection()
        }
    }

    private fun appendMoreDataToFootage() {
        val size = state.value.footage.size
        val range = IntRange(size, size + PAGE_SIZE)
        viewModelScope.launch {
            imagesRepo.getFootage(state.value.footageSearchCriteria, range).collect { lst ->
                val mapped = lst.map { ImageUIDescriptor(it.uri, it.caption)}
                stateMutable.update { state -> state.copy(footage = state.footage + mapped) }
            }
        }
    }

    private fun appendMoreDataToCollection() {
        val size = state.value.collection.size
        val range = IntRange(size, size + PAGE_SIZE)
        viewModelScope.launch {
            imagesRepo.getFootage(state.value.collectionSearchCriteria, range).collect { lst ->
                val mapped = lst.map { ImageUIDescriptor(it.uri, it.caption)}
                stateMutable.update { state -> state.copy(collection = state.collection + mapped) }
            }
        }
    }

    companion object {
        private const val PAGE_SIZE = 300
    }
}

