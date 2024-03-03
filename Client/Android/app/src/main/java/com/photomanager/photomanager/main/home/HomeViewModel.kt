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
            is HomeEvent.OnImagesPicked -> addToFootage(event)
            is HomeEvent.OnAddToCollection -> addToCollection(event.ids)

            HomeEvent.OnApproachingFootageWindowEnd -> appendMoreData(toCollection = false)
            HomeEvent.OnApproachingCollectionWindowEnd -> appendMoreData(toCollection = true)
        }
    }

    private fun addToCollection(ids: List<String>) {
        viewModelScope.launch {
            val idSet = ids.toSet()
            imagesRepo.collection.add(state.value.footage.filter {
                it.id !in idSet
            }.map { it.id })
        }
    }

    private fun addToFootage(event: HomeEvent.OnImagesPicked) {
        viewModelScope.launch {
            imagesRepo.footage.add(event.uris.map { uri ->
                imageProcessorRepo.processExif(uri)
            })
        }
    }

    private fun appendMoreData(toCollection: Boolean) {
        if (state.value.isBusy) return
        stateMutable.update { state -> state.copy(isBusy = true) }
        val curr = if (toCollection) state.value.collection else state.value.footage
        val size = curr.size
        val range = IntRange(size, size + PAGE_SIZE)
        val searchCriteria = if (toCollection) state.value.collectionSearchCriteria else state.value.footageSearchCriteria
        val repo = if (toCollection) {
            imagesRepo.collection
        } else {
            imagesRepo.footage
        }
        viewModelScope.launch {
            val count = repo.getSize(searchCriteria)
            if (count == 0) {
                return@launch
            }
            repo.get(searchCriteria, range).collect { lst ->
                val ids = curr.map { it.id }.toSet()
                val mapped = lst
                    .filter { it.id !in ids }
                    .map { ImageUIDescriptor(it.id, it.uri, it.caption) }
                if (toCollection) {
                    stateMutable.update { state -> state.copy(collection = state.collection + mapped) }
                } else {
                    stateMutable.update { state -> state.copy(footage = state.footage + mapped) }
                }
            }
        }.invokeOnCompletion {
            stateMutable.update { state -> state.copy(isBusy = false) }
        }
    }

    companion object {
        private const val PAGE_SIZE = 300
    }
}

